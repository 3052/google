package play

import (
   "154.pages.dev/protobuf"
   "bytes"
   "fmt"
   "io"
   "net/http"
   "net/url"
)

func (h Header) Acquire(doc string) error {
   var req http.Request
   req.Header = make(http.Header)
   req.Header.Set(h.Authorization())
   req.Header.Set(h.Device_Config())
   req.Header.Set(h.Device_ID())
   // org.videolan.vlc
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
      /*
      m.Add(4, func(m *protobuf.Message) {
         m.Add_Varint(1, 10)
      })
      */
      protobuf.Field{Number: 8, Type: 2,  Value: protobuf.Bytes("")},
      /*
      protobuf.Field{Number: 12, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 0,  Value: protobuf.Varint(version)},
         protobuf.Field{Number: 3, Type: 0,  Value: protobuf.Varint(0)},
      }},
      */
      protobuf.Field{Number: 13, Type: 0,  Value: protobuf.Varint(1)},
      protobuf.Field{Number: 15, Type: 0,  Value: protobuf.Varint(0)},
      protobuf.Field{Number: 22, Type: 2,  Value: protobuf.Bytes("nonce=qlIhuESfWlLW7GI7k6fWej7z403Mynf3o0dl5B9RYfWxmHxGGSGqBARF_TxpzL5RfVPW3oFX0zAHhSETtuUBv7TvrzWOx5hEgolPjFDs1lr_Po9lyH1HJ8UskVSkyMe_gImmY0-hA-I0SSaBfUXyInciRuuMtSNiXsMclJwWoW1bPgjYsQKCn5szQnDPlMvSqz4hBCbIGxGKiWe6L9f6IcmfIwlz8eSRQUA02YN633zvXDbptBIKfrpwE9_P_N0sWrOhc3k9LAQlrn4f4kXor4g98ZQ6BN6U3us7US-2ES-xiCaFvdlbIMiMvpp7_AsnLon1KyvxS_rujvoDaUyOOQ")},
      protobuf.Field{Number: 25, Type: 0,  Value: protobuf.Varint(2)},
      protobuf.Field{Number: 30, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 0,  Value: protobuf.Varint(2)},
      }},
   }
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
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Header["Accept-Language"] = []string{"en-US"}
   req.Header["Connection"] = []string{"Keep-Alive"}
   req.Header["X-Dfe-Client-Id"] = []string{"am-unknown"}
   req.Header["X-Dfe-Mccmnc"] = []string{"310260"}
   req.Header["X-Dfe-Network-Type"] = []string{"4"}
   req.Header["X-Dfe-Request-Params"] = []string{"timeoutMs=35000"}
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return fmt.Errorf(res.Status)
   }
   {
      b, err := io.ReadAll(res.Body)
      if err != nil {
         return err
      }
      if bytes.Contains(b, []byte("Error")) {
         return fmt.Errorf("%q", b)
      }
   }
   return nil
}
