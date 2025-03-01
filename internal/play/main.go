package main

import (
   "41.neocities.org/google/play"
   "flag"
   "fmt"
   "log"
   "net/http"
   "os"
   "path/filepath"
   "strings"
)

func (transport) RoundTrip(req *http.Request) (*http.Response, error) {
   log.Println(req.Method, req.URL)
   return http.DefaultTransport.RoundTrip(req)
}

type transport struct{}

func main() {
   log.SetFlags(log.Ltime)
   http.DefaultClient.Transport = transport{}
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.BoolVar(&f.acquire, "a", false, "acquire")
   flag.StringVar(
      &play.DefaultDevice.Abi,
      "abi",
      play.Abis[0],
      strings.Join(play.Abis[1:], " "),
   )
   flag.BoolVar(&f.checkin, "checkin", false, "checkin request")
   flag.StringVar(&f.app.Id, "i", "", "ID")
   flag.BoolVar(&f.leanback, "leanback", false, play.Leanback)
   flag.BoolVar(&f.single, "s", false, "single APK")
   flag.BoolVar(&f.sync, "sync", false, "sync request")
   flag.StringVar(
      &f.token, "token", "", "accounts.google.com/embedded/setup/v2/android",
   )
   flag.Uint64Var(&f.app.Version, "v", 0, "version code")
   flag.Parse()
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
   case f.token != "":
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
   acquire  bool
   checkin  bool
   home     string
   leanback bool
   single   bool
   sync     bool
   token    string
   app      play.App
}
