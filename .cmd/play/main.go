package main

import (
   "154.pages.dev/google/play"
   "154.pages.dev/log"
   "flag"
   "fmt"
   "strings"
)

type flags struct {
   acquire bool
   app play.Application
   code string
   device bool
   platform play.Platform
   single bool
   level log.Level
}

func main() {
   var f flags
   flag.StringVar(&f.app.ID, "a", "", "application ID")
   flag.BoolVar(&f.acquire, "acquire", false, "acquire application")
   flag.BoolVar(&f.device, "d", false, "checkin and sync device")
   {
      var b strings.Builder
      b.WriteString("oauth_token from ")
      b.WriteString("accounts.google.com/embedded/setup/v2/android")
      flag.StringVar(&f.code, "o", "", b.String())
   }
   flag.Var(&f.platform, "p", fmt.Sprint(play.Platforms))
   flag.BoolVar(&f.single, "s", false, "single APK")
   flag.Uint64Var(&f.app.Version, "v", 0, "version code")
   flag.TextVar(&f.level, "level", f.level, "level")
   flag.Parse()
   log.Set_Transport(0)
   log.Set_Logger(f.level)
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
