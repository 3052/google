package main

import (
   "154.pages.dev/google/play"
   "154.pages.dev/text"
   "fmt"
   "net/http"
   "os"
   "strings"
   "time"
)

func (f *flags) do_delivery() error {
   checkin := &play.GoogleCheckin{}
   auth, err := f.client(checkin)
   if err != nil {
      return err
   }
   deliver, err := auth.Delivery(checkin, &f.app, f.single)
   if err != nil {
      return err
   }
   apks := deliver.Apk()
   for {
      apk, ok := apks()
      if !ok {
         break
      }
      if address, ok := apk.Url(); ok {
         if v, ok := apk.Field1(); ok {
            err := download(address, f.app.Apk(v))
            if err != nil {
               return err
            }
         }
      }
   }
   obbs := deliver.Obb()
   for {
      obb, ok := obbs()
      if !ok {
         break
      }
      if address, ok := obb.Url(); ok {
         if v, ok := obb.Field1(); ok {
            err := download(address, f.app.Obb(v))
            if err != nil {
               return err
            }
         }
      }
   }
   if v, ok := deliver.Url(); ok {
      err := download(v, f.app.Apk(""))
      if err != nil {
         return err
      }
   }
   return nil
}

func (f *flags) do_acquire() error {
   checkin := &play.GoogleCheckin{}
   auth, err := f.client(checkin)
   if err != nil {
      return err
   }
   return auth.Acquire(checkin, f.app.Id)
}

func (f *flags) do_device() error {
   if f.leanback {
      play.Device.Feature = append(play.Device.Feature, play.Leanback)
   }
   checkin, err := play.Device.Checkin()
   if err != nil {
      return err
   }
   err = os.WriteFile(f.device_path(), checkin.Raw, 0666)
   if err != nil {
      return err
   }
   err = checkin.Unmarshal()
   if err != nil {
      return err
   }
   fmt.Println("Sleep(9*time.Second)")
   time.Sleep(9*time.Second)
   return play.Device.Sync(checkin)
}

func (f *flags) device_path() string {
   var b strings.Builder
   b.WriteString(f.home)
   b.WriteByte('/')
   b.WriteString(play.Device.Abi)
   if f.leanback {
      b.WriteString("-leanback")
   }
   b.WriteString(".txt")
   return b.String()
}

func (f *flags) client(checkin *play.GoogleCheckin) (*play.GoogleAuth, error) {
   var (
      token play.GoogleToken
      err error
   )
   token.Raw, err = os.ReadFile(f.home + "/token.txt")
   if err != nil {
      return nil, err
   }
   err = token.Unmarshal()
   if err != nil {
      return nil, err
   }
   auth, err := token.Auth()
   if err != nil {
      return nil, err
   }
   checkin.Raw, err = os.ReadFile(f.device_path())
   if err != nil {
      return nil, err
   }
   err = checkin.Unmarshal()
   if err != nil {
      return nil, err
   }
   return auth, nil
}

func download(address, name string) error {
   dst, err := os.Create(name)
   if err != nil {
      return err
   }
   defer dst.Close()
   resp, err := http.Get(address)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   var meter text.ProgressMeter
   meter.Set(1)
   _, err = dst.ReadFrom(meter.Reader(resp))
   if err != nil {
      return err
   }
   return nil
}

func (f *flags) do_details() (*play.Details, error) {
   checkin := &play.GoogleCheckin{}
   auth, err := f.client(checkin)
   if err != nil {
      return nil, err
   }
   return auth.Details(checkin, f.app.Id, f.single)
}

func (f *flags) do_auth() error {
   var token play.GoogleToken
   err := token.New(f.code)
   if err != nil {
      return err
   }
   os.Mkdir(f.home, 0666)
   return os.WriteFile(f.home + "/token.txt", token.Raw, 0666)
}
