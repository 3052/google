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
   trace bool
   vc uint64
   acquire bool
   platform string
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
   f.home += "/google/play"
   os.MkdirAll(f.home, os.ModePerm)
   flag.BoolVar(&f.acquire, "a", false, "acquire")
   {
      var b strings.Builder
      b.WriteString("oauth_token from ")
      b.WriteString("accounts.google.com/embedded/setup/android")
      flag.StringVar(&f.code, "c", "", b.String())
   }
   flag.StringVar(&f.doc, "d", "", "doc")
   flag.BoolVar(&f.device, "device", false, "create device")
   flag.Func("p", fmt.Sprint(play.Native_Platforms), func(s string) error {
      f.platform = play.Native_Platforms[s]
      return nil
   })
   flag.BoolVar(&f.single, "s", false, "single APK")
   flag.BoolVar(&f.trace, "t", false, "print full HTTP requests")
   flag.Uint64Var(&f.vc, "v", 0, "version code")
   flag.Parse()
   option.No_Location()
   if f.trace {
      option.Trace()
   } else {
      option.Verbose()
   }
   switch {
   case f.acquire:
      err := f.do_acquire()
      if err != nil {
         panic(err)
      }
   case f.code != "":
      err := f.do_auth(dir)
      if err != nil {
         panic(err)
      }
   case f.doc != "":
      if f.vc >= 1 {
         err := f.do_delivery()
         if err != nil {
            panic(err)
         }
      } else {
         err := f.do_details()
         if err != nil {
            panic(err)
         }
      }
   case f.device:
      err := f.do_device(dir)
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}
