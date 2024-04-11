package main

import (
   "154.pages.dev/google/play"
   "154.pages.dev/log"
   "fmt"
   "net/http"
   "os"
   "time"
)

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

func (f flags) do_device() error {
   play.Phone.ABI = f.platform.String()
   var checkin play.GoogleCheckin
   if err := checkin.Checkin(play.Phone); err != nil {
      return err
   }
   f.home += fmt.Sprintf("/%v.bin", f.platform)
   if err := os.WriteFile(f.home, checkin.Data, 0666); err != nil {
      return err
   }
   fmt.Println("Sleep(9*time.Second)")
   time.Sleep(9*time.Second)
   if err := checkin.Unmarshal(); err != nil {
      return err
   }
   return checkin.Sync(play.Phone)
}
func (f flags) do_auth() error {
   var token play.GoogleToken
   if err := token.Auth(f.code); err != nil {
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
   if err := token.Unmarshal(); err != nil {
      return err
   }
   if err := auth.Auth(token); err != nil {
      return err
   }
   checkin.Data, err = os.ReadFile(fmt.Sprint(f.home, "/", f.platform, ".bin"))
   if err != nil {
      return err
   }
   return checkin.Unmarshal()
}

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
   if _, err := dst.ReadFrom(meter.Reader(res)); err != nil {
      return err
   }
   return nil
}
