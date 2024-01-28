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
      if field.Number == 4 {
         if file, ok := field.Message(); ok {
            if f(OBB_File{file}) {
               return
            }
         }
      }
   }
}

func (d Delivery) Config_APK(f func(Config_APK) bool) {
   for _, field := range d.m {
      if field.Number == 15 {
         if config, ok := field.Message(); ok {
            if f(Config_APK{config}) {
               return
            }
         }
      }
   }
}

// developer.android.com/guide/app-bundle
type Config_APK struct {
   m protobuf.Message
}

// developer.android.com/guide/app-bundle
func (s Config_APK) Config() (string, bool) {
   return s.m.String(1)
}

// developer.android.com/guide/app-bundle
func (s Config_APK) URL() (string, bool) {
   return s.m.String(5)
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
   d.m, err = func() (protobuf.Message, error) {
      b, err := io.ReadAll(res.Body)
      if err != nil {
         return nil, err
      }
      return protobuf.Consume(b)
   }()
   if err != nil {
      return err
   }
   d.m.Message(1)
   d.m.Message(21)
   status_code, ok := d.m.Varint(1)
   if !ok {
      return errors.New("status code")
   }
   switch status_code {
   case 3:
      return errors.New("acquire")
   case 5:
      return errors.New("version code")
   }
   d.m.Message(2)
   return nil
}

func (d Delivery) URL() (string, bool) {
   return d.m.String(3)
}

// developer.android.com/google/play/expansion-files
type OBB_File struct {
   m protobuf.Message
}

// developer.android.com/google/play/expansion-files
func (o OBB_File) Role() (uint64, bool) {
   return o.m.Varint(1)
}

func (o OBB_File) URL() (string, bool) {
   return o.m.String(4)
}
