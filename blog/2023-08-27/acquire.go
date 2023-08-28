/*
one or more of the values here, only needs to be provided once
*/
package main

import (
   "154.pages.dev/encoding/protobuf"
   "bytes"
   "fmt"
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
)

var message = protobuf.Message{
   protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("sg.bigo.live")},
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
      protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(651173)},
      protobuf.Field{Number: 3, Type: 0, Value: protobuf.Varint(0)},
   }},
   protobuf.Field{Number: 13, Type: 0, Value: protobuf.Varint(1)},
   protobuf.Field{Number: 25, Type: 0, Value: protobuf.Varint(2)},
}

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Header["Authorization"] = []string{"Bearer ya29.a0AfB_byAl_x_9sQ7Mn0NYODLee5aBJzA-jJOnK92Nsomgl3lrAY-ABzY5Z5qtOXql3S6ULw-cpuP6-_wqomCuTExhp4nAzsxMZTzx4L0JyTeUych8txUMm87NdIBvVsjMwo6rsoFp7SJz-ltGQuq0-rM4XLGGFl8PDTOQn42-zNhe6n__xnqJ50O6in-vmEBwsGOe8OZ9qC9_NYfOx_s19KyhZ4XBKfP_OVprU5hBPvSmKXqxdtu0ViKpvgNT4ZP6lEjhKthHcz8XjFJFkJv89D8vaePl3J9PUwfPtPb6_d_Fr9FijkEk9lThBx2O_yfnxJdSDEBtZQaCgYKARcSARESFQHsvYls1RWcFpHM5RF-XWlQFCrIFA0337"}
   req.Header["X-Dfe-Device-Id"] = []string{"32cf69cd18ab8bcc"}
   req.Header["Accept-Encoding"] = []string{"identity"}
   req.Header["Accept-Language"] = []string{"en-US"}
   req.Header["Connection"] = []string{"Keep-Alive"}
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Header["Host"] = []string{"android.clients.google.com"}
   req.Header["User-Agent"] = []string{"Android-Finsky/22.0.18-21%20%5B0%5D%20%5BPR%5D%20333153705 (api=3,versionCode=82201810,sdk=21,device=generic_x86,hardware=ranchu,product=sdk_google_phone_x86,platformVersionRelease=5.0.2,model=Android%20SDK%20built%20for%20x86,buildId=LSY66K,isWideScreen=0,supportedAbis=x86)"}
   req.Header["X-Dfe-Client-Id"] = []string{"am-google"}
   req.Header["X-Dfe-Content-Filters"] = []string{""}
   req.Header["X-Dfe-Mccmnc"] = []string{"310260"}
   req.Header["X-Dfe-Network-Type"] = []string{"3"}
   req.Header["X-Dfe-Request-Params"] = []string{"timeoutMs=35000"}
   req.Header["X-Public-Android-Id"] = []string{"99512e095ec4e0b8"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "android.clients.google.com"
   req.URL.Path = "/fdfe/acquire"
   req.URL.RawPath = ""
   val := make(url.Values)
   val["theme"] = []string{"2"}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(bytes.NewReader(message.Append(nil)))
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res_body, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%q\n", res_body)
}
