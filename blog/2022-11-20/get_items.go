package main

import (
   "2a.pages.dev/rosso/http"
   "2a.pages.dev/rosso/protobuf"
   "bytes"
   "encoding/base64"
   "fmt"
   "io"
   "net/url"
   "strconv"
   "time"
)

var checkin_body = protobuf.Message{
   4: protobuf.Message{ // checkin
      1: protobuf.Message{ // build
         10: protobuf.Varint(28),
      },
      18: protobuf.Varint(1), // voiceCapable
   },
   14: protobuf.Varint(3),
   18: protobuf.Message{ // deviceConfiguration
      1: protobuf.Varint(0),
      2: protobuf.Varint(0),
      3: protobuf.Varint(0),
      4: protobuf.Varint(0),
      5: protobuf.Varint(0),
      6: protobuf.Varint(0),
      7: protobuf.Varint(0),
      8: protobuf.Varint(0xF_FFFF),
      9: protobuf.Slice[protobuf.String]{
         "android.test.runner",
         "global-miui11-empty.jar",
         "org.apache.http.legacy",
      },
      11: protobuf.String("x86"),
      15: protobuf.Slice[protobuf.String]{
         "GL_OES_compressed_ETC1_RGB8_texture",
         "GL_KHR_texture_compression_astc_ldr",
      },
      26: protobuf.Slice[protobuf.Message]{
         protobuf.Message{
            1: protobuf.String("android.hardware.camera.autofocus"),
         },
         protobuf.Message{
            1: protobuf.String("android.hardware.camera.front"),
         },
         protobuf.Message{
            1: protobuf.String("android.hardware.camera"),
         },
         protobuf.Message{
            1: protobuf.String("android.hardware.location.network"),
         },
         protobuf.Message{
            1: protobuf.String("android.hardware.location.gps"),
         },
         protobuf.Message{
            1: protobuf.String("android.hardware.location"),
         },
         protobuf.Message{
            1: protobuf.String("android.hardware.microphone"),
         },
         protobuf.Message{
            1: protobuf.String("android.hardware.screen.landscape"),
         },
         protobuf.Message{
            1: protobuf.String("android.hardware.screen.portrait"),
         },
         protobuf.Message{
            1: protobuf.String("android.hardware.sensor.accelerometer"),
         },
         protobuf.Message{
            1: protobuf.String("android.hardware.telephony"),
         },
         protobuf.Message{
            1: protobuf.String("android.hardware.touchscreen"),
         },
         protobuf.Message{
            1: protobuf.String("android.hardware.touchscreen.multitouch"),
         },
         protobuf.Message{
            1: protobuf.String("android.hardware.usb.host"),
         },
         protobuf.Message{
            1: protobuf.String("android.hardware.wifi"),
         },
         protobuf.Message{
            1: protobuf.String("android.software.device_admin"),
         },
      },
   },
}

func checkin() (string, error) {
   req, err := http.NewRequest(
      "POST", "https://android.googleapis.com/checkin",
      bytes.NewReader(checkin_body.Marshal()),
   )
   if err != nil {
      return "", err
   }
   req.Header.Set("Content-Type", "application/x-protobuffer")
   res, err := client.Do(req)
   if err != nil {
      return "", err
   }
   defer res.Body.Close()
   body, err := io.ReadAll(res.Body)
   if err != nil {
      return "", err
   }
   message, err := protobuf.Unmarshal(body)
   if err != nil {
      return "", err
   }
   id, err := message.Get_Fixed64(7)
   if err != nil {
      return "", err
   }
   return strconv.FormatUint(id, 16), nil
}

func sync(device string) error {
   req := new(http.Request)
   req.Body = io.NopCloser(bytes.NewReader(sync_body.Marshal()))
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "play-fe.googleapis.com"
   req.URL.Path = "/fdfe/sync"
   req.URL.Scheme = "https"
   req.Header["X-Dfe-Device-Id"] = []string{device}
   if _, err := client.Do(req); err != nil {
      return err
   }
   return nil
}

func main() {
   device, err := checkin()
   if err != nil {
      panic(err)
   }
   if err := sync(device); err != nil {
      panic(err)
   }
   fmt.Println("Sleep", sleep)
   time.Sleep(sleep)
   req := new(http.Request)
   req_body := protobuf.Message{
      2: protobuf.Message{
         1: protobuf.Message{
            1: protobuf.Message{
               1: protobuf.String("com.watchfacestudio.md307digital"),
            },
         },
      },
   }.Marshal()
   req.Body = io.NopCloser(bytes.NewReader(req_body))
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "play-fe.googleapis.com"
   req.URL.Path = "/fdfe/getItems"
   req.URL.Scheme = "https"
   req.Header["X-Dfe-Device-Id"] = []string{device}
   field := protobuf.Message{
      4: protobuf.Bytes{0xFF, 0xFF, 0xFF, 0xFF},
   }.Marshal()
   mask := base64.StdEncoding.EncodeToString(field)
   req.Header.Set("X-Dfe-Item-Field-Mask", mask)
   fmt.Println(device)
   res, err := client.Do(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res_body, err := io.ReadAll(res.Body)
   if err != nil {
      panic(err)
   }
   if bytes.Contains(res_body, []byte("config.xhdpi")) {
      fmt.Println("pass")
   } else {
      fmt.Println("fail")
   }
}

const sleep = 16 * time.Second

var client = http.Default_Client

var sync_body = protobuf.Message{
   1: protobuf.Slice[protobuf.Message]{
      protobuf.Message{
         10: protobuf.Message{
            1: protobuf.Slice[protobuf.Message]{
               protobuf.Message{
                  1: protobuf.String("android.hardware.camera.autofocus"),
               },
               protobuf.Message{
                  1: protobuf.String("android.hardware.camera.front"),
               },
               protobuf.Message{
                  1: protobuf.String("android.hardware.camera"),
               },
               protobuf.Message{
                  1: protobuf.String("android.hardware.location.network"),
               },
               protobuf.Message{
                  1: protobuf.String("android.hardware.location.gps"),
               },
               protobuf.Message{
                  1: protobuf.String("android.hardware.location"),
               },
               protobuf.Message{
                  1: protobuf.String("android.hardware.microphone"),
               },
               protobuf.Message{
                  1: protobuf.String("android.hardware.screen.landscape"),
               },
               protobuf.Message{
                  1: protobuf.String("android.hardware.screen.portrait"),
               },
               protobuf.Message{
                  1: protobuf.String("android.hardware.sensor.accelerometer"),
               },
               protobuf.Message{
                  1: protobuf.String("android.hardware.telephony"),
               },
               protobuf.Message{
                  1: protobuf.String("android.hardware.touchscreen"),
               },
               protobuf.Message{
                  1: protobuf.String("android.hardware.touchscreen.multitouch"),
               },
               protobuf.Message{
                  1: protobuf.String("android.hardware.usb.host"),
               },
               protobuf.Message{
                  1: protobuf.String("android.hardware.wifi"),
               },
               protobuf.Message{
                  1: protobuf.String("android.software.device_admin"),
               },
            },
            2: protobuf.Slice[protobuf.String]{
               "org.apache.http.legacy",
               "android.test.runner",
            },
            4: protobuf.Slice[protobuf.String]{
               "GL_OES_compressed_ETC1_RGB8_texture",
               "GL_KHR_texture_compression_astc_ldr",
            },
         },
      },
      protobuf.Message{
         19: protobuf.Message{
            2: protobuf.Varint(83032110),
         },
      },
      protobuf.Message{
         20: protobuf.Message{
            2: protobuf.Varint(768),
            3: protobuf.Varint(1280),
            4: protobuf.Varint(2),
            5: protobuf.Varint(320),
            1: protobuf.Varint(3),
         },
      },
   },
}
