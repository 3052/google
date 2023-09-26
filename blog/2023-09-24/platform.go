package main

import (
   "154.pages.dev/protobuf"
   "fmt"
)

func main() {
   {
      var m protobuf.Message
      m.Add_String(11, "x86_64")
      m.Add_String(11, "x86")
      b := m.Append(nil)
      fmt.Printf("%v %q\n", len(b), b)
   }
   {
      var m protobuf.Message
      m.Add_String(11, "armeabi-v7a")
      b := m.Append(nil)
      fmt.Printf("%v %q\n", len(b), b)
   }
}
