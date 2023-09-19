package acquire

import (
   "154.pages.dev/google/play"
   "154.pages.dev/protobuf"
   "bytes"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strings"
)

const (
   // android studio
   // device_ID = "344d67278408e17a"
   
   // mine
   device_ID = "30d95bf31df29aac"
   
   // aurora
   //device_ID = "35cd340885ec81f3"
)

func Acquire(h *play.Header, doc string, version uint64) error {
   body := protobuf.Message{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Bytes(doc)},
            protobuf.Field{Number: 2, Type: 0,  Value: protobuf.Varint(1)},
            protobuf.Field{Number: 3, Type: 0,  Value: protobuf.Varint(3)},
         }},
         protobuf.Field{Number: 2, Type: 0,  Value: protobuf.Varint(1)},
         protobuf.Field{Number: 3, Type: 2,  Value: protobuf.Bytes("")},
      }},
      protobuf.Field{Number: 8, Type: 2,  Value: protobuf.Bytes("")},
      protobuf.Field{Number: 12, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 0,  Value: protobuf.Varint(version)},
         protobuf.Field{Number: 3, Type: 0,  Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 13, Type: 0,  Value: protobuf.Varint(1)},
      protobuf.Field{Number: 15, Type: 0,  Value: protobuf.Varint(0)},
      protobuf.Field{Number: 22, Type: 2,  Value: protobuf.Bytes("nonce=z5qVwhnscgbdxNrBx-AsOD1TMGRpF_ohur_MQmR3JCYgaVqP8frLtCofKjPGpBbvbCERVPb7iMFt4entym6eKzsOICEX07ls-7wEITG51SaAa7GCgKXPkLd79hyFhI1FDWVZu5VPpdI9Oxlp_0WvMSVFwI6F5_sHpvI-tQEmOgB9YvTdfRJeyEXHG1da1nzzfAzE7_I38dLEmrHqwhSasb_td4qAceSkpvqvW8dv6NCarUL_Uf_yqJa_UoAjeAgSqKl4-1IgnE0_wQUbJt5CYyeUIrjVx4eSDN23uopFiV0bLR4uALifadia2Df53Fa04SVkwX23TB8T0M-K_L0GyA")},
      protobuf.Field{Number: 25, Type: 0,  Value: protobuf.Varint(2)},
      protobuf.Field{Number: 30, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 0,  Value: protobuf.Varint(2)},
      }},
   }
   req := new(http.Request)
   req.Header = make(http.Header)
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Header["X-Dfe-Device-Id"] = []string{device_ID}
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
   req.Header.Set(h.Authorization())
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   {
      b, err := io.ReadAll(res.Body)
      if err != nil {
         return err
      }
      s := string(b)
      for _, err := range acquire_errors {
         if strings.Contains(s, err) {
            return new_error(err)
         }
      }
   }
   return nil
}

var acquire_errors = []string{
   "Error",
   "Please open my apps to establish a connection with the server.",
   "error",
}

func new_error(s string) error {
   s = strings.TrimSuffix(s, ".")
   return errors.New(strings.ToLower(s))
}
