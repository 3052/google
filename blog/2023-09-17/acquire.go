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

func main() {
   body := protobuf.Message{
      protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2, Omit: false, Value: protobuf.Bytes("com.duolingo")},
            protobuf.Field{Number: 2, Type: 0, Omit: false, Value: protobuf.Varint(1)},
            protobuf.Field{Number: 3, Type: 0, Omit: false, Value: protobuf.Varint(3)},
         }},
         protobuf.Field{Number: 2, Type: 0, Omit: false, Value: protobuf.Varint(1)},
         protobuf.Field{Number: 3, Type: 2, Omit: false, Value: protobuf.Bytes("")},
         protobuf.Field{Number: 7, Type: 0, Omit: false, Value: protobuf.Varint(1)},
      }},
      protobuf.Field{Number: 12, Type: 2,  Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 0, Omit: false, Value: protobuf.Varint(1700)},
      }},
      protobuf.Field{Number: 13, Type: 0, Omit: false, Value: protobuf.Varint(1)},
      protobuf.Field{Number: 22, Type: 2, Omit: false, Value: protobuf.Bytes("nonce=Op6ue4gbbxuuPG5SQV1pAgVP4MG4mpK8Y9jySzRSHFGmdDuF8fYQV2eitbvCmS8hJIvRpCUl7SMtloFIPLFPRmPXnBU2bBJIc5INXjyakfpYoHl1oyQgQtUZZ4QQ-GkS23VyyEUsSwVzqTZlTLpI3x3oE09rCaBsPVNxAcFuPHuR8FQusQAKvyeLfLP8yz2onmOWfjTmyj1pydd2TY629TSXRhaWrFK3poAAXJHE9kCK13ugR8cbOMETcc-22_4sQLdsftFn6wrSCfsBf8r9W72g4pTlmy2zxHxhN6BJv5v7oMFZ1E4aaiH5IQRGR_ckM7et9qXMNsSXs0lwU52xVQ")},
      protobuf.Field{Number: 25, Type: 0, Omit: false, Value: protobuf.Varint(1)},
   }
   var req http.Request
   req.Header = make(http.Header)
   req.Header["Authorization"] = []string{"Bearer ya29.a0AfB_byAoIBf1zus62HhiNeCeeYzfjMoCQrePnesbbnHIE41u1IRfZji2NpjwiUKKa7mUvH473sS5ykvBSAGItOK2lqPTenL5CCimJrE7Se-ZublD7J0ij_O7je5ENbAMj6RWonqKeKv9gchIMCKOsDoRYd-htrKEjJdwadqEKzbhWaZuD3eoehbts6ltNxs4pbI7YT4KQDIRMa-2teUL2jEd8CxNkXi3jMd4YUdUsnCreuwg2nHyuig6Rdgi6eKpDHpvSK9N3HpUP8yEDNFZeiRC58raFlLwvj8rM0CmyUltRHxdKQkyZ39kBjvajOFH70J7-gaCgYKARISARESFQGOcNnC9DFuvWMX4eHk3QpVJ7AQfQ0333"}
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Header["X-Dfe-Device-Id"] = []string{"39f20681ab4baadb"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "play-fe.googleapis.com"
   req.URL.Path = "/fdfe/acquire"
   val := make(url.Values)
   val["theme"] = []string{"1"}
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
