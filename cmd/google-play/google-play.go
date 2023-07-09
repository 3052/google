package main

import (
   "fmt"
   "google-play.pages.dev"
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
   if _, err := file.ReadFrom(res.Body); err != nil {
      return err
   }
   return nil
}

func (f flags) do_header(dir, platform string) (*google_play.Header, error) {
   var head google_play.Header
   err := head.Read_Auth(dir + "/auth.txt")
   if err != nil {
      return nil, err
   }
   if err := head.Auth.Exchange(); err != nil {
      return nil, err
   }
   if err := head.Read_Device(dir + "/" + platform + ".bin"); err != nil {
      return nil, err
   }
   head.Single = f.single
   return &head, nil
}

func (f flags) do_auth(dir string) error {
   if f.file != "" {
      raw, err := os.ReadFile(f.file)
      if err != nil {
         return err
      }
      f.passwd = strings.TrimSpace(string(raw))
   }
   res, err := google_play.New_Auth(f.email, f.passwd)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return res.Write_File(dir + "/auth.txt")
}

func (f flags) do_delivery(head *google_play.Header) error {
   deliver, err := head.Delivery(f.doc, f.vc)
   if err != nil {
      return err
   }
   file := google_play.File{f.doc, f.vc}
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

func (f flags) do_details(head *google_play.Header) ([]byte, error) {
   detail, err := head.Details(f.doc)
   if err != nil {
      return nil, err
   }
   return detail.MarshalText()
}

func (f flags) do_device(dir, platform string) error {
   res, err := google_play.Phone.Checkin(platform)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   fmt.Printf("Sleeping %v for server to process\n", google_play.Sleep)
   time.Sleep(google_play.Sleep)
   return res.Write_File(dir + "/" + platform + ".bin")
}
