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

func encode_base64(p []byte) ([]byte, error) {
   var b bytes.Buffer
   w := base64.NewEncoder(base64.URLEncoding, &b)
   if _, err := w.Write(p); err != nil {
      return nil, err
   }
   if err := w.Close(); err != nil {
      return nil, err
   }
   return b.Bytes(), nil
}

func (g GoogleAuth) authorization(req *http.Request) {
   req.Header.Set("authorization", "Bearer " + g.GetAuth())
}

type Checkin18 struct {
   Field9 []string
   Field11 string
   Field15 []string
   Field26 []string
}

var Phone = Checkin18{
   Field9: []string{
      // com.amctve.amcfullepisodes
      "org.apache.http.legacy",
      // com.binance.dev
      "android.test.runner",
   },
   Field15: []string{
      // com.instagram.android
      "GL_OES_compressed_ETC1_RGB8_texture",
      // com.kakaogames.twodin
      "GL_KHR_texture_compression_astc_ldr",
   },
   Field26: []string{
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
}

const google_play_store = 82941300

// the device actually uses 0x30000, but some apps require a higher version:
// com.axis.drawingdesk.v3
// so lets lie for now
const gl_es_version = 0x30001

const android_api = 31

func (g GoogleCheckin) x_dfe_device_id(req *http.Request) error {
   id, err := g.device_id()
   if err != nil {
      return err
   }
   req.Header.Set("x-dfe-device-id", fmt.Sprintf("%x", id))
   return nil
}

func (g GoogleCheckin) x_ps_rh(req *http.Request) error {
   id, err := g.device_id()
   if err != nil {
      return err
   }
   field1, err := func() ([]byte, error) {
      var m protobuf.Message
      m.Add(3, func(m *protobuf.Message) {
         m.AddBytes(1, fmt.Append(nil, id))
         m.Add(2, func(m *protobuf.Message) {
            v := time.Now().UnixMicro()
            m.AddBytes(1, fmt.Append(nil, v))
         })
      })
      b, err := compress_gzip(m.Encode())
      if err != nil {
         return nil, err
      }
      return encode_base64(b)
   }()
   if err != nil {
      return err
   }
   ps_rh, err := func() ([]byte, error) {
      var m protobuf.Message
      m.Add(1, func(m *protobuf.Message) {
         m.AddBytes(1, field1)
      })
      b, err := compress_gzip(m.Encode())
      if err != nil {
         return nil, err
      }
      return encode_base64(b)
   }()
   if err != nil {
      return err
   }
   req.Header.Set("x-ps-rh", string(ps_rh))
   return nil
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

var ABIs = map[int]string{
   // com.google.android.youtube
   0: "x86",
   // com.sygic.aura
   1: "armeabi-v7a",
   // com.kakaogames.twodin
   2: "arm64-v8a",
}

// play.google.com/store/apps/details?id=com.google.android.apps.youtube.unplugged
type StoreApp struct {
   ID string
   Version uint64
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

func (s StoreApp) APK(v string) string {
   var b []byte
   b = fmt.Append(b, s.ID, "-")
   if v != "" {
      b = fmt.Append(b, v, "-")
   }
   b = fmt.Append(b, s.Version, ".apk")
   return string(b)
}
