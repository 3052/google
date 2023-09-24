package play

import (
   "154.pages.dev/protobuf"
   "bytes"
   "io"
   "net/http"
)

// A Sleep is needed after this.
func (c Config) Checkin(platform string) ([]byte, error) {
   var m protobuf.Message
   m.Add(4, func(m *protobuf.Message) { // checkin
      m.Add(1, func(m *protobuf.Message) { // build
         //protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("PQ3B.190705.003")},
         //protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("sargo")},
         //protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("google")},
         //protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("g670-00011-190411-B-5457439")},
         //protobuf.Field{Number: 5, Type: 2, Value: protobuf.Bytes("b4s4-0.1-5613380")},
         //protobuf.Field{Number: 6, Type: 2, Value: protobuf.Bytes("android-google")},
         //protobuf.Field{Number: 7, Type: 0, Value: protobuf.Varint(1694054582)},
         //protobuf.Field{Number: 8, Type: 0, Value: protobuf.Varint(203615028)},
         //protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("sargo")},
         
         // int sdkVersion_
         // single APK valid range 14 - 28
         // multiple APK valid range 14 - 0x7FFF_FFFF
         m.Add_Varint(10, 28)
         
         //protobuf.Field{Number: 11, Type: 2, Value: protobuf.Bytes("Pixel 3a")},
         //protobuf.Field{Number: 12, Type: 2, Value: protobuf.Bytes("google")},
         //protobuf.Field{Number: 13, Type: 2, Value: protobuf.Bytes("sargo")},
         //protobuf.Field{Number: 14, Type: 0, Value: protobuf.Varint(0)},
      })
      //protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      //protobuf.Field{Number: 6, Type: 2, Value: protobuf.Bytes("334050")},
      //protobuf.Field{Number: 7, Type: 2, Value: protobuf.Bytes("20815")},
      //protobuf.Field{Number: 8, Type: 2, Value: protobuf.Bytes("mobile-notroaming")},
      //protobuf.Field{Number: 9, Type: 0, Value: protobuf.Varint(0)},
      m.Add_Varint(18, 1)
   })
   // int version
   // valid range 2 - 3
   m.Add_Varint(14, 3)
   m.Add(18, func(m *protobuf.Message) { // deviceConfiguration
      m.Add_Varint(1, c.Touch_Screen)
      m.Add_Varint(2, c.Keyboard)
      m.Add_Varint(3, c.Navigation)
      m.Add_Varint(4, c.Screen_Layout)
      m.Add_Bool(5, c.Has_Hard_Keyboard)
      m.Add_Bool(6, c.Has_Five_Way_Navigation)
      m.Add_Varint(7, c.Screen_Density)
      m.Add_Varint(8, c.GL_ES_Version)
      for _, library := range c.System_Shared_Library {
         m.Add_String(9, library)
      }
      m.Add_String(11, platform)
      for _, extension := range c.GL_Extension {
         m.Add_String(15, extension)
      }
      // you cannot swap the next two lines:
      for _, name := range c.System_Available_Feature {
         m.Add(26, func(m *protobuf.Message) {
            m.Add_String(1, name)
         })
      }
   })
   // r.Header.Set("User-Agent", "GoogleAuth/1.4 sargo PQ3B.190705.003")
   res, err := http.Post(
      "https://android.googleapis.com/checkin", "application/x-protobuffer",
      bytes.NewReader(m.Append(nil)),
   )
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return io.ReadAll(res.Body)
}
