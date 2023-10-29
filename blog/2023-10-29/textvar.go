package main

import (
   "flag"
   "fmt"
)

var platforms = map[string]platform{
   "0": "x86",
   "1": "armeabi-v7a",
   "2": "arm64-v8a",
}

func main() {
   var p platform
   flag.TextVar(&p, "p", p, fmt.Sprint(platforms))
   flag.Parse()
   flag.Usage()
   fmt.Printf("%q\n", p)
}

type platform string

func (p platform) MarshalText() ([]byte, error) {
   return []byte(p), nil
}

func (p *platform) UnmarshalText(b []byte) error {
   *p = platforms[string(b)]
   return nil
}
