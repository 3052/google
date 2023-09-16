package acquire

import (
   "154.pages.dev/encoding/protobuf"
   "bytes"
   "net/http"
)

func (client *_GooglePlayClient) uploadDeviceConfig() error {
   c := Phone
   c.Native_Platform = "x86"
   // class UploadDeviceConfigRequest
   var m protobuf.Message
   // DeviceConfiguration.DeviceConfigurationProto deviceConfiguration
   m.Add(1, func(m *protobuf.Message) {
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
   b := m.Append(nil)
   r, _ := http.NewRequest("POST", _UrlUploadDeviceConfig, bytes.NewReader(b))
   r.Header.Set("User-Agent", "Android-Finsky/15.8.23-all [0] [PR] 259261889 (api=3,versionCode=81582300,sdk=28,device=sargo,hardware=sargo,product=sargo,platformVersionRelease=9,model=Pixel 3a,buildId=PQ3B.190705.003,isWideScreen=0,supportedAbis=arm64-v8a;armeabi-v7a;armeabi)")
   r.Header.Set("X-DFE-Device-Id", client.AuthData.GsfID)
   // seems like we need this, what the fuck:
   r.Header.Set("X-DFE-Client-Id", "am-android-google")
   _, status, err := doReq(r)
   if err != nil {
      return err
   }
   if status == 401 {
      return err_GPTokenExpired
   }
   return nil
}
