package main

import (
   "154.pages.dev/google/play"
   "154.pages.dev/http/option"
   "fmt"
   "net/http"
   "os"
   "time"
)

func (f flags) do_delivery() error {
   head, err := f.do_header(dir)
   if err != nil {
      return err
   }
   deliver, err := head.Delivery(f.doc, f.vc)
   if err != nil {
      return err
   }
   file := play.File{f.doc, f.vc}
   option.Location()
   for _, split := range deliver.Split() {
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
   text, err := play.New_Refresh_Token(f.code)
   if err != nil {
      return err
   }
   return os.WriteFile(dir + "/token.txt", text, 0666)
}

func (f flags) do_details() error {
   head, err := f.do_header(dir)
   if err != nil {
      return err
   }
   detail, err := head.Details(f.doc)
   if err != nil {
      return err
   }
   fmt.Println(detail)
   return nil
}

func (f flags) do_device(dir) error {
   data, err := play.Phone.Checkin()
   if err != nil {
      return err
   }
   err = os.WriteFile(dir + "/" + f.platform + ".bin", data, 0666)
   if err != nil {
      return err
   }
   fmt.Println("Sleep(9*time.Second)")
   time.Sleep(9*time.Second)
   return nil
}

func (f flags) do_header(dir) (*play.Header, error) {
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
      b, err := os.ReadFile(dir + "/" + f.platform + ".bin")
      if err != nil {
         return nil, err
      }
      if err := head.Set_Device(b); err != nil {
         return nil, err
      }
   }
   return &head, nil
}

func (f flags) do_acquire() error {
   head, err := f.do_header(dir)
   if err != nil {
      return err
   }
   return head.Acquire(f.doc)
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

