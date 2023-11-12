package main

import (
   "154.pages.dev/google/play"
   "fmt"
   "net/http"
   "os"
   "time"
   option "154.pages.dev/http"
)

func (f flags) client(a *play.Access_Token, c *play.Checkin) error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   home += "/google/play/"
   var token play.Refresh_Token
   token.Raw, err = os.ReadFile(home + "token.txt")
   if err != nil {
      return err
   }
   if err := token.Unmarshal(); err != nil {
      return err
   }
   if err := a.Refresh(token); err != nil {
      return err
   }
   c.Raw, err = os.ReadFile(home + f.platform + ".bin")
   if err != nil {
      return err
   }
   return c.Unmarshal()
}

func (f flags) do_acquire() error {
   var client play.Acquire
   err := f.client(&client.Token, &client.Checkin)
   if err != nil {
      return err
   }
   return client.Acquire(f.app.ID)
}

func (f flags) do_auth() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   token, err := play.Exchange(f.code)
   if err != nil {
      return err
   }
   return os.WriteFile(home + "/google/play/token.txt", token.Raw, 0666)
}

func (f flags) do_delivery() error {
   var client play.Delivery
   err := f.client(&client.Token, &client.Checkin)
   if err != nil {
      return err
   }
   client.App = f.app
   if err := client.Delivery(f.single); err != nil {
      return err
   }
   option.Location()
   for _, apk := range client.Config_APKs() {
      if url, ok := apk.URL(); ok {
         if config, ok := apk.Config(); ok {
            err := f.download(url, f.app.APK(config))
            if err != nil {
               return err
            }
         }
      }
   }
   for _, obb := range client.OBB_Files() {
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

func (f flags) do_device() error {
   name, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   name += "/google/play/" + f.platform + ".bin"
   var check play.Checkin
   play.Pixel_6.Platform = f.platform
   if err := check.Checkin(play.Pixel_6); err != nil {
      return err
   }
   if err := os.WriteFile(name, check.Raw, 0666); err != nil {
      return err
   }
   fmt.Println("Sleep(9*time.Second)")
   time.Sleep(9*time.Second)
   if err := check.Unmarshal(); err != nil {
      return err
   }
   return check.Sync(play.Pixel_6)
}

func (f flags) download(url, name string) error {
   res, err := http.Get(url)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   file, err := os.Create(name)
   if err != nil {
      return err
   }
   defer file.Close()
   pro := option.Progress_Length(res.ContentLength)
   if _, err := file.ReadFrom(pro.Reader(res)); err != nil {
      return err
   }
   return nil
}
