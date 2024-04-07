package main

import (
   "154.pages.dev/google/play"
   "154.pages.dev/log"
   "fmt"
   "net/http"
   "os"
   "time"
)

func (f flags) client(a *play.AccessToken, check *play.Checkin) error {
   var (
      token play.RefreshToken
      err error
   )
   token.Data, err = os.ReadFile(f.home + "/token.txt")
   if err != nil {
      return err
   }
   if err := token.Unmarshal(); err != nil {
      return err
   }
   if err := a.Refresh(token); err != nil {
      return err
   }
   check.Data, err = os.ReadFile(fmt.Sprint(f.home, "/", f.platform, ".bin"))
   if err != nil {
      return err
   }
   return check.Unmarshal()
}

func (f flags) do_delivery() error {
   var client play.Delivery
   err := f.client(&client.Token, &client.Checkin)
   if err != nil {
      return err
   }
   client.App = f.app
   if err := client.Do(f.single); err != nil {
      return err
   }
   for apk := range client.ConfigApk() {
      if url, ok := apk.URL(); ok {
         if config, ok := apk.Config(); ok {
            err := f.download(url, f.app.APK(config))
            if err != nil {
               return err
            }
         }
      }
   }
   for obb := range client.ObbFile() {
      if url, ok := obb.URL(); ok {
         if role, ok := obb.Role(); ok {
            err := f.download(url, f.app.OBB(role))
            if err != nil {
               return err
            }
         }
      }
   }
   if url, ok := client.URL(); ok {
      err := f.download(url, f.app.APK(""))
      if err != nil {
         return err
      }
   }
   return nil
}

func (f flags) do_details() (*play.Details, error) {
   var client play.Details
   err := f.client(&client.Token, &client.Checkin)
   if err != nil {
      return nil, err
   }
   if err := client.Details(f.app.ID, f.single); err != nil {
      return nil, err
   }
   return &client, nil
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

func (f flags) do_acquire() error {
   var client play.Acquire
   err := f.client(&client.Token, &client.Checkin)
   if err != nil {
      return err
   }
   return client.Do(f.app.ID)
}

func (f flags) do_device() error {
   name, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   name += fmt.Sprintf("/google-play/%v.bin", f.platform)
   var check play.Checkin
   play.Phone.Platform = f.platform.String()
   if err := check.Do(play.Phone); err != nil {
      return err
   }
   if err := os.WriteFile(name, check.Data, 0666); err != nil {
      return err
   }
   fmt.Println("Sleep(9*time.Second)")
   time.Sleep(9*time.Second)
   if err := check.Unmarshal(); err != nil {
      return err
   }
   return check.Sync(play.Phone)
}

func (f flags) do_auth() error {
   var token play.RefreshToken
   if err := token.New(f.code); err != nil {
      return err
   }
   return os.WriteFile(f.home + "/token.txt", token.Data, 0666)
}
