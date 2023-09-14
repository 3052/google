package play

import (
   "154.pages.dev/encoding/protobuf"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strconv"
)

func (h Header) Delivery(doc string, vc uint64) (*Delivery, error) {
   req, err := http.NewRequest(
      "GET", "https://play-fe.googleapis.com/fdfe/delivery", nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "doc": {doc},
      "vc": {strconv.FormatUint(vc, 10)},
   }.Encode()
   req.Header.Set(h.Agent())
   req.Header.Set(h.Authorization())
   req.Header.Set(h.Device())
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   data, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   mes, err := protobuf.Consume(data) // ResponseWrapper
   if err != nil {
      return nil, err
   }
   mes, _ = mes.Message(1) // payload
   mes, _ = mes.Message(21) // deliveryResponse
   status, ok := mes.Varint(1)
   if ok {
      switch status {
      case 3:
         return nil, errors.New("purchase required")
      case 5:
         return nil, errors.New("invalid version")
      }
   }
   mes, ok = mes.Message(2)
   if !ok {
      return nil, errors.New("appDeliveryData not found")
   }
   return &Delivery{mes}, nil
}

func (s Split_Delivery_Data) Download_URL() (string, error) {
   v, ok := s.m.String(5)
   if ok {
      return v, nil
   }
   return "", errors.New("split delivery data, download URL")
}

func (s Split_Delivery_Data) ID() (string, error) {
   v, ok := s.m.String(1)
   if ok {
      return v, nil
   }
   return "", errors.New("split delivery data, ID")
}

func (a App_File_Metadata) Download_URL() (string, error) {
   v, ok := a.m.String(4)
   if ok {
      return v, nil
   }
   return "", errors.New("app file metadata, download URL")
}

func (a App_File_Metadata) File_Type() (uint64, error) {
   v, ok := a.m.Varint(1)
   if ok {
      return v, nil
   }
   return 0, errors.New("app file metadata, file type")
}

// AndroidAppDeliveryData
type Delivery struct {
   m protobuf.Message
}

type App_File_Metadata struct {
   m protobuf.Message
}

func (d Delivery) Download_URL() (string, error) {
   v, ok := d.m.String(3)
   if ok {
      return v, nil
   }
   return "", errors.New("delivery, download URL")
}

type Split_Delivery_Data struct {
   m protobuf.Message
}

type File struct {
   Package_Name string
   Version_Code uint64
}

func (f File) APK(id string) string {
   var b []byte
   b = append(b, f.Package_Name...)
   b = append(b, '-')
   if id != "" {
      b = append(b, id...)
      b = append(b, '-')
   }
   b = strconv.AppendUint(b, f.Version_Code, 10)
   b = append(b, ".apk"...)
   return string(b)
}

func (f File) OBB(file_type uint64) string {
   var b []byte
   if file_type >= 1 {
      b = append(b, "patch"...)
   } else {
      b = append(b, "main"...)
   }
   b = append(b, '.')
   b = strconv.AppendUint(b, f.Version_Code, 10)
   b = append(b, '.')
   b = append(b, f.Package_Name...)
   b = append(b, ".obb"...)
   return string(b)
}

func (d Delivery) Additional_File() []App_File_Metadata {
   var files []App_File_Metadata
   for _, f := range d.m {
      if f.Number == 4 { // AppFileMetadata[] additionalFile
         if file, ok := f.Message(); ok {
            files = append(files, App_File_Metadata{file})
         }
      }
   }
   return files
}

func (d Delivery) Split() []Split_Delivery_Data {
   var splits []Split_Delivery_Data
   for _, f := range d.m {
      if f.Number == 15 {
         if split, ok := f.Message(); ok {
            splits = append(splits, Split_Delivery_Data{split})
         }
      }
   }
   return splits
}
