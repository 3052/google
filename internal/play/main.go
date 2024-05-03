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
   v log.Level
   abi string
   feature play.Feature
}

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.BoolVar(&f.acquire, "a", false, "acquire")
   flag.StringVar(
      &f.abi, "b", play.ABIs[0], strings.Join(play.ABIs[1:], " "),
   )
   flag.BoolVar(&f.checkin, "c", false, "checkin and sync device")
   {
      var b strings.Builder
      for i, feature := range Features[1:] {
         if i >= 1 {
            b.WriteByte('\n')
         }
         text, err := feature.MarshalText()
         if err != nil {
            panic(err)
         }
         b.Write(text)
      }
      flag.TextVar(&f.feature, "d", play.Features[0], b.String())
   }
   flag.StringVar(&f.app.ID, "i", "", "app ID")
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
   case f.checkin:
      err := f.do_device()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}
