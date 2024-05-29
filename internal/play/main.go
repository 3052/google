package main

import (
   "154.pages.dev/google/play"
   "154.pages.dev/text"
   "flag"
   "fmt"
   "os"
   "path/filepath"
   "strings"
)

func (f *flags) New() error {
   var err error
   f.home, err = os.UserHomeDir()
   if err != nil {
      return err
   }
   f.home = filepath.ToSlash(f.home) + "/google-play"
   return nil
}

type flags struct {
   acquire bool
   app play.StoreApp
   code string
   checkin bool
   home string
   single bool
   log text.LogLevel
   leanback bool
}

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.BoolVar(&f.acquire, "a", false, "acquire")
   flag.StringVar(
      &play.Device.ABI, "b", play.ABI[0], strings.Join(play.ABI[1:], " "),
   )
   flag.Uint64Var(&f.app.Version, "c", 0, "version code")
   flag.BoolVar(&f.checkin, "d", false, "checkin and sync device")
   flag.StringVar(&f.app.ID, "i", "", "app ID")
   {
      var b strings.Builder
      b.WriteString("oauth_token from ")
      b.WriteString("accounts.google.com/embedded/setup/v2/android")
      flag.StringVar(&f.code, "o", "", b.String())
   }
   flag.BoolVar(&f.single, "s", false, "single APK")
   flag.BoolVar(&f.leanback, "t", false, play.Leanback)
   flag.TextVar(&f.log.Level, "v", f.log.Level, "level")
   flag.Parse()
   f.log.Set()
   f.log.SetTransport(true)
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
   case f.checkin:
      err := f.do_device()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}
