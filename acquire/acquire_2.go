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

func Acquire_2(h *play.Header, doc string, vc uint64) error {
   var acquire_2_body = protobuf.Message{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes(doc)},
            protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1)},
            protobuf.Field{Number: 3, Type: 0, Value: protobuf.Varint(3)},
         }},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1)},
      }},
      protobuf.Field{Number: 4, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(10)},
      }},
      protobuf.Field{Number: 13, Type: 0, Value: protobuf.Varint(1)},
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
