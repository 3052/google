package play

import (
   "154.pages.dev/encoding/protobuf"
   "bytes"
   "errors"
   "fmt"
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "strconv"
   "strings"
)

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

// Checkin$AndroidCheckinResponse
type Device struct {
   m protobuf.Message
}

// androidId
func (d Device) ID() (uint64, error) {
   return d.m.Fixed64(7)
}

func (client *_GooglePlayClient) checkIn() (*checkin, error) {
   b := checkin_body.Append(nil)
   r, _ := http.NewRequest("POST", _UrlCheckIn, bytes.NewReader(b))
   r.Header.Set("User-Agent", "GoogleAuth/1.4 sargo PQ3B.190705.003")
   r.Header.Set("Content-Type", "application/x-protobuffer")
   r.Header.Set("Host", "android.clients.google.com")
   b, _, err := doReq(r)
   if err != nil {
      return nil, err
   }
   var check checkin
   check.m, err = protobuf.Consume(b)
   if err != nil {
      return nil, err
   }
   return &check, nil
}

var checkin_body = protobuf.Message{
   protobuf.Field{Number: 4, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("PQ3B.190705.003")},
         protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("sargo")},
         protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("google")},
         protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("g670-00011-190411-B-5457439")},
         protobuf.Field{Number: 5, Type: 2, Value: protobuf.Bytes("b4s4-0.1-5613380")},
         protobuf.Field{Number: 6, Type: 2, Value: protobuf.Bytes("android-google")},
         protobuf.Field{Number: 7, Type: 0, Value: protobuf.Varint(1694054582)},
         protobuf.Field{Number: 8, Type: 0, Value: protobuf.Varint(203615028)},
         protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("sargo")},
         protobuf.Field{Number: 10, Type: 0, Value: protobuf.Varint(28)},
         protobuf.Field{Number: 11, Type: 2, Value: protobuf.Bytes("Pixel 3a")},
         protobuf.Field{Number: 12, Type: 2, Value: protobuf.Bytes("google")},
         protobuf.Field{Number: 13, Type: 2, Value: protobuf.Bytes("sargo")},
         protobuf.Field{Number: 14, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      protobuf.Field{Number: 6, Type: 2, Value: protobuf.Bytes("334050")},
      protobuf.Field{Number: 7, Type: 2, Value: protobuf.Bytes("20815")},
      protobuf.Field{Number: 8, Type: 2, Value: protobuf.Bytes("mobile-notroaming")},
      protobuf.Field{Number: 9, Type: 0, Value: protobuf.Varint(0)},
   }},
   protobuf.Field{Number: 14, Type: 0, Value: protobuf.Varint(3)},
}

// A Sleep is needed after this.
func (c Config) Checkin(platform string) ([]byte, error) {
   var m protobuf.Message
   m.Add(4, func(m *protobuf.Message) { // checkin
      m.Add(1, func(m *protobuf.Message) { // build
         // int sdkVersion_
         // single APK valid range 14 - 28
         // multiple APK valid range 14 - 0x7FFF_FFFF
         m.Add_Varint(10, 28)
      })
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
