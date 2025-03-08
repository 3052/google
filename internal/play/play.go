package main

import (
   "41.neocities.org/google/play"
   "41.neocities.org/x/progress"
   "flag"
   "fmt"
   "log"
   "net/http"
   "os"
   "path/filepath"
   "strings"
)

func (f *flags) do_delivery() error {
   var checkin play.Checkin
   auth, err := f.client(&checkin)
   if err != nil {
      return err
   }
   deliver, err := auth.Delivery(checkin, &f.app, f.single)
   if err != nil {
      return err
   }
   for obb := range deliver.Obb() {
      err := download(
         obb.Url(), f.app.Obb(obb.Field1()),
      )
      if err != nil {
         return err
      }
   }
   for apk := range deliver.Apk() {
      err := download(
         apk.Url(), f.app.Apk(apk.Field1()),
      )
      if err != nil {
         return err
      }
   }
   err = download(deliver.Url(), f.app.Apk(""))
   if err != nil {
      return err
   }
   return nil
}

func (f *flags) do_acquire() error {
   var checkin play.Checkin
   auth, err := f.client(&checkin)
   if err != nil {
      return err
   }
   return auth.Acquire(checkin, f.app.Id)
}

func (f *flags) device_path() string {
   var b strings.Builder
   b.WriteString(f.home)
   b.WriteByte('/')
   b.WriteString(play.DefaultDevice.Abi)
   if f.leanback {
      b.WriteString("-leanback")
   }
   return b.String()
}

func (f *flags) client(checkin *play.Checkin) (*play.Auth, error) {
   data, err := os.ReadFile(f.home + "/Token")
   if err != nil {
      return nil, err
   }
   var token play.Token
   err = token.Unmarshal(data)
   if err != nil {
      return nil, err
   }
   auth, err := token.Auth()
   if err != nil {
      return nil, err
   }
   data, err = os.ReadFile(f.device_path())
   if err != nil {
      return nil, err
   }
   err = checkin.Unmarshal(data)
   if err != nil {
      return nil, err
   }
   return auth, nil
}
func (transport) RoundTrip(req *http.Request) (*http.Response, error) {
   log.Println(req.Method, req.URL)
   return http.DefaultTransport.RoundTrip(req)
}

type transport struct{}

func (f *flags) New() error {
   var err error
   f.home, err = os.UserHomeDir()
   if err != nil {
      return err
   }
   f.home = filepath.ToSlash(f.home) + "/google/play"
   return nil
}

type flags struct {
   acquire  bool
   app      play.App
   checkin  bool
   home     string
   leanback bool
   single   bool
   sync     bool
   token    string
}

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

func (f *flags) do_details() (*play.Details, error) {
   var checkin play.Checkin
   auth, err := f.client(&checkin)
   if err != nil {
      return nil, err
   }
   return auth.Details(checkin, f.app.Id, f.single)
}

func (f *flags) write_file(name string, data []byte) error {
   log.Println(name)
   return os.WriteFile(name, data, os.ModePerm)
}

func (f *flags) do_auth() error {
   data, err := play.NewToken(f.token)
   if err != nil {
      return err
   }
   return f.write_file(f.home + "/Token", data)
}

func (f *flags) do_checkin() error {
   if f.leanback {
      play.DefaultDevice.Feature = append(
         play.DefaultDevice.Feature, play.Leanback,
      )
   }
   data, err := play.DefaultDevice.Checkin()
   if err != nil {
      return err
   }
   return f.write_file(f.device_path(), data)
}

func (f *flags) do_sync() error {
   data, err := os.ReadFile(f.device_path())
   if err != nil {
      return err
   }
   var checkin play.Checkin
   err = checkin.Unmarshal(data)
   if err != nil {
      return err
   }
   if f.leanback {
      play.DefaultDevice.Feature = append(
         play.DefaultDevice.Feature, play.Leanback,
      )
   }
   return play.DefaultDevice.Sync(checkin)
}

func download(address, name string) error {
   file, err := os.Create(name)
   if err != nil {
      return err
   }
   defer file.Close()
   resp, err := http.Get(address)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   var bytes progress.Bytes
   bytes.Set(resp)
   _, err = file.ReadFrom(&bytes)
   if err != nil {
      return err
   }
   return nil
}
