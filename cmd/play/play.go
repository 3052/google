package main

import (
   "154.pages.dev/google/play"
   "fmt"
   "net/http"
   "os"
   "time"
   option "154.pages.dev/http"
)

func (f flags) do_delivery(head *play.Header) error {
   deliver, err := head.Delivery(f.doc, f.version)
   if err != nil {
      return err
   }
   name := play.Name{f.doc, f.version}
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
      if err := f.download(ref, name.APK(id)); err != nil {
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
      if err := f.download(ref, name.OBB(role)); err != nil {
         return err
      }
   }
   ref, err := deliver.URL()
   if err != nil {
      return err
   }
   return f.download(ref, name.APK(""))
}

func (f flags) do_auth(dir string) error {
   text, err := play.New_Refresh_Token(f.code)
   if err != nil {
      return err
   }
   return os.WriteFile(dir + "/token.txt", text, 0666)
}

func (f flags) do_header(dir string) (*play.Header, error) {
   var head play.Header
   head.Set_Agent(f.single)
   {
      b, err := os.ReadFile(dir + "/token.txt")
      if err != nil {
         return nil, err
      }
      if err := head.Set_Authorization(b); err != nil {
         return nil, err
      }
   }
   {
      b, err := os.ReadFile(dir + "/" + play.Phone.Platform + ".bin")
      if err != nil {
         return nil, err
      }
      if err := head.Set_Device(b); err != nil {
         return nil, err
      }
   }
   return &head, nil
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
func (f flags) do_device(dir string) error {
   data, err := play.New_Checkin(play.Phone)
   if err != nil {
      return err
   }
   os.WriteFile(dir + "/" + play.Phone.Platform + ".bin", data, 0666)
   fmt.Println("Sleep(9*time.Second)")
   time.Sleep(9*time.Second)
   var head play.Header
   head.Set_Agent(f.single)
   if err := head.Set_Device(data); err != nil {
      return err
   }
   return head.Sync(play.Phone)
}

