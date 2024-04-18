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

func (g GoogleCheckin) x_ps_rh(req *http.Request) error {
   id, err := g.device_id()
   if err != nil {
      return err
   }
   var m protobuf.Message
   m.Add(1, func(m *protobuf.Message) {
      m.Add(1, func(m *protobuf.Message) {
         m.Add(3, func(m *protobuf.Message) {
            m.AddBytes(1, fmt.Append(nil, id))
            m.Add(2, func(m *protobuf.Message) {
               now := time.Now().UnixMicro()
               m.AddBytes(1, fmt.Append(nil, now))
            })
         })
      })
   })
   data, err := compress_gzip(m.Encode())
   if err != nil {
      return err
   }
   req.Header.Set("x-ps-rh", base64.URLEncoding.EncodeToString(data))
   return nil
}

func compress_gzip(p []byte) ([]byte, error) {
   var b bytes.Buffer
   w := gzip.NewWriter(&b)
   if _, err := w.Write(p); err != nil {
      return nil, err
   }
   if err := w.Close(); err != nil {
      return nil, err
   }
   return b.Bytes(), nil
}

const android_api = 31

// the device actually uses 0x30000, but some apps require a higher version:
// com.axis.drawingdesk.v3
// so lets lie for now
const gl_es_version = 0x30001

const google_play_store = 82941300

var ABIs = map[int]string{
   // com.google.android.youtube
   0: "x86",
   // com.sygic.aura
   1: "armeabi-v7a",
   // com.kakaogames.twodin
   2: "arm64-v8a",
}

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
   req.Header.Set("User-Agent", string(b))
}

// developer.android.com/ndk/guides/abis
type ABI int

func (a *ABI) Set(s string) error {
   _, err := fmt.Sscan(s, a)
   if err != nil {
      return err
   }
   return nil
}

func (a ABI) String() string {
   return ABIs[int(a)]
}

func (g GoogleAuth) authorization(req *http.Request) {
   req.Header.Set("authorization", "Bearer " + g.get_auth())
}

func (g GoogleCheckin) x_dfe_device_id(req *http.Request) error {
   id, err := g.device_id()
   if err != nil {
      return err
   }
   req.Header.Set("x-dfe-device-id", fmt.Sprintf("%x", id))
   return nil
}

type GoogleDevice struct {
   ABI string
   Feature []string
   Library []string
   Texture []string
}

var Phone = GoogleDevice{
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

// play.google.com/store/apps/details?id=com.google.android.apps.youtube.unplugged
type StoreApp struct {
   ID string
   Version uint64
}

func (s StoreApp) APK(v string) string {
   var b []byte
   b = fmt.Append(b, s.ID, "-")
   if v != "" {
      b = fmt.Append(b, v, "-")
   }
   b = fmt.Append(b, s.Version, ".apk")
   return string(b)
}

func (s StoreApp) OBB(v uint64) string {
   var b []byte
   if v >= 1 {
      b = append(b, "patch"...)
   } else {
      b = append(b, "main"...)
   }
   b = fmt.Append(b, ".", s.Version, ".", s.ID, ".obb")
   return string(b)
}
