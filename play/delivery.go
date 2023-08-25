package play

import (
   "154.pages.dev/encoding/protobuf"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strconv"
)

// AppFileMetadata
type App_File_Metadata struct {
   m protobuf.Message
}

// downloadUrl
func (a App_File_Metadata) Download_URL() (string, error) {
   return a.m.String(4)
}

// fileType
func (a App_File_Metadata) File_Type() (uint64, error) {
   return a.m.Varint(1)
}

// AndroidAppDeliveryData
type Delivery struct {
   m protobuf.Message
}

func (d Delivery) Additional_File() []App_File_Metadata {
   var files []App_File_Metadata
   // additionalFile
   d.m.Messages(4, func(file protobuf.Message) {
      files = append(files, App_File_Metadata{file})
   })
   return files
}

// downloadUrl
func (d Delivery) Download_URL() (string, error) {
   return d.m.String(3)
}

func (d Delivery) Split_Data() []Split_Data {
   var splits []Split_Data
   // splitDeliveryData
   d.m.Messages(15, func(split protobuf.Message) {
      splits = append(splits, Split_Data{split})
   })
   return splits
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
   h.agent(req)
   h.authorization(req)
   h.device(req)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   data, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   // ResponseWrapper
   mes, err := protobuf.Consume(data)
   if err != nil {
      return nil, err
   }
   // payload
   mes, _ = mes.Message(1)
   // deliveryResponse
   mes, _ = mes.Message(21)
   status, err := mes.Varint(1)
   if err != nil {
      return nil, err
   }
   switch status {
   case 3:
      return nil, errors.New("purchase required")
   case 5:
      return nil, errors.New("invalid version")
   }
   mes, err = mes.Message(2)
   if err != nil {
      return nil, errors.New("appDeliveryData not found")
   }
   return &Delivery{mes}, nil
}

// SplitDeliveryData
type Split_Data struct {
   m protobuf.Message
}

// downloadUrl
func (s Split_Data) Download_URL() (string, error) {
   return s.m.String(5)
}

// id
func (s Split_Data) ID() (string, error) {
   return s.m.String(1)
}
