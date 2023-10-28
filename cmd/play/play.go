package main

import (
   "154.pages.dev/google/play"
   "fmt"
   "net/http"
   "os"
   "strings"
   "time"
   option "154.pages.dev/http"
)

func (f flags) do_acquire() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   var req play.Acquire_Request
   req.Token, err = func() (play.Access_Token, error) {
      b, err := os.ReadFile(home + "/google/play/token.txt")
      if err != nil {
         return nil, err
      }
      t, err := play.Raw_Refresh_Token.Token(b)
      if err != nil {
         return nil, err
      }
      return t.Do()
   }()
   if err != nil {
      return err
   }
   req.Checkin, err = func() (*play.Checkin, error) {
      s := func() string {
         var b strings.Builder
         b.WriteString(home)
         b.WriteString("/google/play/")
         b.WriteString(play.Platforms[f.platform])
         b.WriteString(".bin")
         return b.String()
      }()
      b, err := os.ReadFile(s)
      if err != nil {
         return nil, err
      }
      return play.Raw_Checkin.Checkin(b)
   }()
   if err != nil {
      return err
   }
   return req.Do(f.app.ID)
}

func (f flags) do_auth() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   var raw play.Raw_Refresh_Token
   if err := raw.Do(f.code); err != nil {
      return err
   }
   return os.WriteFile(home + "google/play/token.txt", raw, 0666)
}

func (f flags) do_delivery() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   var req play.Delivery_Request
   req.Token, err = func() (play.Access_Token, error) {
      b, err := os.ReadFile(home + "/google/play/token.txt")
      if err != nil {
         return nil, err
      }
      t, err := play.Raw_Refresh_Token.Token(b)
      if err != nil {
         return nil, err
      }
      return t.Do()
   }()
   if err != nil {
      return err
   }
   req.Checkin, err = func() (*play.Checkin, error) {
      s := func() string {
         var b strings.Builder
         b.WriteString(home)
         b.WriteString("/google/play/")
         b.WriteString(play.Platforms[f.platform])
         b.WriteString(".bin")
         return b.String()
      }()
      b, err := os.ReadFile(s)
      if err != nil {
         return nil, err
      }
      return play.Raw_Checkin.Checkin(b)
   }()
   if err != nil {
      return err
   }
   req.App = f.app
   deliver, err := req.Do(f.single)
   if err != nil {
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

func (f flags) do_details() (*play.Details, error) {
   home, err := os.UserHomeDir()
   if err != nil {
      return nil, err
   }
   var req play.Details_Request
   req.Token, err = func() (play.Access_Token, error) {
      b, err := os.ReadFile(home + "/google/play/token.txt")
      if err != nil {
         return nil, err
      }
      t, err := play.Raw_Refresh_Token.Token(b)
      if err != nil {
         return nil, err
      }
      return t.Do()
   }()
   if err != nil {
      return nil, err
   }
   req.Checkin, err = func() (*play.Checkin, error) {
      s := func() string {
         var b strings.Builder
         b.WriteString(home)
         b.WriteString("/google/play/")
         b.WriteString(play.Platforms[f.platform])
         b.WriteString(".bin")
         return b.String()
      }()
      b, err := os.ReadFile(s)
      if err != nil {
         return nil, err
      }
      return play.Raw_Checkin.Checkin(b)
   }()
   if err != nil {
      return nil, err
   }
   return req.Do(f.app.ID, f.single)
}

func (f flags) do_device() error {
   sync := play.Sync_Request{Device: play.Phone}
   sync.Device.Platform = play.Platforms[f.platform]
   raw, err := sync.Device.Checkin()
   if err != nil {
      return err
   }
   name, err := func() (string, error) {
      s, err := os.UserHomeDir()
      if err != nil {
         return "", err
      }
      var b strings.Builder
      b.WriteString(s)
      b.WriteByte('/')
      b.WriteString(sync.Device.Platform)
      b.WriteString(".bin")
      return b.String(), nil
   }()
   if err != nil {
      return err
   }
   os.WriteFile(name, raw, 0666)
   fmt.Println("Sleep(9*time.Second)")
   time.Sleep(9*time.Second)
   sync.Checkin, err = raw.Checkin()
   if err != nil {
      return err
   }
   return sync.Do()
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
