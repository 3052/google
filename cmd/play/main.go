package main

import (
   "154.pages.dev/google/play"
   "154.pages.dev/http"
   "flag"
   "fmt"
   "strings"
)

type flags struct {
   acquire bool
   code string
   device bool
   platform int64
   single bool
   app play.Application
}

func main() {
   var f flags
   flag.StringVar(&f.code, "c", "", func() string {
      var b strings.Builder
      b.WriteString("oauth_token from ")
      b.WriteString("accounts.google.com/embedded/setup/android")
      return b.String()
   }())
   flag.BoolVar(&f.device, "device", false, "create device")
   flag.Int64Var(&f.platform, "p", 0, fmt.Sprint(play.Platforms))
   flag.BoolVar(&f.single, "s", false, "single APK")
   flag.BoolVar(&f.acquire, "acquire", false, "acquire application")
   flag.StringVar(&f.app.ID, "a", "", "application ID")
   flag.Uint64Var(&f.app.Version, "v", 0, "version code")
   flag.Parse()
   http.No_Location()
   http.Verbose()
   if f.code != "" {
      err := f.do_auth()
      if err != nil {
         panic(err)
      }
   } else {
      play.Phone.Platform = play.Platforms[f.platform]
      switch {
      case f.device:
         err := f.do_device()
         if err != nil {
            panic(err)
         }
      case f.app.ID != "":
         switch {
         case f.acquire:
            err := f.do_acquire()
            if err != nil {
               panic(err)
            }
         case f.app.Version >= 1:
            err := f.do_delivery()
            if err != nil {
               panic(err)
            }
         default:
            details, err := f.do_details()
            if err != nil {
               panic(err)
            }
            fmt.Println(details)
         }
      default:
         flag.Usage()
      }
   }
}
