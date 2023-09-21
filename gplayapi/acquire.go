package gplayapi

import (
   "154.pages.dev/protobuf"
   "bytes"
   "fmt"
   "io"
   "net/http"
   "net/url"
)

func (g GooglePlayClient) Acquire(doc string, version uint64) error {
   var req http.Request
   req.Header = make(http.Header)
   req.Header["X-Dfe-Device-Id"] = []string{g.AuthData.GsfID}
   req.Header["Authorization"] = []string{"Bearer " + g.AuthData.AuthToken}
   acquire_body := protobuf.Message{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Bytes(doc)},
            protobuf.Field{Number: 2, Type: 0,  Value: protobuf.Varint(1)},
            protobuf.Field{Number: 3, Type: 0,  Value: protobuf.Varint(3)},
         }},
         protobuf.Field{Number: 2, Type: 0,  Value: protobuf.Varint(1)},
         protobuf.Field{Number: 3, Type: 2,  Value: protobuf.Bytes("")},
         protobuf.Field{Number: 7, Type: 0,  Value: protobuf.Varint(1)},
      }},
      protobuf.Field{Number: 8, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 20, Type: 0,  Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 12, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 0,  Value: protobuf.Varint(version)},
         protobuf.Field{Number: 3, Type: 0,  Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 13, Type: 0,  Value: protobuf.Varint(1)},
      protobuf.Field{Number: 22, Type: 2,  Value: protobuf.Bytes("nonce=EBWK4qj_fNtEs0tox4hbh-G1u4JopnoReuV2oKghIivOHFEeiTi6Sp5RYynfywoaku9lU9HemuJ8qRVxKCCF6jPL1lrWj6i2OGFqYowiAgKzjqPjAgQMFGKYCRWvnxZeQqWjhzLE1yulSwmeFuZ9V380vfBvevWkGK82JemK8cOwWOiYUyYWnKO05ODUrpowvHs8hqFe8HaRM_D3_c9VZYgkMkL-RKsBQ3nn5jvkMDcbNeOt71LZ0INcu28k8lLOaDDJSNb7Ip4aSBLN427tDCnmNFhfKvOJJHwvrSiJCrHTh4GJFOYkfrUI3b1EhcEvA6KVGliLsZMJXJXm8g8mug")},
      protobuf.Field{Number: 25, Type: 0,  Value: protobuf.Varint(2)},
      protobuf.Field{Number: 30, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 0,  Value: protobuf.Varint(2)},
      }},
      protobuf.Field{Number: 31, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 0,  Value: protobuf.Varint(1695255934)},
         protobuf.Field{Number: 2, Type: 0,  Value: protobuf.Varint(361000000)},
      }},
   }
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
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
   req.Body = io.NopCloser(bytes.NewReader(acquire_body.Append(nil)))
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
