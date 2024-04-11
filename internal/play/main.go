package main

import (
   "154.pages.dev/google/play"
   "154.pages.dev/log"
   "flag"
   "fmt"
   "os"
   "path/filepath"
   "strings"
)

type flags struct {
   acquire bool
   app play.StoreApp
   code string
   device bool
   home string
   platform play.ABI
   single bool
   v log.Level
}

func (f *flags) New() error {
   var err error
   f.home, err = os.UserHomeDir()
   if err != nil {
      return err
   }
   f.home = filepath.ToSlash(f.home) + "/google-play"
   return nil
}

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.BoolVar(&f.acquire, "a", false, "acquire")
   flag.Var(&f.platform, "b", fmt.Sprint(play.ABIs))
   flag.BoolVar(&f.device, "d", false, "checkin and sync device")
   flag.StringVar(&f.app.ID, "i", "", "ID")
   flag.TextVar(&f.v.Level, "level", f.v.Level, "level")
   {
      var b strings.Builder
      b.WriteString("oauth_token from ")
      b.WriteString("accounts.google.com/embedded/setup/v2/android")
      flag.StringVar(&f.code, "o", "", b.String())
   }
   flag.BoolVar(&f.single, "s", false, "single APK")
   flag.Uint64Var(&f.app.Version, "v", 0, "version code")
   flag.Parse()
   f.v.Set()
   log.Transport{}.Set()
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
