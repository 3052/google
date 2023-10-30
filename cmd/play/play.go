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
      ref, err := apk.URL()
      if err != nil {
         return err
      }
      id, err := apk.Config()
      if err != nil {
         return err
      }
      if err := f.download(ref, f.app.APK(id)); err != nil {
         return err
      }
   }
   for _, obb := range client.OBB_Files() {
      ref, err := obb.URL()
      if err != nil {
         return err
      }
      role, err := obb.Role()
      if err != nil {
         return err
      }
      if err := f.download(ref, f.app.OBB(role)); err != nil {
         return err
      }
   }
   ref, err := client.URL()
   if err != nil {
      return err
   }
   return f.download(ref, f.app.APK(""))
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
   play.Phone.Platform = f.platform
   if err := check.Checkin(play.Phone); err != nil {
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
   return check.Sync(play.Phone)
}

func (f flags) download(ref, name string) error {
   res, err := http.Get(ref)
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
