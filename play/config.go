package play

import (
   "154.pages.dev/encoding/protobuf"
   "bytes"
   "io"
   "net/http"
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

// A Sleep is needed after this.
func (c Config) Checkin(platform string) (Raw_Device, error) {
   var m protobuf.Message
   m.Add(4, func(m *protobuf.Message) { // checkin
      m.Add(1, func(m *protobuf.Message) { // build
         // int sdkVersion_
         // single APK valid range 14 - 28
         // multiple APK valid range 14 - 0x7FFF_FFFF
         m.Add_Varint(10, 28)
      })
      m.Add_Varint(18, 1) // voiceCapable?
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

type Raw_Device []byte

// Checkin$AndroidCheckinResponse
type Device struct {
   m protobuf.Message
}

func New_Device(r Raw_Device) (*Device, error) {
   var dev Device
   if err := dev.UnmarshalBinary(r); err != nil {
      return nil, err
   }
   return &dev, nil
}

func (d Device) MarshalBinary() ([]byte, error) {
   return d.m.Append(nil), nil
}

func (d *Device) UnmarshalBinary(data []byte) error {
   var err error
   d.m, err = protobuf.Consume(data)
   if err != nil {
      return err
   }
   return nil
}

// androidId
func (d Device) ID() (uint64, error) {
   return d.m.Fixed64(7)
}
