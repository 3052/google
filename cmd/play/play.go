package main

import (
   "154.pages.dev/google/play"
   "154.pages.dev/http/option"
   "fmt"
   "net/http"
   "os"
   "strings"
   "time"
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

func (f flags) do_delivery(head *play.Header) error {
   deliver, err := head.Delivery(f.doc, f.vc)
   if err != nil {
      return err
   }
   file := play.File{f.doc, f.vc}
   option.Location()
   for _, split := range deliver.Split_Data() {
      ref, err := split.Download_URL()
      if err != nil {
         return err
      }
      id, err := split.ID()
      if err != nil {
         return err
      }
      if err := f.download(ref, file.APK(id)); err != nil {
         return err
      }
   }
   for _, add := range deliver.Additional_File() {
      ref, err := add.Download_URL()
      if err != nil {
         return err
      }
      typ, err := add.File_Type()
      if err != nil {
         return err
      }
      if err := f.download(ref, file.OBB(typ)); err != nil {
         return err
      }
   }
   ref, err := deliver.Download_URL()
   if err != nil {
      return err
   }
   return f.download(ref, file.APK(""))
}

func (f flags) do_auth(dir string) error {
   if f.file != "" {
      raw, err := os.ReadFile(f.file)
      if err != nil {
         return err
      }
      f.passwd = strings.TrimSpace(string(raw))
   }
   auth, err := play.New_Auth(f.email, f.passwd)
   if err != nil {
      return err
   }
   {
      b, err := auth.MarshalText()
      if err != nil {
         return err
      }
      os.WriteFile(dir + "/auth.txt", b, 0666)
   }
   return nil
}

func (f flags) do_details(head *play.Header) ([]byte, error) {
   detail, err := head.Details(f.doc)
   if err != nil {
      return nil, err
   }
   return detail.MarshalText()
}

func (f flags) do_device(dir, platform string) error {
   res, err := play.Phone.Checkin(platform)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   fmt.Printf("Sleeping %v for server to process\n", play.Sleep)
   time.Sleep(play.Sleep)
   return res.Write_File(dir + "/" + platform + ".bin")
}

func (f flags) do_header(dir, platform string) (*play.Header, error) {
   var head play.Header
   head.Auth = make(play.Auth)
   {
      b, err := os.ReadFile(dir + "/auth.txt")
      if err != nil {
         return nil, err
      }
      head.Auth.UnmarshalText(b)
   }
   err := head.Auth.Exchange()
   if err != nil {
      return nil, err
   }
   {
      b, err := os.ReadFile(dir + "/" + platform + ".bin")
      if err != nil {
         return nil, err
      }
      head.Device.UnmarshalBinary(b)
   }
   head.Single = f.single
   return &head, nil
}

