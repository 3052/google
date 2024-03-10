package play

import (
   "154.pages.dev/protobuf"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strconv"
)

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
   d.m = <-d.m.Get(1)
   d.m = <-d.m.Get(21)
   switch <-d.m.GetVarint(1) {
   case 3:
      return errors.New("acquire")
   case 5:
      return errors.New("version code")
   }
   d.m = <-d.m.Get(2)
   return nil
}

func (d Delivery) ObbFile() chan ObbFile {
   files := make(chan ObbFile)
   go func() {
      for file := range d.m.Get(4) {
         files <- ObbFile{file}
      }
      close(files)
   }()
   return files
}
func (d Delivery) ConfigApk() chan ConfigApk {
   apks := make(chan ConfigApk)
   go func() {
      for apk := range d.m.Get(15) {
         apks <- ConfigApk{apk}
      }
      close(apks)
   }()
   return apks
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
   if v, ok := <-o.m.GetVarint(1); ok {
      return uint64(v), true
   }
   return 0, false
}

// developer.android.com/guide/app-bundle
func (c ConfigApk) Config() (string, bool) {
   if v, ok := <-c.m.GetBytes(1); ok {
      return string(v), true
   }
   return "", false
}

func (c ConfigApk) URL() (string, bool) {
   if v, ok := <-c.m.GetBytes(5); ok {
      return string(v), true
   }
   return "", false
}

func (o ObbFile) URL() (string, bool) {
   if v, ok := <-o.m.GetBytes(4); ok {
      return string(v), true
   }
   return "", false
}

func (d Delivery) URL() (string, bool) {
   if v, ok := <-d.m.GetBytes(3); ok {
      return string(v), true
   }
   return "", false
}
