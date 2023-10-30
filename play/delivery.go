package play

import (
   "154.pages.dev/protobuf"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strconv"
)

// developer.android.com/guide/app-bundle
type Config_APK struct {
   m protobuf.Message
}

// developer.android.com/guide/app-bundle
func (s Config_APK) Config() (string, error) {
   if s, ok := s.m.String(1); ok {
      return s, nil
   }
   return "", errors.New("config")
}

// developer.android.com/guide/app-bundle
func (s Config_APK) URL() (string, error) {
   if s, ok := s.m.String(5); ok {
      return s, nil
   }
   return "", errors.New("URL")
}

type Delivery struct {
   App Application
   Checkin Checkin
   Token Access_Token
   m protobuf.Message
}

// developer.android.com/guide/app-bundle
func (d Delivery) Config_APKs() []Config_APK {
   var configs []Config_APK
   for _, f := range d.m {
      if f.Number == 15 {
         if config, ok := f.Message(); ok {
            configs = append(configs, Config_APK{config})
         }
      }
   }
   return configs
}

func (d *Delivery) Delivery(single bool) error {
   req, err := http.NewRequest("GET", "https://play-fe.googleapis.com", nil)
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
   x_dfe_device_id(req, d.Checkin)
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

// developer.android.com/google/play/expansion-files
func (d Delivery) OBB_Files() []OBB_File {
   var files []OBB_File
   for _, f := range d.m {
      if f.Number == 4 {
         if file, ok := f.Message(); ok {
            files = append(files, OBB_File{file})
         }
      }
   }
   return files
}

func (d Delivery) URL() (string, error) {
   if v, ok := d.m.String(3); ok {
      return v, nil
   }
   return "", errors.New("URL")
}

// developer.android.com/google/play/expansion-files
type OBB_File struct {
   m protobuf.Message
}

// developer.android.com/google/play/expansion-files
func (o OBB_File) Role() (uint64, error) {
   if v, ok := o.m.Varint(1); ok {
      return v, nil
   }
   return 0, errors.New("role")
}

func (o OBB_File) URL() (string, error) {
   if s, ok := o.m.String(4); ok {
      return s, nil
   }
   return "", errors.New("URL")
}
