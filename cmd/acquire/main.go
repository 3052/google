package main

import (
   "154.pages.dev/google/play"
   "154.pages.dev/http/option"
   "flag"
   "fmt"
   "net/http"
   "os"
   "strings"
   "time"
)

type flags struct {
   code string
   device bool
   doc string
   single bool
   trace bool
   vc uint64
   acquire bool
   platform struct {
      k int64
      v string
   }
}

func main() {
   var f flags
   flag.BoolVar(&f.acquire, "a", false, "acquire")
   {
      var b strings.Builder
      b.WriteString("oauth_token from ")
      b.WriteString("accounts.google.com/embedded/setup/android")
      flag.StringVar(&f.code, "c", "", b.String())
   }
   flag.StringVar(&f.doc, "d", "", "doc")
   flag.BoolVar(&f.device, "device", false, "create device")
   flag.Int64Var(&f.platform, "p", 0, play.Platforms.String())
   flag.BoolVar(&f.single, "s", false, "single APK")
   flag.BoolVar(&f.trace, "t", false, "print full HTTP requests")
   flag.Uint64Var(&f.vc, "v", 0, "version code")
   flag.Parse()
   dir, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   dir += "/google/play"
   if err := os.MkdirAll(dir, os.ModePerm); err != nil {
      panic(err)
   }
   option.No_Location()
   if f.trace {
      option.Trace()
   } else {
      option.Verbose()
   }
   platform := play.Platforms[f.platform]
   switch {
   case f.acquire:
      err := f.do_acquire(platform)
      if err != nil {
         panic(err)
      }
   case f.code != "":
      err := f.do_auth(dir)
      if err != nil {
         panic(err)
      }
   case f.doc != ""
      if f.vc >= 1 {
         err := f.do_delivery(head)
         if err != nil {
            panic(err)
         }
      } else {
         detail, err := head.Details(f.doc)
         if err != nil {
            panic(err)
         }
         fmt.Println(detail)
      }
   case f.device:
      err := f.do_device(dir, platform)
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}
