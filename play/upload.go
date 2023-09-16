package play

import (
   "154.pages.dev/encoding/protobuf"
   "bytes"
   "errors"
   "net/http"
)

func (h Header) Upload(c Config) error {
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
   req, err := http.NewRequest(
      "POST", "https://android.clients.google.com/fdfe/uploadDeviceConfig",
      bytes.NewReader(m.Append(nil)),
   )
   if err != nil {
      return err
   }
   // seems like we need this, what the fuck:
   req.Header.Set("X-DFE-Client-Id", "am-android-google")
   req.Header.Set(h.Device())
   req.Header.Set(h.Agent())
   //req.Header.Set(h.Authorization())
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   return nil
}
