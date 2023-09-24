// X-DFE-Cookie
package header

import (
   "154.pages.dev/protobuf"
   "encoding/base64"
)

var cookie_two_out = protobuf.Message{
   protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 3, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Bytes("4530614045536013680")},
         protobuf.Field{Number: 2, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Bytes("1695512393015169")},
            protobuf.Field{Number: 2, Type: 2, Value: protobuf.Prefix{
               protobuf.Field{Number: 1, Type: 0,  Value: protobuf.Varint(1695512393)},
               protobuf.Field{Number: 2, Type: 0,  Value: protobuf.Varint(15169000)},
            }},
         }},
      }},
   }},
}

func cookie_two(s string) (protobuf.Message, error) {
   data, err := base64.StdEncoding.DecodeString(s)
   if err != nil {
      return nil, err
   }
   return protobuf.Consume(data)
}

func cookie_one(s string) (protobuf.Message, error) {
   data, err := base64.RawStdEncoding.DecodeString(s)
   if err != nil {
      return nil, err
   }
   return protobuf.Consume(data)
}

const cookie_two_in = "CjgaNgoTNDUzMDYxNDA0NTUzNjAxMzY4MBIfChAxNjk1NTEyMzkzMDE1MTY5EgsIye69qAYQ6OudBw=="

var cookie_one_out = protobuf.Message{
   protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1)},
   protobuf.Field{Number: 3, Type: 0, Value: protobuf.Varint(0)},
   protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("US")},
   protobuf.Field{Number: 6, Type: 2, Value: protobuf.Bytes("CjgaNgoTNDUzMDYxNDA0NTUzNjAxMzY4MBIfChAxNjk1NTEyMzkzMDE1MTY5EgsIye69qAYQ6OudBw==")},
   protobuf.Field{Number: 8, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("US-TX")},
      protobuf.Field{Number: 2, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(1697931593)},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(191308000)},
      }},
   }},
   protobuf.Field{Number: 9, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("US")},
      protobuf.Field{Number: 2, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(1697931593)},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(191373000)},
      }},
   }},
   protobuf.Field{Number: 11, Type: 0, Value: protobuf.Varint(0)},
}

const cookie_one_in = "EAEYACICVVMyUENqZ2FOZ29UTkRVek1EWXhOREEwTlRVek5qQXhNelk0TUJJZkNoQXhOamsxTlRFeU16a3pNREUxTVRZNUVnc0l5ZTY5cUFZUTZPdWRCdz09QhQKBVVTLVRYEgsIycLRqQYQ4MGcW0oRCgJVUxILCMnC0akGEMi9oFtYAA"
