package play

import (
   "154.pages.dev/protobuf"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strconv"
)

type Name struct {
   Package string
   Version_Code uint64
}

func (n Name) APK(config string) string {
   var b []byte
   b = append(b, n.Package...)
   b = append(b, '-')
   if config != "" {
      b = append(b, config...)
      b = append(b, '-')
   }
   b = strconv.AppendUint(b, n.Version_Code, 10)
   b = append(b, ".apk"...)
   return string(b)
}

func (n Name) OBB(role uint64) string {
   var b []byte
   if role >= 1 {
      b = append(b, "patch"...)
   } else {
      b = append(b, "main"...)
   }
   b = append(b, '.')
   b = strconv.AppendUint(b, n.Version_Code, 10)
   b = append(b, '.')
   b = append(b, n.Package...)
   b = append(b, ".obb"...)
   return string(b)
}

func (h Header) Delivery(doc string, vc uint64) (*Delivery, error) {
   req, err := http.NewRequest("GET", "https://play-fe.googleapis.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/fdfe/delivery"
   req.URL.RawQuery = url.Values{
      "doc": {doc},
      "vc":  {strconv.FormatUint(vc, 10)},
   }.Encode()
   req.Header.Set(h.Authorization())
   req.Header.Set(h.User_Agent())
   req.Header.Set(h.X_DFE_Device_ID())
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   mes, err := func() (protobuf.Message, error) {
      b, err := io.ReadAll(res.Body)
      if err != nil {
         return nil, err
      }
      return protobuf.Consume(b)
   }()
   if err != nil {
      return nil, err
   }
   mes.Message(1)
   mes.Message(21)
   status_code, ok := mes.Varint(1)
   if !ok {
      return nil, errors.New("status code")
   }
   switch status_code {
   case 3:
      return nil, errors.New("acquire")
   case 5:
      return nil, errors.New("version code")
   }
   mes.Message(2)
   return &Delivery{mes}, nil
}

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
