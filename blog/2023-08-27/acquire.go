package main

import (
   "154.pages.dev/encoding/protobuf"
   "bytes"
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
)

var req_body = protobuf.Message{
   protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.gamma.scan")},
         //protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1)},
         //protobuf.Field{Number: 3, Type: 0, Value: protobuf.Varint(3)},
      }},
      //protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1)},
      //protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("")},
   }},
   /*
   protobuf.Field{Number: 8, Type: 2, Value: protobuf.Bytes("\xa0\x01\x00")},
   protobuf.Field{Number: 8, Type: 2, Value: protobuf.Maybe{
      protobuf.Field{Number: 20, Type: 0, Value: protobuf.Varint(0)},
   }},
   protobuf.Field{Number: 12, Type: 2, Value: protobuf.Bytes("\b\xa2\x01\x18\x00")},
   protobuf.Field{Number: 12, Type: 2, Value: protobuf.Maybe{
      protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(162)},
      protobuf.Field{Number: 3, Type: 0, Value: protobuf.Varint(0)},
   }},
   protobuf.Field{Number: 13, Type: 0, Value: protobuf.Varint(1)},
   protobuf.Field{Number: 22, Type: 2, Value: protobuf.Bytes("nonce=vTuw21ROE5t4cunOg3LNhJOSOKo3e4LVVQ7pvzKborEo7IwOyGuCzGEHyFUvNq2GyAY-e7DN3uYOYDHocNlG4do2AkVpaaZuGlQKPgIMDdSZy_XCHwPF1kCLXGeK9eh0yr4YlvL_dY570AEoKHr6eGKEVVac7wYBUNyE2AflIk6wVxxWv0KGJuY1rP4FCKVvxBTL5K1mpxkidwAIzQh8tUbDUTjINK_lH_bkIDJ5nJiA8x5LTl9mgYbnaEJvg8DhXziroPQ6VKWIIlBqpoN5socLWGW_pigCoZoTBurAI6TAR3ZEK0GCBD_4nNNQ1dlPFfNHxNRGyKXnSi2WUIaW2A")},
   protobuf.Field{Number: 25, Type: 0, Value: protobuf.Varint(2)},
   protobuf.Field{Number: 30, Type: 2, Value: protobuf.Bytes("\b\x02\x10\x00")},
   protobuf.Field{Number: 30, Type: 2, Value: protobuf.Maybe{
      protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(2)},
      protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
   }},
   protobuf.Field{Number: 31, Type: 2, Value: protobuf.Bytes("\b\xa9թ\xa7\x06\x10\x80\xb7\x8a<")},
   protobuf.Field{Number: 31, Type: 2, Value: protobuf.Maybe{
      protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(1693084329)},
      protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(126000000)},
   }},
   */
}

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "play-fe.googleapis.com"
   req.URL.Path = "/fdfe/acquire"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(bytes.NewReader(req_body.Append(nil)))
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Header["X-Dfe-Device-Id"] = []string{"37e9443505dd7eb0"}
   req.Header["Authorization"] = []string{"Bearer ya29.a0AfB_byCaSkA9DZP1Iagl6sJ5edyQT2Tsf0LJHlTlakY6WPXeMMySmLh6pvYJ8HCLtOwG1D2RZdaCbmTgaCgHtb-EW9WBr0jSyB6TOpZS2MLQ2h4o-_HF1cqsq-yECT7y154kwAApvVeMBe9l8pAp-FV2hqC2mUuSPHe6Ml65qFqDIuNEYd4YDuFoW6mHlyaqWcWQmVAQHvaaoMuxKaoJ8t8PMGn4cGaoz7yUfWbF9CMdeeblJHnlpO1OJaBVox8FCjnZRhJEdl87QcAxwlNBjJ2xD4-VVxpima1ONw-BXCI6G2nAz3ehQU5pmMYFQgm2vxU1cnHU9waCgYKAdQSARESFQHsvYlslYCNuThThWmcCGx-zV9xPA0337"}
   //val := make(url.Values)
   //val["theme"] = []string{"2"}
   //req.URL.RawQuery = val.Encode()
   //req.Header["Accept-Encoding"] = []string{"identity"}
   //req.Header["Accept-Language"] = []string{"en-US"}
   //req.Header["Connection"] = []string{"Keep-Alive"}
   //req.Header["Host"] = []string{"play-fe.googleapis.com"}
   //req.Header["User-Agent"] = []string{"Android-Finsky/30.3.21-21%20%5B0%5D%20%5BPR%5D%20445437866 (api=3,versionCode=83032110,sdk=21,device=generic_x86,hardware=ranchu,product=sdk_google_phone_x86,platformVersionRelease=5.0.2,model=Android%20SDK%20built%20for%20x86,buildId=LSY66K,isWideScreen=0,supportedAbis=x86)"}
   //req.Header["X-Dfe-Client-Id"] = []string{"am-google"}
   //req.Header["X-Dfe-Cookie"] = []string{"EAEYACICVVNKEgoCVVMSDAjVpr2oBhDA1MrRAVgA"}
   //req.Header["X-Dfe-Encoded-Targets"] = []string{"CAEaggFXjAUF0I+BBgnEAgQBAQ3nAX0vugFWASEKIw0iN4QBDgxjCViqAROMAhRMWQEBFFsZ4gUEkgI8nAKnB6QEWb0GkgrgBeAfrhzuBuQHmxoCAZAFpQelAu0K9werAuMJgwO8AssBFuoYuiafP4MGRVYBjgXUAq4R6RqIDvYL9wKbA8c7"}
   //req.Header["X-Dfe-Mccmnc"] = []string{"310260"}
   //req.Header["X-Dfe-Network-Type"] = []string{"3"}
   //req.Header["X-Dfe-Request-Params"] = []string{"timeoutMs=35000"}
   //req.Header["X-Ps-Rh"] = []string{"H4sIAAAAAAAAAONqZuRqYPQwKfZ0hIJ4U99K0wpPI98Q58piF8dK36pII1_zwPISl0Bj36pkQ9ccn3TXLEeD4qpAo2IXZ6MSH2egnGOVX5VnmEmpu0lybl5OclaiU1ahhUdRlmtSZYRHRXKaW3mVY6JPvFuGZXxlYpWrY36Ao1OOboW_UxTITgAYYVMrhgAAAA"}
   //req.Header["X-Public-Android-Id"] = []string{"3ac4cbd266c19f06"}
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

