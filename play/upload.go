package play

import (
   "154.pages.dev/encoding/protobuf"
   "bytes"
   "errors"
   "net/http"
)

func (c Config) UploadDeviceConfig(h Header, platform string) error {
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
   r.Header.Set(h.Device())
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
   // valid range 0x3_0001 - 0x7FFF_FFFF
   GL_ES_Version: 0xF_FFFF,
}
