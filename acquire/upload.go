package acquire

import (
   "154.pages.dev/encoding/protobuf"
   "bytes"
   "net/http"
)

func (client *_GooglePlayClient) uploadDeviceConfig() error {
   b := upload_body.Append(nil)
   r, _ := http.NewRequest("POST", _UrlUploadDeviceConfig, bytes.NewReader(b))
   data := client.AuthData
   r.Header.Set("Authorization", "Bearer "+data._AuthToken)
   r.Header.Set("User-Agent", "Android-Finsky/15.8.23-all [0] [PR] 259261889 (api=3,versionCode=81582300,sdk=28,device=sargo,hardware=sargo,product=sargo,platformVersionRelease=9,model=Pixel 3a,buildId=PQ3B.190705.003,isWideScreen=0,supportedAbis=arm64-v8a;armeabi-v7a;armeabi)")
   r.Header.Set("X-DFE-Device-Id", data.GsfID)
   // seems like we need this, what the fuck:
   r.Header.Set("X-DFE-Client-Id", "am-android-google")
   _, status, err := doReq(r)
   if err != nil {
      return err
   }
   if status == 401 {
      return err_GPTokenExpired
   }
   return nil
}

var upload_body = protobuf.Message{
   protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(3)},
      protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1)},
      protobuf.Field{Number: 3, Type: 0, Value: protobuf.Varint(1)},
      protobuf.Field{Number: 4, Type: 0, Value: protobuf.Varint(2)},
      protobuf.Field{Number: 5, Type: 0, Value: protobuf.Varint(0)},
      protobuf.Field{Number: 6, Type: 0, Value: protobuf.Varint(0)},
      protobuf.Field{Number: 7, Type: 0, Value: protobuf.Varint(490)},
      protobuf.Field{Number: 8, Type: 0, Value: protobuf.Varint(196610)},
      // com.amctve.amcfullepisodes
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("org.apache.http.legacy")},
      // com.binance.dev
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("android.test.runner")},
      protobuf.Field{Number: 11, Type: 2, Value: protobuf.Bytes("x86")},
      // com.instagram.android
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_compressed_ETC1_RGB8_texture")},
      // com.kakaogames.twodin
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_KHR_texture_compression_astc_ldr")},
      ///////////////////////////////////////////////////////////////////////////
      protobuf.Field{Number: 16, Type: 0, Value: protobuf.Varint(0)},
      protobuf.Field{Number: 19, Type: 0, Value: protobuf.Varint(0)},
      protobuf.Field{Number: 20, Type: 0, Value: protobuf.Varint(8589935000)},
      protobuf.Field{Number: 21, Type: 0, Value: protobuf.Varint(8)},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.audio.low_latency")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.audio.output")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.audio.pro")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.bluetooth")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.bluetooth_le")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera.any")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera.autofocus")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera.capability.manual_post_processing")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera.capability.manual_sensor")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera.capability.raw")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera.flash")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera.front")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera.level.full")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.faketouch")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.fingerprint")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.location")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.location.gps")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.location.network")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.microphone")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.nfc")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.nfc.any")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.nfc.hce")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.nfc.hcef")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.opengles.aep")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.ram.normal")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.screen.landscape")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.screen.portrait")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.accelerometer")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.assist")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.barometer")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.compass")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.gyroscope")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.hifi_sensors")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.light")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.proximity")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.stepcounter")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.stepdetector")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.strongbox_keystore")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.telephony")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.telephony.carrierlock")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.telephony.cdma")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.telephony.euicc")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.telephony.gsm")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen.multitouch")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen.multitouch.distinct")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen.multitouch.jazzhand")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.usb.accessory")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.usb.host")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.vulkan.compute")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.vulkan.level")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.vulkan.version")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.wifi")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.wifi.aware")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.wifi.direct")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.wifi.passpoint")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.wifi.rtt")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.activities_on_secondary_displays")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.app_widgets")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.autofill")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.backup")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.cant_save_state")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.companion_device_setup")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.connectionservice")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.cts")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.device_admin")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.device_id_attestation")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.file_based_encryption")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.home_screen")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.input_methods")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.live_wallpaper")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.managed_users")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.midi")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.picture_in_picture")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.print")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.securely_removes_users")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.sip")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.sip.voip")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.verified_boot")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.voice_recognizers")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.webview")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.android.apps.dialer.SUPPORTED")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.android.apps.photos.PIXEL_2019_MIDYEAR_PRELOAD")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.android.feature.EXCHANGE_6_2")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.android.feature.GOOGLE_BUILD")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.android.feature.GOOGLE_EXPERIENCE")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.android.feature.PIXEL_2017_EXPERIENCE")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.android.feature.PIXEL_2018_EXPERIENCE")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.android.feature.PIXEL_2019_MIDYEAR_EXPERIENCE")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.android.feature.PIXEL_EXPERIENCE")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.android.feature.TURBO_PRELOAD")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.android.feature.WELLBEING")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.android.feature.ZERO_TOUCH")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.hardware.camera.easel_2018")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.verizon.hardware.telephony.ehrpd")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.verizon.hardware.telephony.lte")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
   }},
}
