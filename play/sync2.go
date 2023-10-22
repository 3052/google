package play

import (
   "154.pages.dev/protobuf"
   "bytes"
   "errors"
   "io"
   "net/http"
   "net/url"
)

func (h Header) Sync2(Device) error {
   req := new(http.Request)
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "android.clients.google.com"
   req.URL.Path = "/fdfe/sync"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(bytes.NewReader(sync_body.Append(nil)))
   req.Header.Set(h.Device_ID())
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   return nil
}

var sync_body = protobuf.Message{
   protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Prefix{
         // app.source.getcontact
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Bytes("android.hardware.location.gps")},
         }},
         // br.com.rodrigokolb.realdrum
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Bytes("android.software.midi")},
         }},
         // com.clearchannel.iheartradio.controller
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Bytes("android.hardware.microphone")},
         }},
         // com.google.android.apps.walletnfcrel
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Bytes("android.software.device_admin")},
         }},
         // com.google.android.youtube
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Bytes("android.hardware.touchscreen")},
         }},
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Bytes("android.hardware.wifi")},
         }},
         // com.madhead.tos.zh
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Bytes("android.hardware.sensor.accelerometer")},
         }},
         // com.pinterest
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Bytes("android.hardware.camera")},
         }},
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Bytes("android.hardware.location")},
         }},
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Bytes("android.hardware.screen.portrait")},
         }},
         // com.supercell.brawlstars
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Bytes("android.hardware.touchscreen.multitouch")},
         }},
         // com.sygic.aura
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Bytes("android.hardware.location.network")},
         }},
         // kr.sira.metal
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Bytes("android.hardware.sensor.compass")},
         }},
         // org.thoughtcrime.securesms
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Bytes("android.hardware.telephony")},
         }},
         // org.videolan.vlc
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Bytes("android.hardware.screen.landscape")},
         }},
         // com.amctve.amcfullepisodes
         protobuf.Field{Number: 2, Type: 2,  Value: protobuf.Bytes("org.apache.http.legacy")},
         // com.binance.dev
         protobuf.Field{Number: 2, Type: 2,  Value: protobuf.Bytes("android.test.runner")},
         // com.instagram.android
         protobuf.Field{Number: 4, Type: 2,  Value: protobuf.Bytes("GL_OES_compressed_ETC1_RGB8_texture")},
         // com.kakaogames.twodin
         protobuf.Field{Number: 4, Type: 2,  Value: protobuf.Bytes("GL_KHR_texture_compression_astc_ldr")},
      }},
   }},
   protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 0,  Value: protobuf.Varint(0)},
         protobuf.Field{Number: 2, Type: 0,  Value: protobuf.Varint(1585762304)},
         protobuf.Field{Number: 3, Type: 0,  Value: protobuf.Varint(4)},
         protobuf.Field{Number: 4, Type: 2,  Value: protobuf.Bytes("x86")},
      }},
   }},
   protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 18, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Bytes("am-unknown")},
         protobuf.Field{Number: 2, Type: 2,  Value: protobuf.Bytes("play-ms-android-google")},
         protobuf.Field{Number: 3, Type: 2,  Value: protobuf.Bytes("play-ad-ms-android-google")},
      }},
   }},
   protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 19, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 2, Type: 0,  Value: protobuf.Varint(82941300)},
      }},
   }},
   protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 21, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Bytes("google/sdk_gphone_x86/generic_x86:8.0.0/OSR1.180418.026/6741039:userdebug/dev-keys")},
         protobuf.Field{Number: 2, Type: 0,  Value: protobuf.Varint(26)},
         protobuf.Field{Number: 3, Type: 2,  Value: protobuf.Bytes("")},
         protobuf.Field{Number: 4, Type: 2,  Value: protobuf.Bytes("")},
         protobuf.Field{Number: 5, Type: 2,  Value: protobuf.Bytes("")},
         protobuf.Field{Number: 6, Type: 0,  Value: protobuf.Varint(196609)},
      }},
   }},
}
