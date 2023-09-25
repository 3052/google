package play

import (
   "154.pages.dev/protobuf"
   "encoding/base64"
   "errors"
   "fmt"
   "net/url"
   "strconv"
   "strings"
   "time"
)

var Phone = Config{
   System_Available_Feature: []string{
      // app.source.getcontact
      "android.hardware.location.gps",
      // br.com.rodrigokolb.realdrum
      "android.software.midi",
      // com.app.xt
      "android.hardware.camera.front",
      // com.cabify.rider
      "android.hardware.camera.flash",
      // com.clearchannel.iheartradio.controller
      "android.hardware.microphone",
      // com.google.android.apps.walletnfcrel
      "android.software.device_admin",
      // com.google.android.youtube
      "android.hardware.touchscreen",
      "android.hardware.wifi",
      // com.madhead.tos.zh
      "android.hardware.sensor.accelerometer",
      // com.miHoYo.GenshinImpact
      "android.hardware.opengles.aep",
      // com.pinterest
      "android.hardware.camera",
      "android.hardware.location",
      "android.hardware.screen.portrait",
      // com.supercell.brawlstars
      "android.hardware.touchscreen.multitouch",
      // com.sygic.aura
      "android.hardware.location.network",
      // com.xiaomi.smarthome
      "android.hardware.bluetooth",
      "android.hardware.bluetooth_le",
      "android.hardware.camera.autofocus",
      "android.hardware.usb.host",
      // kr.sira.metal
      "android.hardware.sensor.compass",
      // org.thoughtcrime.securesms
      "android.hardware.telephony",
      // org.videolan.vlc
      "android.hardware.screen.landscape",
   },
   System_Shared_Library: []string{
      // com.amctve.amcfullepisodes
      "org.apache.http.legacy",
      // com.binance.dev
      "android.test.runner",
      // com.miui.weather2
      "global-miui11-empty.jar",
   },
   GL_Extension: []string{
      // com.instagram.android
      "GL_OES_compressed_ETC1_RGB8_texture",
      // com.kakaogames.twodin
      "GL_KHR_texture_compression_astc_ldr",
   },
   // com.axis.drawingdesk.v3
   // valid range 0x03_00_01 - math.MaxInt32  
   GL_ES_Version: 0xFF_FF_FF,
}

// These can use default values, but they must all be included
type Config struct {
   GL_ES_Version uint64
   GL_Extension []string
   Has_Five_Way_Navigation bool
   Has_Hard_Keyboard bool
   Keyboard uint64
   Navigation uint64
   Screen_Density uint64
   Screen_Layout uint64
   System_Available_Feature []string
   System_Shared_Library []string
   Touch_Screen uint64
}

// Checkin$AndroidCheckinResponse
type Device struct {
   m protobuf.Message
}

func (d Device) android_ID() (uint64, error) {
   v, ok := d.m.Fixed64(7) // long androidId_
   if ok {
      return v, nil
   }
   return 0, errors.New("androidId_")
}

type Native_Platform map[int64]string

var Platforms = Native_Platform{
   // com.google.android.youtube
   0: "x86",
   // com.miui.weather2
   1: "armeabi-v7a",
   // com.kakaogames.twodin
   2: "arm64-v8a",
}

func (n Native_Platform) String() string {
   var b []byte
   b = append(b, "native platform"...)
   for key, value := range n {
      b = append(b, '\n')
      b = strconv.AppendInt(b, key, 10)
      b = append(b, ": "...)
      b = append(b, value...)
   }
   return string(b)
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

func (d Details) Currency_Code() (string, error) {
   if v, ok := d.m.Message(8); ok { // Common.Offer[] offer
      if v, ok := v.String(2); ok { // String currencyCode
         return v, nil
      }
   }
   return "", fmt.Errorf("details, currency")
}

func (d Details) File() []File_Metadata {
   d.m, _ = d.m.Message(13) // DocDetails.DocumentDetails details
   d.m, _ = d.m.Message(1) // AppDetails appDetails
   var files []File_Metadata
   for _, f := range d.m {
      if f.Number == 17 { // FileMetadata[] file
         if file, ok := f.Message(); ok {
            files = append(files, File_Metadata{file})
         }
      }
   }
   return files
}

func (d Details) Installation_Size() (uint64, error) {
   d.m, _ = d.m.Message(13) // DocDetails.DocumentDetails details
   d.m, _ = d.m.Message(1) // AppDetails appDetails
   v, ok := d.m.Varint(9) // long installationSize
   if ok {
      return v, nil
   }
   return 0, fmt.Errorf("details, installation size")
}

func (d Details) Micros() (uint64, error) {
   d.m, _ = d.m.Message(8) // Common.Offer[] offer
   v, ok := d.m.Varint(1) // long micros
   if ok {
      return v, nil
   }
   return 0, fmt.Errorf("details, micros")
}

func (d Details) Num_Downloads() (uint64, error) {
   // DocDetails.DocumentDetails details
   d.m, _ = d.m.Message(13)
   // AppDetails appDetails
   d.m, _ = d.m.Message(1)
   // I dont know the name of field 70, but the similar field 13 is called
   // numDownloads
   v, ok := d.m.Varint(70)
   if ok {
      return v, nil
   }
   return 0, fmt.Errorf("details, num downloads")
}

func (d Details) Title() (string, error) {
   v, ok := d.m.String(5) // String title
   if ok {
      return v, nil
   }
   return "", fmt.Errorf("details, title")
}

func (d Details) Upload_Date() (string, error) {
   d.m, _ = d.m.Message(13) // DocDetails.DocumentDetails details
   d.m, _ = d.m.Message(1) // AppDetails appDetails
   v, ok := d.m.String(16) // String uploadDate
   if ok {
      return v, nil
   }
   return "", fmt.Errorf("details, upload date")
}

func (d Details) Version() (string, error) {
   d.m, _ = d.m.Message(13) // DocDetails.DocumentDetails details
   d.m, _ = d.m.Message(1) // AppDetails appDetails
   v, ok := d.m.String(4) // String versionString
   if ok {
      return v, nil
   }
   return "", fmt.Errorf("details, version")
}

func (d Details) Version_Code() (uint64, error) {
   d.m, _ = d.m.Message(13) // DocDetails.DocumentDetails details
   d.m, _ = d.m.Message(1) // AppDetails appDetails
   v, ok := d.m.Varint(3) // int versionCode
   if ok {
      return v, nil
   }
   return 0, fmt.Errorf("details, version code")
}

// FileMetadata
// This is similar to AppFileMetadata, but notably field 4 is different.
type File_Metadata struct {
   m protobuf.Message
}

func (f File_Metadata) File_Type() (uint64, error) {
   v, ok := f.m.Varint(1) // int fileType
   if ok {
      return v, nil
   }
   return 0, fmt.Errorf("file metadata, file type")
}
type Split_Delivery_Data struct {
   m protobuf.Message
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

type App_File_Metadata struct {
   m protobuf.Message
}

// AndroidAppDeliveryData
type Delivery struct {
   m protobuf.Message
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

func (d Delivery) Download_URL() (string, error) {
   v, ok := d.m.String(3)
   if ok {
      return v, nil
   }
   return "", errors.New("delivery, download URL")
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

type Header struct {
   Agent func() (string, string)
   Authorization func() (string, string)
   Device_Config func() (string, string)
   Device_ID func() (string, string)
}

func (h *Header) Set_Device(device []byte) error {
   var (
      dev Device
      err error
   )
   dev.m, err = protobuf.Consume(device)
   if err != nil {
      return err
   }
   id, err := dev.android_ID()
   if err != nil {
      return err
   }
   h.Device_ID = func() (string, string) {
      return "X-DFE-Device-ID", strconv.FormatUint(id, 16)
   }
   h.Device_Config = func() (string, string) {
      id := strconv.FormatUint(id, 10)
      date := strconv.FormatInt(time.Now().UnixMicro(), 10)
      b := protobuf.Message{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 3, Type: 2, Value: protobuf.Prefix{
               protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Bytes(id)},
               protobuf.Field{Number: 2, Type: 2, Value: protobuf.Prefix{
                  protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Bytes(date)},
               }},
            }},
         }},
      }.Append(nil)
      return "X-DFE-Device-Config-Token", base64.StdEncoding.EncodeToString(b)
   }
   return nil
}

func parse_query(query string) (map[string]string, error) {
   values := make(map[string]string)
   for query != "" {
      var line string
      line, query, _ = strings.Cut(query, "\n")
      key, value, _ := strings.Cut(line, "=")
      var err error
      key, err = url.QueryUnescape(key)
      if err != nil {
         return nil, err
      }
      value, err = url.QueryUnescape(value)
      if err != nil {
         return nil, err
      }
      values[key] = value
   }
   return values, nil
}

type Access_Token map[string]string

func (a Access_Token) auth() string {
   return a["Auth"]
}

func (h *Header) Set_Agent(single bool) {
   var b []byte
   // `sdk` is needed for `/fdfe/delivery`
   b = append(b, "Android-Finsky (sdk="...)
   // valid range 0 - 0x7FFF_FFFF
   b = strconv.AppendInt(b, 9, 10)
   // com.android.vending
   b = append(b, ",versionCode="...)
   if single {
      // valid range 8_03_2_00_00 - 8_09_1_99_99
      b = strconv.AppendInt(b, 8_09_1_99_99, 10)
   } else {
      // valid range 8_09_2_00_00 - math.MaxInt32
      b = strconv.AppendInt(b, 9_99_9_99_99, 10)
   }
   b = append(b, ')')
   h.Agent = func() (string, string) {
      return "User-Agent", string(b)
   }
}

func (r Refresh_Token) token() string {
   return r["Token"]
}

type Refresh_Token map[string]string
