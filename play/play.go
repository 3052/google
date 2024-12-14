package play

import (
   "41.neocities.org/protobuf"
   "bytes"
   "compress/gzip"
   "encoding/base64"
   "net/http"
   "time"
   "strconv"
)

// com.roku.web.trc
const Leanback = "android.software.leanback"

// the device actually uses 0x30000, but some apps require a higher version:
// com.axis.drawingdesk.v3
// so lets lie for now
const gl_es_version = 0x30001

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

func user_agent(req *http.Request, single bool) {
   // `sdk` is needed for `/fdfe/delivery`
   data := []byte("Android-Finsky (sdk=")
   // with `/fdfe/acquire`, requests will be rejected with certain apps, if the
   // device was created with too low a version here:
   data = strconv.AppendInt(data, android_api, 10)
   data = append(data, ",versionCode="...)
   // for multiple APKs just tell the truth. for single APK we have to lie.
   // below value is the last version that works.
   if single {
      data = strconv.AppendInt(data, 80919999, 10)
   } else {
      data = strconv.AppendInt(data, google_play_store, 10)
   }
   data = append(data, ')')
   req.Header.Set("user-agent", string(data))
}

const google_play_store = 82941300

const android_api = 31

func compress_gzip(in []byte) ([]byte, error) {
   var out bytes.Buffer
   w := gzip.NewWriter(&out)
   _, err := w.Write(in)
   if err != nil {
      return nil, err
   }
   err = w.Close()
   if err != nil {
      return nil, err
   }
   return out.Bytes(), nil
}

type GoogleDevice struct {
   Abi     string
   Feature []string
   Library []string
   Texture []string
}

func authorization(req *http.Request, auth GoogleAuth) {
   req.Header.Set("authorization", "Bearer "+auth.auth())
}

func (s *StoreApp) Apk(value string) string {
   data := []byte(s.Id)
   data = append(data, '-')
   if value != "" {
      data = append(data, value...)
      data = append(data, '-')
   }
   data = strconv.AppendUint(data, s.Version, 10)
   data = append(data, ".apk"...)
   return string(data)
}

func (s *StoreApp) Obb(value uint64) string {
   var data []byte
   if value >= 1 {
      data = append(data, "patch."...)
   } else {
      data = append(data, "main."...)
   }
   data = strconv.AppendUint(data, s.Version, 10)
   data = append(data, '.')
   data = append(data, s.Id...)
   data = append(data, ".obb"...)
   return string(data)
}

// play.google.com/store/apps/details?id=com.google.android.apps.youtube.unplugged
type StoreApp struct {
   Id      string
   Version uint64
}

func x_dfe_device_id(req *http.Request, check GoogleCheckin) {
   req.Header.Set("x-dfe-device-id", strconv.FormatUint(check.field_7(), 16))
}

func x_ps_rh(req *http.Request, check GoogleCheckin) error {
   id := strconv.FormatUint(check.field_7(), 10)
   now := strconv.FormatInt(time.Now().UnixMicro(), 10)
   message := protobuf.Message{
      1: {protobuf.Message{
         1: {protobuf.Message{
            3: {protobuf.Message{
               1: {protobuf.Bytes(id)},
               2: {protobuf.Message{
                  1: {protobuf.Bytes(now)},
               }},
            }},
         }},
      }},
   }
   data, err := compress_gzip(message.Marshal())
   if err != nil {
      return err
   }
   req.Header.Set("x-ps-rh", base64.URLEncoding.EncodeToString(data))
   return nil
}

// developer.android.com/ndk/guides/abis
var Abis = []string{
   // com.google.android.youtube
   "x86",
   "x86_64",
   // com.sygic.aura
   "armeabi-v7a",
   // com.kakaogames.twodin
   "arm64-v8a",
}
