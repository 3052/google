package acquire

import (
   "154.pages.dev/encoding/protobuf"
   "154.pages.dev/google/play"
   "bytes"
   "errors"
   "fmt"
   "io"
   "net/http"
   "net/url"
)

func acquire_2(h *play.Header, doc string, vc uint64) error {
   var acquire_2_body = protobuf.Message{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes(doc)},
            protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1)},
            protobuf.Field{Number: 3, Type: 0, Value: protobuf.Varint(3)},
         }},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1)},
         protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("")},
         protobuf.Field{Number: 7, Type: 0, Value: protobuf.Varint(1)},
      }},
      protobuf.Field{Number: 4, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(10)},
      }},
      protobuf.Field{Number: 8, Type: 2, Value: protobuf.Bytes("")},
      protobuf.Field{Number: 12, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(vc)},
      }},
      protobuf.Field{Number: 13, Type: 0, Value: protobuf.Varint(1)},
      //protobuf.Field{Number: 22, Type: 2, Value: protobuf.Bytes("nonce=zPXXYPxcgwrtB1eIsiqpQrEfdQmj66MyODvudKhhNu_gY3xTVjqdG76ra9Gw3oXjc3xLjg679AwF-wDQQTwK5P71n-0Ibt7EslTYUZNb8MlSMAiPZHy5CUXlP6YsZO65l-2CfgtDINF7ND77Ym6oBL4lpsiKp7KfbGuLXyru0j6roligDJnak6f-zQ7UuNLFa_lQqyezDlLevV2HILD73VUWtDvmNDVR7lqxgsFuZevOMywJ3zVikTXkcYcMeqkthPa5QA9gErMALdpoE1tdnyoHrba-63v2r7lQNY6biCO7TwEsGRR9jZjkfCTEOcqZ7laxXbFa3U1v1w-u03P9vw")},
      protobuf.Field{Number: 25, Type: 0, Value: protobuf.Varint(2)},
      protobuf.Field{Number: 28, Type: 0, Value: protobuf.Varint(1)},
   }
   req := new(http.Request)
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "android.clients.google.com"
   req.URL.Path = "/fdfe/acquire"
   val := make(url.Values)
   val["theme"] = []string{"2"}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(bytes.NewReader(acquire_2_body.Append(nil)))
   req.Header["X-Dfe-Device-Id"] = []string{device_ID}
   req.Header.Set(h.Authorization())
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   {
      b, err := io.ReadAll(res.Body)
      if err != nil {
         return err
      }
      fmt.Printf("%q\n", b)
   }
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   return nil
}
