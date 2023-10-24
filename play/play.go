package play

import (
   "154.pages.dev/protobuf"
   "encoding/base64"
   "strconv"
   "time"
)

const (
   google_play_store = 82941300
   // WARNING 28 is the last version that supports single APK
   google_services_framework = 28
)

func (h *Header) Set_Agent(single bool) {
   var b []byte
   // `sdk` is needed for `/fdfe/delivery`
   b = append(b, "Android-Finsky (sdk="...)
   // with `/fdfe/acquire`, requests will be rejected with certain apps, if the
   // device was created with too low a version here:
   b = strconv.AppendInt(b, google_services_framework, 10)
   b = append(b, ",versionCode="...)
   // for multiple APKs just tell the truth. for single APK we have to lie.
   // below value is the last version that works.
   if single {
      b = strconv.AppendInt(b, 80919999, 10)
   } else {
      b = strconv.AppendInt(b, google_play_store, 10)
   }
   b = append(b, ')')
   h.Agent = func() (string, string) {
      return "User-Agent", string(b)
   }
}

func (h *Header) Set_Device(device []byte) error {
   var (
      check Checkin
      err   error
   )
   check.m, err = protobuf.Consume(device)
   if err != nil {
      return err
   }
   id, err := check.device_ID()
   if err != nil {
      return err
   }
   h.Device_Config = func() (string, string) {
      var m protobuf.Message
      m.Add(1, func(m *protobuf.Message) {
         m.Add(3, func(m *protobuf.Message) {
            m.Add_String(1, strconv.FormatUint(id, 10))
            m.Add(2, func(m *protobuf.Message) {
               now := time.Now().UnixMicro()
               m.Add_String(1, strconv.FormatInt(now, 10))
            })
         })
      })
      token := base64.StdEncoding.EncodeToString(m.Append(nil))
      return "X-DFE-Device-Config-Token", token
   }
   h.Device_ID = func() (string, string) {
      return "X-DFE-Device-ID", strconv.FormatUint(id, 16)
   }
   return nil
}

var Platforms = map[int64]string{
   // com.google.android.youtube
   0: "x86",
   // com.miui.weather2
   1: "armeabi-v7a",
   // com.kakaogames.twodin
   2: "arm64-v8a",
}

type Device struct {
   // developer.android.com/ndk/guides/abis
   Platform string
   // developer.android.com/guide/topics/manifest/supports-gl-texture-element
   Texture []string
   // developer.android.com/guide/topics/manifest/uses-library-element
   Library []string
   // developer.android.com/guide/topics/manifest/uses-feature-element
   Feature []string
}

var Phone = Device{
   Texture: []string{
      // com.instagram.android
      "GL_OES_compressed_ETC1_RGB8_texture",
      // com.kakaogames.twodin
      "GL_KHR_texture_compression_astc_ldr",
   },
   Library: []string{
      // com.amctve.amcfullepisodes
      "org.apache.http.legacy",
      // com.binance.dev
      "android.test.runner",
   },
   Feature: []string{
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
}

type Header struct {
   Agent         func() (string, string)
   Authorization func() (string, string)
   Device_Config func() (string, string)
   Device_ID     func() (string, string)
}
