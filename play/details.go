package play

import (
   "154.pages.dev/encoding/protobuf"
   "154.pages.dev/strconv"
   "fmt"
   "io"
   "net/http"
)

func (d Details) Currency_Code() (string, error) {
   if v, ok := d.m.Message(8); ok { // Common.Offer[] offer
      if v, ok := v.String(2); ok { // String currencyCode
         return v, nil
      }
   }
   return "", fmt.Errorf("details, currency")
}

func (d Details) File() []File_Metadata {
   // details
   d.m, _ = d.m.Message(13)
   // appDetails
   d.m, _ = d.m.Message(1)
   var files []File_Metadata
   d.m.Messages(17, func(file protobuf.Message) {
      files = append(files, File_Metadata{file})
   })
   return files
}

func (d Details) Installation_Size() (uint64, error) {
   // details
   d.m, _ = d.m.Message(13)
   // appDetails
   d.m, _ = d.m.Message(1)
   // installationSize
   return d.m.Varint(9)
}

func (d Details) Micros() (uint64, error) {
   // offer
   d.m, _ = d.m.Message(8)
   return d.m.Varint(1)
}

func (d Details) Num_Downloads() (uint64, error) {
   // details
   d.m, _ = d.m.Message(13)
   // appDetails
   d.m, _ = d.m.Message(1)
   // I dont know the name of field 70, but the similar field 13 is called
   // numDownloads
   return d.m.Varint(70)
}

func (d Details) String() string {
   var b []byte
   b = append(b, "creator: "...)
   {
      v, _ := d.Creator()
      b = append(b, v...)
   }
   b = append(b, "\nfile:"...)
   for _, file := range d.File() {
      if v, _ := file.File_Type(); v >= 1 {
         b = append(b, " OBB"...)
      } else {
         b = append(b, " APK"...)
      }
   }
   b = append(b, "\ninstallation size: "...)
   {
      v, _ := d.Installation_Size()
      b = fmt.Append(b, strconv.Size(v))
   }
   b = append(b, "\ndownloads: "...)
   {
      v, _ := d.Num_Downloads()
      b = fmt.Append(b, strconv.Cardinal(v))
   }
   b = append(b, "\noffer: "...)
   {
      v, _ := d.Micros()
      b = fmt.Append(b, v)
   }
   b = append(b, ' ')
   {
      v, _ := d.Currency_Code()
      b = append(b, v...)
   }
   b = append(b, "\ntitle: "...)
   {
      v, _ := d.Title()
      b = append(b, v...)
   }
   b = append(b, "\nupload date: "...)
   {
      v, _ := d.Upload_Date()
      b = append(b, v...)
   }
   b = append(b, "\nversion: "...)
   {
      v, _ := d.Version()
      b = append(b, v...)
   }
   b = append(b, "\nversion code: "...)
   {
      v, _ := d.Version_Code()
      b = fmt.Append(b, v)
   }
   return string(b)
}

func (d Details) Title() (string, error) {
   return d.m.String(5)
}

func (d Details) Upload_Date() (string, error) {
   // details
   d.m, _ = d.m.Message(13)
   // appDetails
   d.m, _ = d.m.Message(1)
   return d.m.String(16)
}

func (d Details) Version() (string, error) {
   // details
   d.m, _ = d.m.Message(13)
   // appDetails
   d.m, _ = d.m.Message(1)
   // versionString
   return d.m.String(4)
}

func (d Details) Version_Code() (uint64, error) {
   // details
   d.m, _ = d.m.Message(13)
   // appDetails
   d.m, _ = d.m.Message(1)
   return d.m.Varint(3)
}

// FileMetadata
// This is similar to AppFileMetadata, but notably field 4 is different.
type File_Metadata struct {
   m protobuf.Message
}

// fileType
func (f File_Metadata) File_Type() (uint64, error) {
   return f.m.Varint(1)
}

func (h Header) Details(doc string) (*Details, error) {
   req, err := http.NewRequest(
      "GET", "https://android.clients.google.com/fdfe/details?doc=" + doc, nil,
   )
   if err != nil {
      return nil, err
   }
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
   // ResponseWrapper
   mes, err := protobuf.Consume(data)
   if err != nil {
      return nil, err
   }
   mes, err = mes.Message(1)
   if err != nil {
      return nil, fmt.Errorf("payload not found")
   }
   // detailsResponse
   mes, _ = mes.Message(2)
   // docV2
   mes, _ = mes.Message(4)
   return &Details{mes}, nil
}

type Details struct {
   m protobuf.Message
}

func (d Details) Creator() (string, error) {
   v, ok := d.m.String(6)
   if ok {
      return v, nil
   }
   return "", fmt.Errorf("details, creator")
}
