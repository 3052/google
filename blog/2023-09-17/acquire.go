package main

import (
   "154.pages.dev/protobuf"
   "bytes"
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
)

func main() {
   body := protobuf.Message{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Bytes("com.duolingo")},
            protobuf.Field{Number: 2, Type: 0,  Value: protobuf.Varint(1)},
            protobuf.Field{Number: 3, Type: 0,  Value: protobuf.Varint(3)},
         }},
         protobuf.Field{Number: 2, Type: 0,  Value: protobuf.Varint(1)},
         protobuf.Field{Number: 3, Type: 2,  Value: protobuf.Bytes("")},
      }},
      protobuf.Field{Number: 8, Type: 2,  Value: protobuf.Bytes("")},
      /*
      protobuf.Field{Number: 12, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 0,  Value: protobuf.Varint(1700)},
         protobuf.Field{Number: 3, Type: 0,  Value: protobuf.Varint(0)},
      }},
      */
      protobuf.Field{Number: 13, Type: 0,  Value: protobuf.Varint(1)},
      protobuf.Field{Number: 15, Type: 0,  Value: protobuf.Varint(0)},
      //protobuf.Field{Number: 22, Type: 2,  Value: protobuf.Bytes("nonce=z5qVwhnscgbdxNrBx-AsOD1TMGRpF_ohur_MQmR3JCYgaVqP8frLtCofKjPGpBbvbCERVPb7iMFt4entym6eKzsOICEX07ls-7wEITG51SaAa7GCgKXPkLd79hyFhI1FDWVZu5VPpdI9Oxlp_0WvMSVFwI6F5_sHpvI-tQEmOgB9YvTdfRJeyEXHG1da1nzzfAzE7_I38dLEmrHqwhSasb_td4qAceSkpvqvW8dv6NCarUL_Uf_yqJa_UoAjeAgSqKl4-1IgnE0_wQUbJt5CYyeUIrjVx4eSDN23uopFiV0bLR4uALifadia2Df53Fa04SVkwX23TB8T0M-K_L0GyA")},
      protobuf.Field{Number: 25, Type: 0,  Value: protobuf.Varint(2)},
      protobuf.Field{Number: 30, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 0,  Value: protobuf.Varint(2)},
      }},
   }
   var req http.Request
   req.Header = make(http.Header)
   req.Header["Authorization"] = []string{"Bearer ya29.a0AfB_byAXBBcOb0E6BfWrnUzWf1rR7aJGW_jElONLMcws3JhdOrPAX5Zt8wzkD_fcodkBwxeofxtK5htEp4xOX10hf_MQiPN9_AagkzndyraqPp39XY-qDMzqaTi-mR4ZIEGvaYaLRo27cZTKg8MpenWG2TxN12igV3VVOFtOHK-Igf-SX4XHiF2-tyF62Zg6hQoZKBelr_kERd4haLboTznh67syvGB6ElnH8j4QwQZdTyxWNycrC9qaSJz14LPWYW7XwhfFS08JL2i57JaiFdflXd4eIM69UMjABgPWn_DGAnpg-dKVMZmJUAvHBNZ5A8Bc3waCgYKAaISARESFQGOcNnCuty79iAXfjU2kIxMjHf-nw0333"}
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Header["X-Dfe-Device-Id"] = []string{"344d67278408e17a"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "play-fe.googleapis.com"
   req.URL.Path = "/fdfe/acquire"
   val := make(url.Values)
   val["theme"] = []string{"2"}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(bytes.NewReader(body.Append(nil)))
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res_body, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(res_body)
}
