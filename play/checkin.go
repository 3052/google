package play

import (
   "154.pages.dev/encoding/protobuf"
   "bytes"
   "strconv"
   "time"
)

const Sleep = 4 * time.Second

// These can use default values, but they must all be included
type Config struct {
   GL_ES_Version uint64
   GL_Extension []string
   Has_Five_Way_Navigation uint64
   Has_Hard_Keyboard uint64
   Keyboard uint64
   Navigation uint64
   New_System_Available_Feature []string
   Screen_Density uint64
   Screen_Layout uint64
   System_Shared_Library []string
   Touch_Screen uint64
}

var Phone = Config{
   New_System_Available_Feature: []string{
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
         protobuf.Number(1).Varint(c.Touch_Screen),
         protobuf.Number(2).Varint(c.Keyboard),
         protobuf.Number(3).Varint(c.Navigation),
         protobuf.Number(4).Varint(c.Screen_Layout),
         protobuf.Number(5).Varint(c.Has_Hard_Keyboard),
         protobuf.Number(6).Varint(c.Has_Five_Way_Navigation),
         protobuf.Number(7).Varint(c.Screen_Density),
         protobuf.Number(8).Varint(c.GL_ES_Version),
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

// Checkin$AndroidCheckinResponse
type Device struct {
   m protobuf.Message
}

// androidId
func (d Device) ID() (uint64, error) {
   return d.m.Fixed64(7)
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
