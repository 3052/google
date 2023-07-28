package play

import (
   "154.pages.dev/encoding/protobuf"
   "154.pages.dev/tls"
   "bytes"
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
)

// A Sleep is needed after this.
func (c Config) Checkin(platform string) (*Response, error) {
   body := protobuf.Message{
      // Checkin$AndroidCheckinRequest
      protobuf.Number(4).Prefix( // checkin
         // Logs$AndroidCheckinProto
         protobuf.Number(1).Prefix( // build
            // Logs$AndroidBuildProto
            // multiple APK valid range 14 - 0x7FFF_FFFF
            // single APK valid range 14 - 28
            protobuf.Number(10).Varint(28), // sdkVersion
         ),
         protobuf.Number(18).Varint(1), // voiceCapable
      ),
      // valid range 2 - 3
      protobuf.Number(14).Varint(3), // version
      protobuf.Number(18).Prefix( // deviceConfiguration
         // DeviceConfiguration
         protobuf.Number(1).Uvarint(c.Touch_Screen),
         protobuf.Number(2).Uvarint(c.Keyboard),
         protobuf.Number(3).Uvarint(c.Navigation),
         protobuf.Number(4).Uvarint(c.Screen_Layout),
         protobuf.Number(5).Uvarint(c.Has_Hard_Keyboard),
         protobuf.Number(6).Uvarint(c.Has_Five_Way_Navigation),
         protobuf.Number(7).Uvarint(c.Screen_Density),
         protobuf.Number(8).Uvarint(c.GL_ES_Version),
         func() (m protobuf.Message) {
            for _, library := range c.System_Shared_Library {
               // systemSharedLibrary
               m = append(m, protobuf.Number(9).String(library))
            }
            return
         }(),
         protobuf.Number(11).String(platform), // nativePlatform
         func() (m protobuf.Message) {
            for _, extension := range c.GL_Extension {
               // glExtension
               m = append(m, protobuf.Number(15).String(extension))
            }
            return
         }(),
         func() (m protobuf.Message) {
            for _, name := range c.New_System_Available_Feature {
               // newSystemAvailableFeature
               m = append(m, protobuf.Number(26).Prefix(
                  protobuf.Number(1).String(name),
               ))
            }
            return
         }(),
      ),
   }
   res, err := client.Post(
      "https://android.googleapis.com/checkin", "application/x-protobuffer",
      bytes.NewReader(body.Append(nil)),
   )
   if err != nil {
      return nil, err
   }
   return &Response{res}, nil
}

func (r Response) Write_File(name string) error {
   data, err := io.ReadAll(r.Body)
   if err != nil {
      return err
   }
   return os.WriteFile(name, data, 0666)
}

type Response struct {
   *http.Response
}

// You can also use host "android.clients.google.com", but it also uses
// TLS fingerprinting.
func New_Auth(email, passwd string) (*Response, error) {
   body := url.Values{
      "Email": {email},
      "Passwd": {passwd},
      "client_sig": {""},
      "droidguard_results": {"-"},
   }.Encode()
   req, err := http.NewRequest(
      "POST", "https://android.googleapis.com/auth", strings.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   client := http.DefaultClient
   client.Transport = &tls.Transport{Spec: tls.Android_API_26}
   res, err := client.Do(req)
   if err != nil {
      return nil, err
   }
   return &Response{res}, nil
}
