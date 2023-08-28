# august 28 2023

https://github.com/1268/google/commit/95b10e8e8a05c81704ef70601e997e06c6ebd9d5

in that commit, user posted this value:

~~~
H4sIAAAAAAAAADWRTavjdBTGp3cY79CNcldyURBxUMRck3_emsUIJ--vTdMmbZpNyFuT9C1N0zZpVzJ-AT-DS8cPoN9iwJ24Ety4dydi9Y4HHnjgwHN-PKf_Ta__dU-lGg3eTkhbZ7rTkOV6eLMUiKEIbePC2RK9y_ASIVk2c2kJeHNxkLV32uHyKvaqi-dQR7nK2GFnByco6ogl6NKizSV1VHMqqGTZ3Z7D01weC5qngp1aRTEM_r05_vmm_-am_2NPAGtSONIwSuJT4Si7MMZFvk1NNdYU1the-HwmAivIvGnNxGNTUZ6RWEnZorZVmba0hH2-6wo_45MKUH7OIK4KYB0ZbDBBbKBr6QGYrVi1GqhmpYDWGBxYcAbxCOlABMUBxlEkYDSFyEcrMRcy3hdSMfB4eTyht7FmXnZsJNqo4KvE5jYmNmkS5e7b_9CDnHCLjJEhZsPaUVJI1CVQK-WYo245bK0t0yg826xg5AigawSYjgWjRAI7cWnHAybfB3zupkJrIyEXkAhXAEsyr8g8WNV05RDrK5-4yR0oxosrAmXKziZszLmJzqrd2J7V-oom5S9f3j8J5L4ogTQHQRO8a-kyVLInOwfLlSZCW266ZlbxhVgdUg6fFJIBa96T8kbbNbFcw9zBTbdM5Rzim-fP9k8ut_1n2Tb0Jq96vz_t__YUeFk6EFN9wJWy4i-UvVdIFk2rgEt4Q9Ektd7lU3_h6kE48ISAnpf7tjn6pzRdTMUJjkK2tkisnJz9Wc7n47ErzrgDZcdkcxDL0UZExvo03c7WGwiUzoMS7ZtZYWGr-Tx1LJHf-ZKHMYLGXINcpLBDVw5hLdqeZhv1TIrqJHCoZi4ZbJCNHcs3pIm3PmkJXYRuGB9Jo1v61TKbk8vLypY3hzLDaFA2uphqU8Fr11jHp7NVRzuZYvnFOJKyxL4cV8NL6A6yscqFhyPn1XVLr_W0O15ku3QDQ60XUdxQILnXgJa8vnF5NpiNSWKb8DBQKZZTKt9t83EUIL2kyDrRqQhRu0kne9yMXrJuwDaLmUQejXDnM5QDOjSUN9VLnR_kdRYcL-HQOGUCv1hOhKLtdL_UkeEoNjec6YrrMevzcFFjqTCajlfefs1m9AxJbn0qRv5ob8l1Jl0rBzTiVNUUMV8lWSfx-LkkqCsrWiEay1mBMZi9rtHtd70v-h88__DuXcchxAeE4wROP-A4ur8h8I9-_evPT-_60QY7blfbqt2-3_u-F_VvT3HVDZjd3TuP5v727fpjJJfbZnX-kmQf0AMxwBD3AuEvaP4q8dGNxo-WpgcIkdyA_Pz_NPz1H6_-_uH1T2---qX3Wf8TdhHRBBnFGIXHOEaxOIFxLMFhHENkERsn2YIi3rv5B2gbeRnaBAAA
~~~

but I also found this shorter value, from a fork of my own code:

~~~
H4sIAAAAAAAAAONqZuRqYPQwKfZ0hIJ4U99KkwpPo5Isx0q_Kv8qXx9fgxIXxwrfqlBj3xDXStconzzXLEcDX3Nfg2IXR6OSkMAK36xQ98r0nIpk0wqzeHdXo1wLZ_ckU6OwoDyv5DS3cgsPxzInxwzf8HhLc2cfzxBnx3In3SAnM7NwsI22tgCrOwMMhgAAAA
~~~

https://github.com/mxrch/googleplay/blob/main/research/2022-11-21/autoUpdate.go

I found this which gave me a clue:

<https://github.com/doug-leith/android-protobuf-decoding/blob/e286302dc2c4938f58a4a0f88bc665e1a588222e/decoding_helpers.py#L558-L565>

but it seems with Go its actually `compress/gzip` and not `compress/zlib`. if you run this on the value:

~~~go
package main

import (
   "154.pages.dev/encoding/protobuf"
   "bytes"
   "compress/gzip"
   "encoding/base64"
   "fmt"
   "io"
)

const x_ps_rh = "H4sIAAAAAAAAAONqZuRqYPQwKfZ0hIJ4U99KkwpPo5Isx0q_Kv8qXx9fgxIXxwrfqlBj3xDXStconzzXLEcDX3Nfg2IXR6OSkMAK36xQ98r0nIpk0wqzeHdXo1wLZ_ckU6OwoDyv5DS3cgsPxzInxwzf8HhLc2cfzxBnx3In3SAnM7NwsI22tgCrOwMMhgAAAA"

func main() {
   b, err := base64.RawURLEncoding.DecodeString(x_ps_rh)
   if err != nil {
      panic(err)
   }
   r, err := gzip.NewReader(bytes.NewReader(b))
   if err != nil {
      panic(err)
   }
   defer r.Close()
   b, err = io.ReadAll(r)
   if err != nil {
      panic(err)
   }
   m, err := protobuf.Consume(b)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%#v\n", m)
}
~~~

you get this result:

~~~go
protobuf.Message{
   protobuf.Field{Number:1, Type:2, Value:protobuf.Prefix{
      protobuf.Field{Number:1, Type:2, Value:protobuf.Bytes("H4sIAAAAAAAA_5My4xI2tjAyNzOzMLM0tDAxMzU3MTEyEZLnEjA0M7M0sDA2tTQxMjUGyglxc5x6_GE2m8CGb52VRnJcfFw8HAvBAhMW_97CLITCAwB-RB66WAAAAA==")},
   }},
}
~~~

if you run this on the inner value:

~~~go
package main

import (
   "154.pages.dev/encoding/protobuf"
   "bytes"
   "compress/gzip"
   "encoding/base64"
   "fmt"
   "io"
)

const x_ps_rh_2 = "H4sIAAAAAAAA_5My4xI2tjAyNzOzMLM0tDAxMzU3MTEyEZLnEjA0M7M0sDA2tTQxMjUGyglxc5x6_GE2m8CGb52VRnJcfFw8HAvBAhMW_97CLITCAwB-RB66WAAAAA=="

func main() {
   b, err := base64.URLEncoding.DecodeString(x_ps_rh_2)
   if err != nil {
      panic(err)
   }
   r, err := gzip.NewReader(bytes.NewReader(b))
   if err != nil {
      panic(err)
   }
   defer r.Close()
   b, err = io.ReadAll(r)
   if err != nil {
      panic(err)
   }
   m, err := protobuf.Consume(b)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%#v\n", m)
}
~~~

you get this:

~~~go
protobuf.Message{
   protobuf.Field{Number: 3, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("3827668691846574424")},
      protobuf.Field{Number: 2, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("1669083594253918")},
         protobuf.Field{Number: 2, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(1669083594)},
            protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(253918000)},
         }},
      }},
   }},
   protobuf.Field{Number: 6, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(1669083553)},
            protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(916378000)},
         }},
      }},
      protobuf.Field{Number: 2, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(1669083553)},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(916378000)},
      }},
   }},
}
~~~

I am not sure what all these are, but looks promising.
