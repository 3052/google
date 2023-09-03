package main

import (
   "154.pages.dev/encoding/protobuf"
   "encoding/json"
   "os"
)

func main() {
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   enc.Encode(pass)
}

var pass = protobuf.Message{
   protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 94, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 3, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 3, Type: 2, Value: protobuf.Prefix{
               protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(2)},
               protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(3)},
               protobuf.Field{Number: 3, Type: 2, Value: protobuf.Prefix{
                  protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1693695217)},
                  protobuf.Field{Number: 3, Type: 1, Value: protobuf.Fixed64(4744890680618272632)},
               }},
               protobuf.Field{Number: 4, Type: 2, Value: protobuf.Prefix{
                  protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
                     protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.zutgames.ilovehue")},
                     protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1)},
                     protobuf.Field{Number: 3, Type: 0, Value: protobuf.Varint(3)},
                  }},
                  protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1)},
                  protobuf.Field{Number: 3, Type: 0, Value: protobuf.Varint(244010920912327096)},
                  protobuf.Field{Number: 5, Type: 2, Value: protobuf.Prefix{
                     protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("HZJ20VakK5h5OKjZ0JoMoYfg7Qc")},
                  }},
                  protobuf.Field{Number: 12, Type: 0, Value: protobuf.Varint(1693695217052)},
               }},
               protobuf.Field{Number: 6, Type: 2, Value: protobuf.Bytes("3")},
            }},
            protobuf.Field{Number: 7, Type: 0, Value: protobuf.Varint(0)},
            protobuf.Field{Number: 10, Type: 2, Value: protobuf.Prefix{
               protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(1)},
               protobuf.Field{Number: 2, Type: 2, Value: protobuf.Prefix{
                  protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
                     protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("RESPONSE_CODE")},
                     protobuf.Field{Number: 5, Type: 0, Value: protobuf.Varint(0)},
                  }},
               }},
            }},
         }},
         protobuf.Field{Number: 4, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
               protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
                  protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.zutgames.ilovehue")},
                  protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(7)},
               }},
               protobuf.Field{Number: 2, Type: 2, Value: protobuf.Prefix{
                  protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("CAE=")},
               }},
               protobuf.Field{Number: 5, Type: 2, Value: protobuf.Prefix{
                  protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
                     protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.zutgames.ilovehue")},
                     protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(7)},
                  }},
                  protobuf.Field{Number: 2, Type: 2, Value: protobuf.Prefix{
                     protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("CAE=")},
                  }},
                  protobuf.Field{Number: 5, Type: 0, Value: protobuf.Varint(1)},
               }},
            }},
            protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("")},
            protobuf.Field{Number: 5, Type: 2, Value: protobuf.Prefix{
               protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(1)},
            }},
            protobuf.Field{Number: 7, Type: 2, Value: protobuf.Bytes("")},
            protobuf.Field{Number: 11, Type: 2, Value: protobuf.Prefix{
               protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1)},
            }},
            protobuf.Field{Number: 22, Type: 2, Value: protobuf.Prefix{
               protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("-7492140743066199944")},
            }},
            protobuf.Field{Number: 35, Type: 2, Value: protobuf.Prefix{
               protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("")},
            }},
         }},
         protobuf.Field{Number: 6, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(0)},
            protobuf.Field{Number: 2, Type: 2, Value: protobuf.Prefix{
               protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(5)},
               protobuf.Field{Number: 6, Type: 2, Value: protobuf.Prefix{
                  protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(2)},
                  protobuf.Field{Number: 2, Type: 2, Value: protobuf.Prefix{
                     protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(0)},
                  }},
                  protobuf.Field{Number: 3, Type: 2, Value: protobuf.Prefix{
                     protobuf.Field{Number: 2, Type: 2, Value: protobuf.Prefix{
                        protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.zutgames.ilovehue")},
                        protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1)},
                        protobuf.Field{Number: 3, Type: 0, Value: protobuf.Varint(3)},
                     }},
                  }},
                  protobuf.Field{Number: 4, Type: 0, Value: protobuf.Varint(1)},
                  protobuf.Field{Number: 5, Type: 2, Value: protobuf.Prefix{
                     protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(0)},
                     protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("USD")},
                  }},
                  protobuf.Field{Number: 11, Type: 2, Value: protobuf.Bytes("")},
                  protobuf.Field{Number: 18, Type: 2, Value: protobuf.Bytes("")},
                  protobuf.Field{Number: 19, Type: 0, Value: protobuf.Varint(0)},
                  protobuf.Field{Number: 22, Type: 0, Value: protobuf.Varint(0)},
                  protobuf.Field{Number: 23, Type: 2, Value: protobuf.Bytes("")},
                  protobuf.Field{Number: 33, Type: 2, Value: protobuf.Bytes("-7492140743066199944")},
                  protobuf.Field{Number: 42, Type: 2, Value: protobuf.Prefix{
                     protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(1)},
                     protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
                  }},
               }},
            }},
            protobuf.Field{Number: 3, Type: 0, Value: protobuf.Varint(1)},
         }},
         protobuf.Field{Number: 8, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 3, Type: 2, Value: protobuf.Prefix{
               protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
                  protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
                     protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("RESPONSE_CODE")},
                     protobuf.Field{Number: 5, Type: 0, Value: protobuf.Varint(0)},
                  }},
               }},
               protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1)},
            }},
         }},
         protobuf.Field{Number: 10, Type: 0, Value: protobuf.Varint(1)},
      }},
   }},
   protobuf.Field{Number: 5, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(822)},
   }},
   protobuf.Field{Number: 9, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(18)},
      protobuf.Field{Number: 19, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(1693695216432165)},
            protobuf.Field{Number: 2, Type: 5, Value: protobuf.Fixed32(177376833)},
            protobuf.Field{Number: 3, Type: 5, Value: protobuf.Fixed32(2818903093)},
         }},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1)},
      }},
   }},
}
