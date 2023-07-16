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
      4: protobuf.Message{ // checkin
         // Logs$AndroidCheckinProto
         1: protobuf.Message{ // build
            // Logs$AndroidBuildProto
            // multiple APK valid range 14 - 0x7FFF_FFFF
            // single APK valid range 14 - 28
            10: protobuf.Varint(28), // sdkVersion
         },
         18: protobuf.Varint(1), // voiceCapable
      },
      // valid range 2 - 3
      14: protobuf.Varint(3), // version
      18: protobuf.Message{ // deviceConfiguration
         // DeviceConfiguration
         1: protobuf.Varint(c.Touch_Screen),
         2: protobuf.Varint(c.Keyboard),
         3: protobuf.Varint(c.Navigation),
         4: protobuf.Varint(c.Screen_Layout),
         5: protobuf.Varint(c.Has_Hard_Keyboard),
         6: protobuf.Varint(c.Has_Five_Way_Navigation),
         7: protobuf.Varint(c.Screen_Density),
         8: protobuf.Varint(c.GL_ES_Version),
         11: protobuf.String(platform), // nativePlatform
      },
   }
   for _, library := range c.System_Shared_Library {
      // .deviceConfiguration.systemSharedLibrary
      body.Get(18).Add_String(9, library)
   }
   for _, extension := range c.GL_Extension {
      // .deviceConfiguration.glExtension
      body.Get(18).Add_String(15, extension)
   }
   for _, name := range c.New_System_Available_Feature {
      // .deviceConfiguration.newSystemAvailableFeature
      body.Get(18).Add(26, protobuf.Message{
         1: protobuf.String(name),
      })
   }
   res, err := client.Post(
      "https://android.googleapis.com/checkin",
      "application/x-protobuffer",
      bytes.NewReader(body.Marshal()),
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
