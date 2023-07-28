package main

import (
   "154.pages.dev/google/play"
   "fmt"
   "net/http"
   "os"
   "strings"
   "time"
)

func (f flags) download(ref, name []byte) error {
   res, err := http.Get(string(ref))
   if err != nil {
      return err
   }
   defer res.Body.Close()
   file, err := os.Create(string(name))
   if err != nil {
      return err
   }
   defer file.Close()
   if _, err := file.ReadFrom(res.Body); err != nil {
      return err
   }
   return nil
}

func (f flags) do_header(dir, platform string) (*play.Header, error) {
   var head play.Header
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
   res, err := play.New_Auth(f.email, f.passwd)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return res.Write_File(dir + "/auth.txt")
}

func (f flags) do_delivery(head *play.Header) error {
   deliver, err := head.Delivery(f.doc, f.vc)
   if err != nil {
      return err
   }
   file := play.File{f.doc, f.vc}
   for _, split := range deliver.Split_Data() {
      _, ref := split.Download_URL()
      _, id := split.ID()
      if err := f.download(ref, file.APK(id)); err != nil {
         return err
      }
   }
   for _, add := range deliver.Additional_File() {
      _, ref := add.Download_URL()
      _, typ := add.File_Type()
      if err := f.download(ref, file.OBB(typ)); err != nil {
         return err
      }
   }
   _, ref := deliver.Download_URL()
   return f.download(ref, file.APK(nil))
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
