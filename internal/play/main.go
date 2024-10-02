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

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.BoolVar(&f.acquire, "a", false, "acquire")
   flag.StringVar(
      &play.Device.Abi, "abi", play.Abi[0], strings.Join(play.Abi[1:], " "),
   )
   flag.StringVar(
      &f.auth, "auth", "", "accounts.google.com/embedded/setup/v2/android",
   )
   flag.BoolVar(&f.checkin, "checkin", false, "checkin request")
   flag.StringVar(&f.app.Id, "i", "", "ID")
   flag.BoolVar(&f.leanback, "leanback", false, play.Leanback)
   flag.BoolVar(&f.single, "s", false, "single APK")
   flag.BoolVar(&f.sync, "sync", false, "sync request")
   flag.Uint64Var(&f.app.Version, "v", 0, "version code")
   flag.Parse()
   text.Transport{}.Set(true)
   switch {
   case f.app.Id != "":
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
   case f.auth != "":
      err := f.do_auth()
      if err != nil {
         panic(err)
      }
   case f.checkin:
      err := f.do_checkin()
      if err != nil {
         panic(err)
      }
   case f.sync:
      err := f.do_sync()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}

type flags struct {
   app play.StoreApp
   sync bool
   single bool
   leanback bool
   home string
   checkin bool
   auth string
   acquire bool
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
