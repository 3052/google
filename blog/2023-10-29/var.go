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
   flag.Var(&p, "p", fmt.Sprint(platforms))
   flag.Parse()
   flag.Usage()
   fmt.Printf("%q\n", p)
}

type platform string

func (p platform) String() string {
   return string(p)
}

func (p *platform) Set(s string) error {
   *p = platforms[s]
   return nil
}

