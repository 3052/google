package gplayapi

import (
   "154.pages.dev/encoding/protobuf"
   "bytes"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strings"
)

func (a AuthData) Acquire(doc string, version uint64) error {
   req := new(http.Request)
   req.Header = make(http.Header)
   req.Header["X-Dfe-Device-Id"] = []string{a.GsfID}
   req.Header["Authorization"] = []string{"Bearer " + a.AuthToken}
   body := protobuf.Message{
      protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2, Omit: false, Value: protobuf.Bytes(doc)},
            protobuf.Field{Number: 2, Type: 0, Omit: false, Value: protobuf.Varint(1)},
            protobuf.Field{Number: 3, Type: 0, Omit: false, Value: protobuf.Varint(3)},
         }},
         protobuf.Field{Number: 2, Type: 0, Omit: false, Value: protobuf.Varint(1)},
         protobuf.Field{Number: 3, Type: 2, Omit: false, Value: protobuf.Bytes("")},
         protobuf.Field{Number: 7, Type: 0, Omit: false, Value: protobuf.Varint(1)},
      }},
      protobuf.Field{Number: 12, Type: 2,  Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 0, Omit: false, Value: protobuf.Varint(version)},
      }},
      protobuf.Field{Number: 13, Type: 0, Omit: false, Value: protobuf.Varint(1)},
      protobuf.Field{Number: 22, Type: 2, Omit: false, Value: protobuf.Bytes("nonce=_c4trPkYCClWis3v9vu1CoFlTEiNrbWoY79uFHBwsw_k7BG2peExk2V5eW_D4tTH_XV7aG9UFt3lrsFdc3F8KGXbK6sPaSIyiOhiXk93kqGarXYEBYRF-tOo9JiuBKVeA-VGCw56QY4lbr1vUgNh-De097T1Ba-1z-o62f2H3qVzaVzkKsvdcs6d7H8nHF8G2tU0Od8T3NGayZowiWGOfH5I-DNC9JztUz7owabG6l2R1WKoigyyvLUJ0dX_BcrSkjhs9XUujHuE9_2RX0WTO46rXyNJ3vijKnqcyswPilhQVsL71UK2g4uIR_VbxbXTe_NAT2KdV56G8p8cnyY8VQ")},
      protobuf.Field{Number: 25, Type: 0, Omit: false, Value: protobuf.Varint(1)},
   }
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "android.clients.google.com"
   req.URL.Path = "/fdfe/acquire"
   val := make(url.Values)
   val["theme"] = []string{"1"}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(bytes.NewReader(body.Append(nil)))
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   data, err := io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   if strings.Contains(string(data), connection) {
      return new_error(connection)
   }
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   return nil
}

const connection = "Please open my apps to establish a connection with the server."

func new_error(s string) error {
   s = strings.TrimSuffix(s, ".")
   return errors.New(strings.ToLower(s))
}
