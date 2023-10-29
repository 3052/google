package main

import (
   "154.pages.dev/google/play"
   "fmt"
   "net/http"
   "os"
   "time"
   option "154.pages.dev/http"
)

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

func (f flags) do_acquire() error {
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
   token.Unmarshal()
   var client play.Acquire
   client.Token.Refresh(token)
   client.Checkin.Raw, err = os.ReadFile(home + f.platform + ".bin")
   if err != nil {
      return err
   }
   client.Checkin.Unmarshal()
   return client.Acquire(f.app.ID)
}

func (f flags) do_device() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   home += "/google/play/"
   play.Phone.Platform = f.platform
   check, err := play.Phone.Checkin()
   if err != nil {
      return err
   }
   os.WriteFile(home + f.platform + ".bin", check.Raw, 0666)
   fmt.Println("Sleep(9*time.Second)")
   time.Sleep(9*time.Second)
   check.Unmarshal()
   return play.Phone.Sync(check)
}

func (f flags) do_details() (*play.Details, error) {
   home, err := os.UserHomeDir()
   if err != nil {
      return nil, err
   }
   home += "/google/play/"
   var token play.Refresh_Token
   token.Raw, err = os.ReadFile(home + "token.txt")
   if err != nil {
      return nil, err
   }
   token.Unmarshal()
   var client play.Details
   client.Token.Refresh(token)
   client.Checkin.Raw, err = os.ReadFile(home + f.platform + ".bin")
   if err != nil {
      return nil, err
   }
   client.Checkin.Unmarshal()
   if err := client.Details(f.app.ID, f.single); err != nil {
      return nil, err
   }
   return &client, nil
}

func (f flags) do_delivery() error {
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
   token.Unmarshal()
   var client play.Delivery
   client.Token.Refresh(token)
   client.Checkin.Raw, err = os.ReadFile(home + f.platform + ".bin")
   if err != nil {
      return err
   }
   client.Checkin.Unmarshal()
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
