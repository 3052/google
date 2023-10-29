package main

import (
   "154.pages.dev/google/play"
   "fmt"
   "net/http"
   "os"
   "time"
   option "154.pages.dev/http"
)

func (f flags) client(c *play.Client) error {
   dir, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   dir += "/google/play/"
   var token play.Refresh_Token
   token.Raw, err = os.ReadFile(dir + "token.txt")
   if err != nil {
      return err
   }
   if err := token.Unmarshal(); err != nil {
      return err
   }
   if err := c.Token.Refresh(token); err != nil {
      return err
   }
   c.Checkin.Raw, err = os.ReadFile(dir + f.platform + ".bin")
   if err != nil {
      return err
   }
   return c.Checkin.Unmarshal()
}

func (f flags) do_device() error {
   name, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   name += "/google/play/" + f.platform + ".bin"
   var check play.Checkin
   play.Phone.Platform = f.platform
   check.Checkin(play.Phone)
   os.WriteFile(name, check.Raw, 0666)
   fmt.Println("Sleep(9*time.Second)")
   time.Sleep(9*time.Second)
   check.Unmarshal()
   return check.Sync(play.Phone)
}

func (f flags) do_acquire() error {
   var client play.Client
   f.client(&client)
   return client.Acquire(f.app.ID)
}

func (f flags) do_details() (*play.Details, error) {
   var detail play.Details
   f.client(&detail.Client)
   if err := detail.Details(f.app.ID, f.single); err != nil {
      return nil, err
   }
   return &detail, nil
}

func (f flags) do_delivery() error {
   var deliver play.Delivery
   f.client(&deliver.Client)
   deliver.App = f.app
   if err := deliver.Delivery(f.single); err != nil {
      return err
   }
   option.Location()
   for _, apk := range deliver.Config_APKs() {
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
   for _, obb := range deliver.OBB_Files() {
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
   ref, err := deliver.URL()
   if err != nil {
      return err
   }
   return f.download(ref, f.app.APK(""))
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
