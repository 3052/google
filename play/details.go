package play

import (
   "154.pages.dev/encoding/protobuf"
   "154.pages.dev/strconv"
   "fmt"
   "io"
   "net/http"
)

func (d Details) Creator() (int, []byte) {
   return d.m.Bytes(6)
}

func (d Details) MarshalText() ([]byte, error) {
   var b []byte
   b = append(b, "creator: "...)
   if i, v := d.Creator(); i >= 0 {
      b = append(b, v...)
   }
   b = append(b, "\nfile:"...)
   for _, file := range d.File() {
      if _, v := file.File_Type(); v >= 1 {
         b = append(b, " OBB"...)
      } else {
         b = append(b, " APK"...)
      }
   }
   b = append(b, "\ninstallation size: "...)
   if i, v := d.Installation_Size(); i >= 0 {
      b = fmt.Append(b, strconv.Size(v))
   }
   b = append(b, "\ndownloads: "...)
   if i, v := d.Num_Downloads(); i >= 0 {
      b = fmt.Append(b, strconv.Cardinal(v))
   }
   b = append(b, "\noffer: "...)
   if i, v := d.Micros(); i >= 0 {
      b = fmt.Append(b, v)
   }
   b = append(b, ' ')
   if i, v := d.Currency_Code(); i >= 0 {
      b = append(b, v...)
   }
   b = append(b, "\ntitle: "...)
   if i, v := d.Title(); i >= 0 {
      b = append(b, v...)
   }
   b = append(b, "\nupload date: "...)
   if i, v := d.Upload_Date(); i >= 0 {
      b = append(b, v...)
   }
   b = append(b, "\nversion: "...)
   if i, v := d.Version(); i >= 0 {
      b = append(b, v...)
   }
   b = append(b, "\nversion code: "...)
   if i, v := d.Version_Code(); i >= 0 {
      b = fmt.Append(b, v)
   }
   return b, nil
}

func (d Details) Currency_Code() (int, []byte) {
   // offer
   _, d.m = d.m.Message(8)
   return d.m.Bytes(2)
}

// fileType
func (f File_Metadata) File_Type() (int, uint64) {
   return f.m.Uvarint(1)
}

func (d Details) File() []File_Metadata {
   var files []File_Metadata
   // details
   _, d.m = d.m.Message(13)
   // appDetails
   _, d.m = d.m.Message(1)
   for {
      i, file := d.m.Message(17)
      if i == -1 {
         return files
      }
      files = append(files, File_Metadata{file})
      d.m = d.m[i+1:]
   }
}

func (d Details) Installation_Size() (int, uint64) {
   // details
   _, d.m = d.m.Message(13)
   // appDetails
   _, d.m = d.m.Message(1)
   // installationSize
   return d.m.Uvarint(9)
}

func (h Header) Details(doc string) (*Details, error) {
   req, err := http.NewRequest(
      "GET", "https://android.clients.google.com/fdfe/details?doc=" + doc, nil,
   )
   if err != nil {
      return nil, err
   }
   // half of the apps I test require User-Agent,
   // so just set it for all of them
   h.Set_Agent(req.Header)
   h.Set_Auth(req.Header)
   h.Set_Device(req.Header)
   res, err := client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   body, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   // ResponseWrapper
   mes, err := protobuf.Unmarshal(body)
   if err != nil {
      return nil, err
   }
   // payload
   _, mes = mes.Message(1)
   // detailsResponse
   _, mes = mes.Message(2)
   // docV2
   _, mes = mes.Message(4)
   return &Details{mes}, nil
}

// FileMetadata
// This is similar to AppFileMetadata, but notably field 4 is different.
type File_Metadata struct {
   m protobuf.Message
}

var err_device = fmt.Errorf("your device isn't compatible with this version")

type Details struct {
   m protobuf.Message
}

func (d Details) Micros() (int, uint64) {
   // offer
   _, d.m = d.m.Message(8)
   return d.m.Uvarint(1)
}

func (d Details) Title() (int, []byte) {
   return d.m.Bytes(5)
}

func (d Details) Upload_Date() (int, []byte) {
   // details
   _, d.m = d.m.Message(13)
   // appDetails
   _, d.m = d.m.Message(1)
   return d.m.Bytes(16)
}

func (d Details) Version() (int, []byte) {
   // details
   _, d.m = d.m.Message(13)
   // appDetails
   _, d.m = d.m.Message(1)
   // versionString
   return d.m.Bytes(4)
}

func (d Details) Version_Code() (int, uint64) {
   // details
   _, d.m = d.m.Message(13)
   // appDetails
   _, d.m = d.m.Message(1)
   return d.m.Uvarint(3)
}

func (d Details) Num_Downloads() (int, uint64) {
   // details
   _, d.m = d.m.Message(13)
   // appDetails
   _, d.m = d.m.Message(1)
   // I dont know the name of field 70, but the similar field 13 is called
   // numDownloads
   return d.m.Uvarint(70)
}

