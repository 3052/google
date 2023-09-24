package gplayapi

import (
   "154.pages.dev/http/option"
   "154.pages.dev/protobuf"
   "bytes"
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "os"
   "testing"
)

// 10 minute fail
func Test_Acquire(t *testing.T) {
   text, err := os.ReadFile("client.json")
   if err != nil {
      t.Fatal(err)
   }
   var client GooglePlayClient
   if err := json.Unmarshal(text, &client); err != nil {
      t.Fatal(err)
   }
   option.No_Location()
   option.Trace()
   if err := client.Acquire("kr.sira.metal", 74); err != nil {
      t.Fatal(err)
   }
}

func Test_Client(t *testing.T) {
   token, err := func() (string, error) {
      b, err := os.ReadFile(`C:\Users\Steven\google\play\token.txt`)
      if err != nil {
         return "", err
      }
      return parseResponse(string(b))["Token"], nil
   }()
   if err != nil {
      t.Fatal(err)
   }
   client, err := NewClientWithDeviceInfo(
      "srpen6@gmail.com", token, Emulator_x86,
   )
   if err != nil {
      t.Fatal(err)
   }
   file, err := os.Create("client.json")
   if err != nil {
      t.Fatal(err)
   }
   defer file.Close()
   enc := json.NewEncoder(file)
   enc.SetIndent("", " ")
   enc.Encode(client)
}

func (g GooglePlayClient) Acquire(doc string, version uint64) error {
   var req http.Request
   req.Header = make(http.Header)
   
   req.Header.Set("X-DFE-Cookie", "")
   req.Header.Set("X-DFE-Device-Config-Token", "")
   
   req.Header["X-Dfe-Device-Id"] = []string{g.AuthData.GsfID}
   req.Header["Authorization"] = []string{"Bearer " + g.AuthData.AuthToken}
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Header["Accept-Language"] = []string{"en-US"}
   req.Header["Connection"] = []string{"Keep-Alive"}
   req.Header["X-Dfe-Client-Id"] = []string{"am-unknown"}
   req.Header["X-Dfe-Mccmnc"] = []string{"310260"}
   req.Header["X-Dfe-Network-Type"] = []string{"4"}
   req.Header["X-Dfe-Request-Params"] = []string{"timeoutMs=35000"}
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
      protobuf.Field{Number: 8, Type: 2,  Value: protobuf.Bytes("")},
      protobuf.Field{Number: 12, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 0,  Value: protobuf.Varint(version)},
         protobuf.Field{Number: 3, Type: 0,  Value: protobuf.Varint(0)},
      }},
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
