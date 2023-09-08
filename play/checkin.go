package play

import (
   "154.pages.dev/encoding/protobuf"
   "bytes"
   "io"
   "net/http"
)

// Checkin$AndroidCheckinResponse
type Device struct {
   m protobuf.Message
}

// androidId
func (d Device) ID() (uint64, error) {
   return d.m.Fixed64(7)
}

// A Sleep is needed after this.
func (c Config) Checkin(platform string) ([]byte, error) {
   var checkin_body = protobuf.Message{
      protobuf.Field{Number: 4, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("PQ3B.190705.003")},
            protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("sargo")},
            protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("google")},
            protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("g670-00011-190411-B-5457439")},
            protobuf.Field{Number: 5, Type: 2, Value: protobuf.Bytes("b4s4-0.1-5613380")},
            protobuf.Field{Number: 6, Type: 2, Value: protobuf.Bytes("android-google")},
            protobuf.Field{Number: 7, Type: 0, Value: protobuf.Varint(1694054582)},
            protobuf.Field{Number: 8, Type: 0, Value: protobuf.Varint(203615028)},
            protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("sargo")},
            protobuf.Field{Number: 10, Type: 0, Value: protobuf.Varint(28)},
            protobuf.Field{Number: 11, Type: 2, Value: protobuf.Bytes("Pixel 3a")},
            protobuf.Field{Number: 12, Type: 2, Value: protobuf.Bytes("google")},
            protobuf.Field{Number: 13, Type: 2, Value: protobuf.Bytes("sargo")},
            protobuf.Field{Number: 14, Type: 0, Value: protobuf.Varint(0)},
         }},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
         protobuf.Field{Number: 6, Type: 2, Value: protobuf.Bytes("334050")},
         protobuf.Field{Number: 7, Type: 2, Value: protobuf.Bytes("20815")},
         protobuf.Field{Number: 8, Type: 2, Value: protobuf.Bytes("mobile-notroaming")},
         protobuf.Field{Number: 9, Type: 0, Value: protobuf.Varint(0)},
         protobuf.Field{Number: 18, Type: 0, Value: protobuf.Varint(1)},
      }},
      protobuf.Field{Number: 14, Type: 0, Value: protobuf.Varint(3)},
   }
   // r.Header.Set("User-Agent", "GoogleAuth/1.4 sargo PQ3B.190705.003")
   res, err := http.Post(
      "https://android.googleapis.com/checkin", "application/x-protobuffer",
      bytes.NewReader(checkin_body.Append(nil)),
   )
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return io.ReadAll(res.Body)
}
