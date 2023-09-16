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
   single bool
   vc uint64
   acquire bool
   home string
}

func main() {
   var (
      f flags
      err error
   )
   f.home, err = os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   f.home += "/google/play/"
   flag.BoolVar(&f.acquire, "a", false, "acquire")
   {
      var b strings.Builder
      b.WriteString("oauth_token from ")
      b.WriteString("accounts.google.com/embedded/setup/android")
      flag.StringVar(&f.code, "c", "", b.String())
   }
   flag.StringVar(&f.doc, "d", "", "doc")
   flag.BoolVar(&f.device, "device", false, "create device")
   {
      play.Phone.Native_Platform = "x86"
      flag.Func("p", fmt.Sprint(play.Native_Platforms), func(s string) error {
         play.Phone.Native_Platform = play.Native_Platforms[s]
         return nil
      })
   }
   flag.BoolVar(&f.single, "s", false, "single APK")
   flag.Uint64Var(&f.vc, "v", 0, "version code")
   flag.Parse()
   option.No_Location()
   option.Verbose()
   switch {
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
   case f.doc != "":
      switch {
      case f.acquire:
         err := f.do_acquire()
         if err != nil {
            panic(err)
         }
      case f.vc >= 1:
         err := f.do_delivery()
         if err != nil {
            panic(err)
         }
      default:
         err := f.do_details()
         if err != nil {
            panic(err)
         }
      }
   default:
      flag.Usage()
   }
}
