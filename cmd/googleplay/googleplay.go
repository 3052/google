package main

import (
   "2a.pages.dev/googleplay"
   "2a.pages.dev/rosso/http"
   "fmt"
   "io"
   "os"
   "time"
)

func (f flags) do_header(dir, platform string) (*googleplay.Header, error) {
   var head googleplay.Header
   err := head.Open_Auth(dir + "/auth.txt")
   if err != nil {
      return nil, err
   }
   if err := f.Exchange(&head.Auth); err != nil {
      return nil, err
   }
   if err := head.Open_Device(dir + "/" + platform + ".bin"); err != nil {
      return nil, err
   }
   head.Single = f.single
   return &head, nil
}

func (f flags) do_details(h *googleplay.Header) ([]byte, error) {
   detail, err := f.Details(h, f.doc)
   if err != nil {
      return nil, err
   }
   return detail.MarshalText()
}

func (f flags) do_device(dir, platform string) error {
   res, err := f.Checkin(googleplay.Phone, platform)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   fmt.Printf("Sleeping %v for server to process\n", googleplay.Sleep)
   time.Sleep(googleplay.Sleep)
   return res.Create(dir + "/" + platform + ".bin")
}

func (f flags) do_auth(dir string) error {
   res, err := f.Auth(f.email, f.passwd)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return res.Create(dir + "/auth.txt")
}

func (f flags) download(ref, name string) error {
   f.CheckRedirect = nil
   req := http.Get()
   err := req.URL_String(ref)
   if err != nil {
      return err
   }
   res, err := f.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   file, err := os.Create(name)
   if err != nil {
      return err
   }
   defer file.Close()
   pro := http.Progress_Bytes(file, res.ContentLength)
   if _, err := io.Copy(pro, res.Body); err != nil {
      return err
   }
   return nil
}

func (f flags) do_delivery(h *googleplay.Header) error {
   deliver, err := f.Delivery(h, f.doc, f.vc)
   if err != nil {
      return err
   }
   file := googleplay.File{f.doc, f.vc}
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

