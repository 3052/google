package main

import (
   "flag"
   "fmt"
)

var platforms = map[string]string{
   "0": "x86",
   "1": "armeabi-v7a",
   "2": "arm64-v8a",
}

func main() {
   var platform string
   flag.Func("p", fmt.Sprint(platforms), func(s string) error {
      platform = platforms[s]
      return nil
   })
   flag.Parse()
   flag.Usage()
   fmt.Printf("%q\n", platform)
}
