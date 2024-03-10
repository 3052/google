package play

import (
   "154.pages.dev/protobuf"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strconv"
)

func (d Delivery) ObbFile() func() (*ObbFile, bool) {
   iterate := d.m.Iterate(4)
   return func() (*ObbFile, bool) {
      if v, ok := iterate(); ok {
         return &ObbFile{v}, true
      }
      return nil, false
   }
}

func (d Delivery) ConfigApk() func() (*ConfigApk, bool) {
   iterate := d.m.Iterate(15)
   return func() (*ConfigApk, bool) {
      if v, ok := iterate(); ok {
         return &ConfigApk{v}, true
      }
      return nil, false
   }
}

// developer.android.com/guide/app-bundle
type ConfigApk struct {
   m protobuf.Message
}

type Delivery struct {
   App Application
   Checkin Checkin
   Token AccessToken
   m protobuf.Message
}

// developer.android.com/google/play/expansion-files
type ObbFile struct {
   m protobuf.Message
}

// developer.android.com/google/play/expansion-files
func (o ObbFile) Role() (uint64, bool) {
   if v, ok := o.m.GetVarint(1); ok {
      return uint64(v), true
   }
   return 0, false
}

// developer.android.com/guide/app-bundle
func (c ConfigApk) Config() (string, bool) {
   if v, ok := c.m.GetBytes(1); ok {
      return string(v), true
   }
   return "", false
}

func (c ConfigApk) URL() (string, bool) {
   if v, ok := c.m.GetBytes(5); ok {
      return string(v), true
   }
   return "", false
}

func (o ObbFile) URL() (string, bool) {
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

func (d *Delivery) Do(single bool) error {
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
   var ok bool
   if d.m, ok = d.m.Get(1); ok {
      if d.m, ok = d.m.Get(21); ok {
         if v, ok := d.m.GetVarint(1); ok {
            switch v {
            case 3:
               return errors.New("acquire")
            case 5:
               return errors.New("version code")
            }
            if d.m, ok = d.m.Get(2); ok {
               return nil
            }
            return errors.New("Delivery.Get[1][21][2]")
         }
         return errors.New("Delivery.Get[1][21][1]")
      }
      return errors.New("Delivery.Get[1][21]")
   }
   return errors.New("Delivery.Get[1]")
}
