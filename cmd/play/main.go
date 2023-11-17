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
   app play.Application
   code string
   device bool
   single bool
   platform play.Platform
}

func main() {
   var f flags
   flag.StringVar(&f.app.ID, "a", "", "application ID")
   flag.BoolVar(&f.acquire, "acquire", false, "acquire application")
   flag.BoolVar(&f.device, "d", false, "checkin and sync device")
   flag.StringVar(&f.code, "o", "", func() string {
      var b strings.Builder
      b.WriteString("oauth_token from ")
      b.WriteString("accounts.google.com/embedded/setup/v2/android")
      return b.String()
   }())
   flag.BoolVar(&f.single, "s", false, "single APK")
   flag.Uint64Var(&f.app.Version, "v", 0, "version code")
   flag.Var(&f.platform, "p", fmt.Sprint(play.Platforms))
   flag.Parse()
   http.No_Location()
   http.Verbose()
   switch {
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
   case f.code != "":
      err := f.do_auth()
      if err != nil {
         panic(err)
      }
   case f.device:
      err := f.do_device()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}
