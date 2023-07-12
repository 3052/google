package play

import (
   "154.pages.dev/encoding/protobuf"
   "154.pages.dev/strconv"
   "fmt"
   "io"
   "net/http"
)

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
   response_wrapper, err := protobuf.Unmarshal(body)
   if err != nil {
      return nil, err
   }
   var det Details
   // .payload.detailsResponse.docV2
   det.Message = response_wrapper.Get(1).Get(2).Get(4)
   return &det, nil
}

// .details.appDetails.installationSize
func (d Details) Installation_Size() (uint64, error) {
   value, err := d.Get(13).Get(1).Get_Varint(9)
   if err != nil {
      return 0, err_device
   }
   return value, nil
}

// .details.appDetails.file
func (d Details) File() []File_Metadata {
   var files []File_Metadata
   for _, file := range d.Get(13).Get(1).Get_Messages(17) {
      files = append(files, File_Metadata{file})
   }
   return files
}

// FileMetadata
// This is similar to AppFileMetadata, but notably field 4 is different.
type File_Metadata struct {
   protobuf.Message
}

// .fileType
func (f File_Metadata) File_Type() (uint64, error) {
   return f.Get_Varint(1)
}

var err_device = fmt.Errorf("your device isn't compatible with this version")

type Details struct {
   protobuf.Message
}

// .creator
func (d Details) Creator() (string, error) {
   return d.Get_String(6)
}

// .offer.currencyCode
func (d Details) Currency_Code() (string, error) {
   return d.Get(8).Get_String(2)
}

// .offer.micros
func (d Details) Micros() (uint64, error) {
   return d.Get(8).Get_Varint(1)
}

// .title
func (d Details) Title() (string, error) {
   return d.Get_String(5)
}

// .details.appDetails.uploadDate
func (d Details) Upload_Date() (string, error) {
   value, err := d.Get(13).Get(1).Get_String(16)
   if err != nil {
      return "", err_device
   }
   return value, nil
}

// .details.appDetails.versionString
func (d Details) Version() (string, error) {
   value, err := d.Get(13).Get(1).Get_String(4)
   if err != nil {
      return "", err_device
   }
   return value, nil
}

// .details.appDetails.versionCode
func (d Details) Version_Code() (uint64, error) {
   value, err := d.Get(13).Get(1).Get_Varint(3)
   if err != nil {
      return 0, err_device
   }
   return value, nil
}

// .details.appDetails
// I dont know the name of field 70, but the similar field 13 is called
// .numDownloads
func (d Details) Num_Downloads() (uint64, error) {
   return d.Get(13).Get(1).Get_Varint(70)
}

func (d Details) MarshalText() ([]byte, error) {
   var b []byte
   b = append(b, "creator: "...)
   if v, err := d.Creator(); err != nil {
      return nil, err
   } else {
      b = append(b, v...)
   }
   b = append(b, "\nfile:"...)
   for _, file := range d.File() {
      if v, err := file.File_Type(); err != nil {
         return nil, err
      } else if v >= 1 {
         b = append(b, " OBB"...)
      } else {
         b = append(b, " APK"...)
      }
   }
   b = append(b, "\ninstallation size: "...)
   if v, err := d.Installation_Size(); err != nil {
      return nil, err
   } else {
      b = fmt.Append(b, strconv.Size(v))
   }
   b = append(b, "\ndownloads: "...)
   if v, err := d.Num_Downloads(); err != nil {
      return nil, err
   } else {
      b = fmt.Append(b, strconv.Cardinal(v))
   }
   b = append(b, "\noffer: "...)
   if v, err := d.Micros(); err != nil {
      return nil, err
   } else {
      b = fmt.Append(b, v)
   }
   b = append(b, ' ')
   if v, err := d.Currency_Code(); err != nil {
      return nil, err
   } else {
      b = append(b, v...)
   }
   b = append(b, "\ntitle: "...)
   if v, err := d.Title(); err != nil {
      return nil, err
   } else {
      b = append(b, v...)
   }
   b = append(b, "\nupload date: "...)
   if v, err := d.Upload_Date(); err != nil {
      return nil, err
   } else {
      b = append(b, v...)
   }
   b = append(b, "\nversion: "...)
   if v, err := d.Version(); err != nil {
      return nil, err
   } else {
      b = append(b, v...)
   }
   b = append(b, "\nversion code: "...)
   if v, err := d.Version_Code(); err != nil {
      return nil, err
   } else {
      b = fmt.Append(b, v)
   }
   return b, nil
}

