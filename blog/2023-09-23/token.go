// X-DFE-Device-Config-Token
package header

import (
   "154.pages.dev/protobuf"
   "encoding/base64"
)

var token_one_out = protobuf.Message{
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

func token_one(s string) (protobuf.Message, error) {
   data, err := base64.StdEncoding.DecodeString(s)
   if err != nil {
      return nil, err
   }
   return protobuf.Consume(data)
}

const token_one_in = "CjgaNgoTNDUzMDYxNDA0NTUzNjAxMzY4MBIfChAxNjk1NTEyMzkzMDE1MTY5EgsIye69qAYQ6OudBw=="
