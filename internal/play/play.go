package main

import (
   "154.pages.dev/google/play"
   "154.pages.dev/log"
   "fmt"
   "net/http"
   "os"
   "time"
)

func (f flags) do_delivery() error {
   var (
      auth play.GoogleAuth
      checkin play.GoogleCheckin
   )
   err := f.client(&auth, &checkin)
   if err != nil {
      return err
   }
   deliver, err := checkin.Delivery(auth, f.app, f.single)
   if err != nil {
      return err
   }
   for apk := range deliver.APK() {
      if url, ok := apk.URL(); ok {
         if v, ok := apk.Field1(); ok {
            err := f.download(url, f.app.APK(v))
            if err != nil {
               return err
            }
         }
      }
   }
   for obb := range deliver.OBB() {
      if url, ok := obb.URL(); ok {
         if v, ok := obb.Field1(); ok {
            err := f.download(url, f.app.OBB(v))
            if err != nil {
               return err
            }
         }
      }
   }
   if v, ok := deliver.URL(); ok {
      err := f.download(v, f.app.APK(""))
      if err != nil {
         return err
      }
   }
   return nil
}

func (f flags) download(url, name string) error {
   dst, err := os.Create(name)
   if err != nil {
      return err
   }
   defer dst.Close()
   res, err := http.Get(url)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   var meter log.ProgressMeter
   meter.Set(1)
   _, err = dst.ReadFrom(meter.Reader(res))
   if err != nil {
      return err
   }
   return nil
}

func (f flags) do_acquire() error {
   var (
      auth play.GoogleAuth
      checkin play.GoogleCheckin
   )
   err := f.client(&auth, &checkin)
   if err != nil {
      return err
   }
   return checkin.Acquire(auth, f.app.ID)
}

func (f flags) do_auth() error {
   var token play.GoogleToken
   err := token.Auth(f.code)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/token.txt", token.Data, 0666)
}

func (f flags) do_details() (*play.Details, error) {
   var (
      auth play.GoogleAuth
      checkin play.GoogleCheckin
   )
   err := f.client(&auth, &checkin)
   if err != nil {
      return nil, err
   }
   return checkin.Details(auth, f.app.ID, f.single)
}

func (f flags) client(auth *play.GoogleAuth, checkin *play.GoogleCheckin) error {
   var (
      token play.GoogleToken
      err error
   )
   token.Data, err = os.ReadFile(f.home + "/token.txt")
   if err != nil {
      return err
   }
   err = token.Unmarshal()
   if err != nil {
      return err
   }
   err = auth.Auth(token)
   if err != nil {
      return err
   }
   device, err := f.device.MarshalText()
   if err != nil {
      return err
   }
   f.home += func() string {
      var b strings.Builder
      b.WriteByte('/')
      b.Write(device)
      b.WriteByte('-')
      b.WriteString(f.device.ABI)
      b.WriteString(".bin")
      return b.String()
   }()
   checkin.Data, err = os.ReadFile(f.home)
   if err != nil {
      return err
   }
   return checkin.Unmarshal()
}

func (f flags) do_device() error {
   var checkin play.GoogleCheckin
   err := checkin.Checkin(f.device)
   if err != nil {
      return err
   }
   f.home += fmt.Sprintf("/%v.bin", f.device.ABI)
   err = os.WriteFile(f.home, checkin.Data, 0666)
   if err != nil {
      return err
   }
   fmt.Println("Sleep(9*time.Second)")
   time.Sleep(9*time.Second)
   err = checkin.Unmarshal()
   if err != nil {
      return err
   }
   return checkin.Sync(f.device)
}
