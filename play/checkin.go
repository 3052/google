package play

import (
   "154.pages.dev/encoding/protobuf"
   "bytes"
   "errors"
   "io"
   "net/http"
)

func (c Config) Checkin() ([]byte, error) {
   // Checkin$AndroidCheckinRequest
   var m protobuf.Message
   // Logs$AndroidCheckinProto checkin_
   m.Add(4, func(m *protobuf.Message) {
      // Logs$AndroidBuildProto build_
      m.Add(1, func(m *protobuf.Message) {
         // int sdkVersion_
         // single APK valid range 14 - 28
         // multiple APK valid range 14 - math.MaxInt32
         m.Add_Varint(10, 28)
      })
      m.Add_Varint(18, 1)
   })
   // int version
   // valid range 2 - 3
   m.Add_Varint(14, 3)
   // Config$DeviceConfigurationProto deviceConfiguration_
   m.Add(18, func(m *protobuf.Message) {
      // int touchScreen
      m.Add_Varint(1, c.Touch_Screen)
      // int keyboard
      m.Add_Varint(2, c.Keyboard)
      // int navigation
      m.Add_Varint(3, c.Navigation)
      // int screenLayout
      m.Add_Varint(4, c.Screen_Layout)
      // boolean hasHardKeyboard
      m.Add_Bool(5, c.Has_Hard_Keyboard)
      // boolean hasFiveWayNavigation
      m.Add_Bool(6, c.Has_Five_Way_Navigation)
      // int screenDensity
      m.Add_Varint(7, c.Screen_Density)
      // int glEsVersion
      m.Add_Varint(8, c.GL_ES_Version)
      for _, library := range c.System_Shared_Library {
         // String[] systemSharedLibrary
         m.Add_String(9, library)
      }
      // String[] nativePlatform
      m.Add_String(11, c.Native_Platform)
      for _, extension := range c.GL_Extension {
         // String[] glExtension
         m.Add_String(15, extension)
      }
      // you cannot swap the next two lines:
      for _, feature := range c.System_Available_Feature {
         m.Add(26, func(m *protobuf.Message) {
            // String[] systemAvailableFeature
            m.Add_String(1, feature)
         })
      }
   })
   req, err := http.NewRequest(
      "POST", "https://android.googleapis.com/checkin",
      bytes.NewReader(m.Append(nil)),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Content-Type": {"application/x-protobuffer"},
      "User-Agent": {"GoogleAuth/1.4 sargo PQ3B.190705.003"},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   return io.ReadAll(res.Body)
}

func (d Device) android_ID() (uint64, error) {
   v, ok := d.m.Fixed64(7) // long androidId_
   if ok {
      return v, nil
   }
   return 0, errors.New("androidId_")
}

// Checkin$AndroidCheckinResponse
type Device struct {
   m protobuf.Message
}
