package main

import (
   "154.pages.dev/google/play"
   "154.pages.dev/http/option"
   "flag"
   "fmt"
   "os"
   "strings"
)

type flags struct {
   code string
   device bool
   doc string
   platform int64
   purchase bool
   single bool
   trace bool
   vc uint64
}

func main() {
   var f flags
   {
      var b strings.Builder
      b.WriteString("oauth_token from ")
      b.WriteString("accounts.google.com/embedded/setup/android")
      flag.StringVar(&f.code, "c", "", b.String())
   }
   flag.StringVar(&f.doc, "d", "", "doc")
   flag.BoolVar(&f.device, "device", false, "create device")
   flag.Int64Var(&f.platform, "p", 0, play.Platforms.String())
   flag.BoolVar(&f.purchase, "purchase", false, "purchase request")
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
   if f.code != "" {
      err := f.do_auth(dir)
      if err != nil {
         panic(err)
      }
   } else {
      platform := play.Platforms[f.platform]
      if f.device {
         err := f.do_device(dir, platform)
         if err != nil {
            panic(err)
         }
      } else if f.doc != "" {
         head, err := f.do_header(dir, platform)
         if err != nil {
            panic(err)
         }
         if f.purchase {
            err := head.Purchase(f.doc)
            if err != nil {
               panic(err)
            }
         } else if f.vc >= 1 {
            err := f.do_delivery(head)
            if err != nil {
               panic(err)
            }
         } else {
            detail, err := f.do_details(head)
            if err != nil {
               panic(err)
            }
            fmt.Println(string(detail))
         }
      } else {
         flag.Usage()
      }
   }
}
