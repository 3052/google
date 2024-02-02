package play

import (
   "154.pages.dev/protobuf"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strconv"
)

func (d Delivery) OBB_File(f func(OBB_File) bool) {
   for _, field := range d.m {
      if file, ok := field.Get(4); ok {
         if !f(OBB_File{file}) {
            return
         }
      }
   }
}

func (d Delivery) Config_APK(f func(Config_APK) bool) {
   for _, field := range d.m {
      if config, ok := field.Get(15); ok {
         if !f(Config_APK{config}) {
            return
         }
      }
   }
}

// developer.android.com/guide/app-bundle
type Config_APK struct {
   m protobuf.Message
}

type Delivery struct {
   App Application
   Checkin Checkin
   Token Access_Token
   m protobuf.Message
}

func (d *Delivery) Delivery(single bool) error {
   req, err := http.NewRequest("GET", "https://android.clients.google.com", nil)
   if err != nil {
      return err
   }
   req.URL.Path = "/fdfe/delivery"
   req.URL.RawQuery = url.Values{
      "doc": {d.App.ID},
      "vc":  {strconv.FormatUint(d.App.Version, 10)},
   }.Encode()
   authorization(req, d.Token)
   user_agent(req, single)
   if err := x_dfe_device_id(req, d.Checkin); err != nil {
      return err
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   data, err := io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   if err := d.m.Consume(data); err != nil {
      return err
   }
   d.m, _ = d.m.Get(1)
   d.m, _ = d.m.Get(21)
   status_code, ok := d.m.GetVarint(1)
   if !ok {
      return errors.New("status code")
   }
   switch status_code {
   case 3:
      return errors.New("acquire")
   case 5:
      return errors.New("version code")
   }
   d.m, _ = d.m.Get(2)
   return nil
}

// developer.android.com/google/play/expansion-files
type OBB_File struct {
   m protobuf.Message
}

// developer.android.com/google/play/expansion-files
func (o OBB_File) Role() (uint64, bool) {
   if v, ok := o.m.GetVarint(1); ok {
      return uint64(v), true
   }
   return 0, false
}

// developer.android.com/guide/app-bundle
func (s Config_APK) Config() (string, bool) {
   if v, ok := s.m.GetBytes(1); ok {
      return string(v), true
   }
   return "", false
}

func (s Config_APK) URL() (string, bool) {
   if v, ok := s.m.GetBytes(5); ok {
      return string(v), true
   }
   return "", false
}

func (o OBB_File) URL() (string, bool) {
   if v, ok := o.m.GetBytes(4); ok {
      return string(v), true
   }
   return "", false
}

func (d Delivery) URL() (string, bool) {
   if v, ok := d.m.GetBytes(3); ok {
      return string(v), true
   }
   return "", false
}
