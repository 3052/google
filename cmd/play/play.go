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
   var acquire play.Acquire
   acquire.Token, err = func() (play.Access_Token, error) {
      b, err := os.ReadFile(home + "/google/play/token.txt")
      if err != nil {
         return nil, err
      }
      m, err := play.Raw_Refresh_Token.Refresh_Token(b)
      if err != nil {
         return nil, err
      }
      return m.Auth()
   }()
   if err != nil {
      return err
   }
   acquire.Checkin, err = func() (*play.Checkin, error) {
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
   return acquire.Acquire(f.app.ID)
}

func (f flags) do_auth() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   raw, err := play.Auth(f.code)
   if err != nil {
      return err
   }
   return os.WriteFile(home + "/google/play/token.txt", raw, 0666)
}
func (f flags) do_delivery() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   var deliver play.Delivery
   deliver.Token, err = func() (play.Access_Token, error) {
      b, err := os.ReadFile(home + "/google/play/token.txt")
      if err != nil {
         return nil, err
      }
      m, err := play.Raw_Refresh_Token.Refresh_Token(b)
      if err != nil {
         return nil, err
      }
      return m.Auth()
   }()
   if err != nil {
      return err
   }
   deliver.Checkin, err = func() (*play.Checkin, error) {
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

func (f flags) do_details() (*play.Details, error) {
   home, err := os.UserHomeDir()
   if err != nil {
      return nil, err
   }
   var detail play.Details
   detail.Token, err = func() (play.Access_Token, error) {
      b, err := os.ReadFile(home + "/google/play/token.txt")
      if err != nil {
         return nil, err
      }
      m, err := play.Raw_Refresh_Token.Refresh_Token(b)
      if err != nil {
         return nil, err
      }
      return m.Auth()
   }()
   if err != nil {
      return nil, err
   }
   detail.Checkin, err = func() (*play.Checkin, error) {
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
   if err := detail.Details(f.app.ID, f.single); err != nil {
      return nil, err
   }
   return &detail, nil
}

func (f flags) do_device() error {
   play.Phone.Platform = play.Platforms[f.platform]
   name, err := func() (string, error) {
      s, err := os.UserHomeDir()
      if err != nil {
         return "", err
      }
      var b strings.Builder
      b.WriteString(s)
      b.WriteByte('/')
      b.WriteString(play.Phone.Platform)
      b.WriteString(".bin")
      return b.String(), nil
   }()
   if err != nil {
      return err
   }
   raw, err := play.Phone.Checkin()
   if err != nil {
      return err
   }
   os.WriteFile(name, raw, 0666)
   fmt.Println("Sleep(9*time.Second)")
   time.Sleep(9*time.Second)
   check, err := raw.Checkin()
   if err != nil {
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
