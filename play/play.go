package play

import (
   "154.pages.dev/protobuf"
   "bytes"
   "compress/gzip"
   "encoding/base64"
   "fmt"
   "net/http"
   "time"
)

const google_play_store = 82941300

func user_agent(req *http.Request, single bool) {
   var b []byte
   // `sdk` is needed for `/fdfe/delivery`
   b = append(b, "Android-Finsky (sdk="...)
   // with `/fdfe/acquire`, requests will be rejected with certain apps, if the
   // device was created with too low a version here:
   b = fmt.Append(b, android_api)
   b = append(b, ",versionCode="...)
   // for multiple APKs just tell the truth. for single APK we have to lie.
   // below value is the last version that works.
   if single {
      b = fmt.Append(b, 80919999)
   } else {
      b = fmt.Append(b, google_play_store)
   }
   b = append(b, ')')
   req.Header.Set("user-agent", string(b))
}

// developer.android.com/ndk/guides/abis
var Abi = []string{
   // com.google.android.youtube
   "x86",
   // com.sygic.aura
   "armeabi-v7a",
   // com.kakaogames.twodin
   "arm64-v8a",
}

var Device = GoogleDevice{
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
      // com.pinterest
      "android.hardware.camera",
      "android.hardware.location",
      "android.hardware.screen.portrait",
      // com.roku.web.trc
      "android.hardware.screen.landscape",
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
   },
   Library: []string{
      // com.amctve.amcfullepisodes
      "org.apache.http.legacy",
      // com.binance.dev
      "android.test.runner",
   },
   Texture: []string{
      // com.instagram.android
      "GL_OES_compressed_ETC1_RGB8_texture",
      // com.kakaogames.twodin
      "GL_KHR_texture_compression_astc_ldr",
   },
}

func compress_gzip(data []byte) ([]byte, error) {
   var buf bytes.Buffer
   w := gzip.NewWriter(&buf)
   _, err := w.Write(data)
   if err != nil {
      return nil, err
   }
   err = w.Close()
   if err != nil {
      return nil, err
   }
   return buf.Bytes(), nil
}

type GoogleDevice struct {
   Abi     string
   Feature []string
   Library []string
   Texture []string
}

func (s *StoreApp) Apk(value string) string {
   var b []byte
   b = fmt.Append(b, s.Id, "-")
   if value != "" {
      b = fmt.Append(b, value, "-")
   }
   b = fmt.Append(b, s.Version, ".apk")
   return string(b)
}

// com.roku.web.trc
const Leanback = "android.software.leanback"

const android_api = 31

// play.google.com/store/apps/details?id=com.google.android.apps.youtube.unplugged
type StoreApp struct {
   Id      string
   Version uint64
}

// the device actually uses 0x30000, but some apps require a higher version:
// com.axis.drawingdesk.v3
// so lets lie for now
const gl_es_version = 0x30001

func (s *StoreApp) Obb(value uint64) string {
   var b []byte
   if value >= 1 {
      b = append(b, "patch"...)
   } else {
      b = append(b, "main"...)
   }
   b = fmt.Append(b, ".", s.Version, ".", s.Id, ".obb")
   return string(b)
}

func authorization(req *http.Request, auth *GoogleAuth) {
   req.Header.Set("authorization", "Bearer "+auth.auth())
}

func x_dfe_device_id(req *http.Request, checkin *GoogleCheckin) error {
   id, err := checkin.device_id()
   if err != nil {
      return err
   }
   req.Header.Set("x-dfe-device-id", fmt.Sprintf("%x", id))
   return nil
}

func x_ps_rh(req *http.Request, checkin *GoogleCheckin) error {
   id, err := checkin.device_id()
   if err != nil {
      return err
   }
   message := protobuf.Message{}
   message.Add(1, func(m protobuf.Message) {
      m.Add(1, func(m protobuf.Message) {
         m.Add(3, func(m protobuf.Message) {
            m.AddBytes(1, fmt.Append(nil, id))
            m.Add(2, func(m protobuf.Message) {
               now := time.Now().UnixMicro()
               m.AddBytes(1, fmt.Append(nil, now))
            })
         })
      })
   })
   data, err := compress_gzip(message.Marshal())
   if err != nil {
      return err
   }
   req.Header.Set("x-ps-rh", base64.URLEncoding.EncodeToString(data))
   return nil
}
