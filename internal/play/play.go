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

func (f flags) do_device() error {
   if f.leanback {
      play.Device.Feature = append(play.Device.Feature, play.Leanback)
   }
   data, err := play.Device.GoogleCheckin()
   if err != nil {
      return err
   }
   err = os.WriteFile(f.device_path(), data, 0666)
   if err != nil {
      return err
   }
   var checkin play.GoogleCheckin
   err = checkin.Unmarshal(data)
   if err != nil {
      return err
   }
   fmt.Println("Sleep(9*time.Second)")
   time.Sleep(9*time.Second)
   return play.Device.Sync(checkin)
}

func (f flags) do_delivery() error {
   var checkin play.GoogleCheckin
   auth, err := f.client(&checkin)
   if err != nil {
      return err
   }
   deliver, err := auth.Delivery(checkin, f.app, f.single)
   if err != nil {
      return err
   }
   for apk := range deliver.Apk() {
      if address, ok := apk.Url(); ok {
         if v, ok := apk.Field1(); ok {
            err := f.download(address, f.app.Apk(v))
            if err != nil {
               return err
            }
         }
      }
   }
   for obb := range deliver.Obb() {
      if address, ok := obb.Url(); ok {
         if v, ok := obb.Field1(); ok {
            err := f.download(address, f.app.Obb(v))
            if err != nil {
               return err
            }
         }
      }
   }
   if v, ok := deliver.Url(); ok {
      err := f.download(v, f.app.Apk(""))
      if err != nil {
         return err
      }
   }
   return nil
}

func (f flags) download(address, name string) error {
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
func (f flags) do_acquire() error {
   var checkin play.GoogleCheckin
   auth, err := f.client(&checkin)
   if err != nil {
      return err
   }
   return auth.Acquire(checkin, f.app.Id)
}

func (f flags) do_auth() error {
   text, err := play.NewGoogleToken(f.code)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/token.txt", text, 0666)
}
func (f flags) do_details() (*play.Details, error) {
   var checkin play.GoogleCheckin
   auth, err := f.client(&checkin)
   if err != nil {
      return nil, err
   }
   return auth.Details(checkin, f.app.Id, f.single)
}

func (f flags) client(checkin *play.GoogleCheckin) (*play.GoogleAuth, error) {
   data, err := os.ReadFile(f.home + "/token.txt")
   if err != nil {
      return nil, err
   }
   var token play.GoogleToken
   err = token.Unmarshal(data)
   if err != nil {
      return nil, err
   }
   auth, err := token.Auth()
   if err != nil {
      return nil, err
   }
   data, err = os.ReadFile(f.device_path())
   if err != nil {
      return nil, err
   }
   err = checkin.Unmarshal(data)
   if err != nil {
      return nil, err
   }
   return auth, nil
}

func (f flags) device_path() string {
   var b strings.Builder
   b.WriteString(f.home)
   b.WriteByte('/')
   b.WriteString(play.Device.Abi)
   if f.leanback {
      b.WriteString("-leanback")
   }
   b.WriteString(".bin")
   return b.String()
}
