package play

import (
   "154.pages.dev/protobuf"
   "bytes"
   "errors"
   "net/http"
)

func (h Header) Upload(c Config, platform string) error {
   var m protobuf.Message
   m.Add(1, func(m *protobuf.Message) {
      //protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(3)},
      m.Add_Varint(1, c.Touch_Screen)
      //protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1)},
      m.Add_Varint(2, c.Keyboard)
      //protobuf.Field{Number: 3, Type: 0, Value: protobuf.Varint(1)},
      m.Add_Varint(3, c.Navigation)
      //protobuf.Field{Number: 4, Type: 0, Value: protobuf.Varint(2)},
      m.Add_Varint(4, c.Screen_Layout)
      //protobuf.Field{Number: 5, Type: 0, Value: protobuf.Varint(0)},
      m.Add_Bool(5, c.Has_Hard_Keyboard)
      //protobuf.Field{Number: 6, Type: 0, Value: protobuf.Varint(0)},
      m.Add_Bool(6, c.Has_Five_Way_Navigation)
      //protobuf.Field{Number: 7, Type: 0, Value: protobuf.Varint(490)},
      m.Add_Varint(7, c.Screen_Density)
      //protobuf.Field{Number: 8, Type: 0, Value: protobuf.Varint(196610)},
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
   r, _ := http.NewRequest(
      "POST",
      "https://android.clients.google.com/fdfe/uploadDeviceConfig",
      bytes.NewReader(m.Append(nil)),
   )
   r.Header.Set("User-Agent", "Android-Finsky/15.8.23-all [0] [PR] 259261889 (api=3,versionCode=81582300,sdk=28,device=sargo,hardware=sargo,product=sargo,platformVersionRelease=9,model=Pixel 3a,buildId=PQ3B.190705.003,isWideScreen=0,supportedAbis=arm64-v8a;armeabi-v7a;armeabi)")
   // seems like we need this, what the fuck:
   r.Header.Set("X-DFE-Client-Id", "am-android-google")
   r.Header.Set(h.Authorization())
   r.Header.Set(h.Device_ID())
   res, err := http.DefaultClient.Do(r)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   return nil
}
