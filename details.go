package googleplay

import (
   "2a.pages.dev/rosso/protobuf"
   "2a.pages.dev/rosso/strconv"
   "errors"
   "io"
   "net/http"
   "net/url"
)

var err_device = errors.New("your device isn't compatible with this version")

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

func (d Details) MarshalText() ([]byte, error) {
   var b []byte
   b = append(b, "Title: "...)
   if val, err := d.Title(); err != nil {
      return nil, err
   } else {
      b = append(b, val...)
   }
   b = append(b, "\nCreator: "...)
   if val, err := d.Creator(); err != nil {
      return nil, err
   } else {
      b = append(b, val...)
   }
   b = append(b, "\nUpload Date: "...)
   if val, err := d.Upload_Date(); err != nil {
      return nil, err
   } else {
      b = append(b, val...)
   }
   b = append(b, "\nVersion: "...)
   if val, err := d.Version(); err != nil {
      return nil, err
   } else {
      b = append(b, val...)
   }
   b = append(b, "\nVersion Code: "...)
   if val, err := d.Version_Code(); err != nil {
      return nil, err
   } else {
      b = strconv.AppendUint(b, val, 10)
   }
   b = append(b, "\nNum Downloads: "...)
   if val, err := d.Num_Downloads(); err != nil {
      return nil, err
   } else {
      b = strconv.New_Number(val).Cardinal(b)
   }
   b = append(b, "\nInstallation Size: "...)
   if val, err := d.Installation_Size(); err != nil {
      return nil, err
   } else {
      b = strconv.New_Number(val).Size(b)
   }
   b = append(b, "\nFile:"...)
   for _, file := range d.File() {
      if val, err := file.File_Type(); err != nil {
         return nil, err
      } else if val >= 1 {
         b = append(b, " OBB"...)
      } else {
         b = append(b, " APK"...)
      }
   }
   b = append(b, "\nOffer: "...)
   if val, err := d.Micros(); err != nil {
      return nil, err
   } else {
      b = strconv.AppendUint(b, val, 10)
   }
   b = append(b, ' ')
   if val, err := d.Currency_Code(); err != nil {
      return nil, err
   } else {
      b = append(b, val...)
   }
   return append(b, '\n'), nil
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

func (h Header) Details(app string) (*Details, error) {
   req, err := http.NewRequest(
      "GET", "https://android.clients.google.com/fdfe/details", nil,
   )
   if err != nil {
      return nil, err
   }
   // half of the apps I test require User-Agent,
   // so just set it for all of them
   h.Set_Agent(req.Header)
   h.Set_Auth(req.Header)
   h.Set_Device(req.Header)
   req.URL.RawQuery = "doc=" + url.QueryEscape(app)
   res, err := Client.Do(req)
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
