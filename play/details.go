package play

import (
   "154.pages.dev/encoding/protobuf"
   "154.pages.dev/strconv"
   "fmt"
   "io"
   "net/http"
)

type Details struct {
   m protobuf.Message
}

func (d Details) Installation_Size() (uint64, error) {
   // details
   d.m, _ = d.m.Message(13)
   // appDetails
   d.m, _ = d.m.Message(1)
   // installationSize
   return d.m.Varint(9)
}

// FileMetadata
// This is similar to AppFileMetadata, but notably field 4 is different.
type File_Metadata struct {
   m protobuf.Message
}

func (d Details) Micros() (uint64, error) {
   // offer
   d.m, _ = d.m.Message(8)
   return d.m.Varint(1)
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

func (d Details) Num_Downloads() (uint64, error) {
   // details
   d.m, _ = d.m.Message(13)
   // appDetails
   d.m, _ = d.m.Message(1)
   // I dont know the name of field 70, but the similar field 13 is called
   // numDownloads
   return d.m.Varint(70)
}

func (d Details) Creator() (string, error) {
   return d.m.String(6)
}

func (d Details) Currency_Code() (string, error) {
   // offer
   d.m, _ = d.m.Message(8)
   return d.m.String(2)
}

// fileType
func (f File_Metadata) File_Type() (uint64, error) {
   return f.m.Varint(1)
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

func (d Details) MarshalText() ([]byte, error) {
   var b []byte
   b = append(b, "creator: "...)
   if v, err := d.Creator(); err == nil {
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
   if v, err := d.Installation_Size(); err == nil {
      b = fmt.Append(b, strconv.Size(v))
   }
   b = append(b, "\ndownloads: "...)
   if v, err := d.Num_Downloads(); err == nil {
      b = fmt.Append(b, strconv.Cardinal(v))
   }
   b = append(b, "\noffer: "...)
   if v, err := d.Micros(); err == nil {
      b = fmt.Append(b, v)
   }
   b = append(b, ' ')
   if v, err := d.Currency_Code(); err == nil {
      b = append(b, v...)
   }
   b = append(b, "\ntitle: "...)
   if v, err := d.Title(); err == nil {
      b = append(b, v...)
   }
   b = append(b, "\nupload date: "...)
   if v, err := d.Upload_Date(); err == nil {
      b = append(b, v...)
   }
   b = append(b, "\nversion: "...)
   if v, err := d.Version(); err == nil {
      b = append(b, v...)
   }
   b = append(b, "\nversion code: "...)
   if v, err := d.Version_Code(); err == nil {
      b = fmt.Append(b, v)
   }
   return b, nil
}

type Header struct {
   h http.Header
}

func (a Access_Token) Header(d *Device, single bool) (*Header, error) {
   h := make(http.Header)
   h.Set("Authorization", "Bearer " + a.auth())
   {
      var b []byte
      // `sdk` is needed for `/fdfe/delivery`
      b = append(b, "Android-Finsky (sdk="...)
      // valid range 0 - 0x7FFF_FFFF
      b = fmt.Append(b, 9)
      // com.android.vending
      b = append(b, ",versionCode="...)
      if single {
         // valid range 8032_0000 - 8091_9999
         b = fmt.Append(b, 8091_9999)
      } else {
         // valid range 8092_0000 - 0x7FFF_FFFF
         b = fmt.Append(b, 9999_9999)
      }
      b = append(b, ')')
      h.Set("User-Agent", string(b))
   }
   {
      id, err := d.ID()
      if err != nil {
         return nil, err
      }
      h.Set("X-DFE-Device-ID", fmt.Sprintf("%x", id))
   }
   return &Header{h}, nil
}

func (h Header) Details(doc string) (*Details, error) {
   req, err := http.NewRequest(
      "GET", "https://android.clients.google.com/fdfe/details?doc=" + doc, nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header = h.h
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
