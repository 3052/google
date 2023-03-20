package main

import (
   "2a.pages.dev/googleplay"
   "flag"
   "fmt"
   "os"
)

type flags struct {
   app string
   device bool
   email string
   password string
   platform int64
   purchase bool
   single bool
   version uint64
   verbose bool
}

func main() {
   var f flags
   // a
   flag.StringVar(&f.app, "a", "", "app")
   // device
   flag.BoolVar(&f.device, "device", false, "create device")
   // email
   flag.StringVar(&f.email, "email", "", "your email")
   // p
   flag.Int64Var(&f.platform, "p", 0, googleplay.Platforms.String())
   // password
   flag.StringVar(&f.password, "password", "", "your password")
   // purchase
   flag.BoolVar(&f.purchase, "purchase", false, "purchase app")
   // s
   flag.BoolVar(&f.single, "s", false, "single APK")
   // v
   flag.Uint64Var(&f.version, "v", 0, "app version")
   // verbose
   flag.BoolVar(&f.verbose, "verbose", false, "verbose")
   flag.Parse()
   if f.verbose {
      googleplay.Client.Log_Level = 2
   }
   dir, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   dir += "/googleplay"
   os.Mkdir(dir, os.ModePerm)
   if f.password != "" {
      err := f.do_auth(dir)
      if err != nil {
         panic(err)
      }
   } else {
      platform := googleplay.Platforms[f.platform]
      if f.device {
         err := do_device(dir, platform)
         if err != nil {
            panic(err)
         }
      } else if f.app != "" {
         head, err := f.do_header(dir, platform)
         if err != nil {
            panic(err)
         }
         if f.purchase {
            err := head.Purchase(f.app)
            if err != nil {
               panic(err)
            }
         } else if f.version >= 1 {
            err := f.do_delivery(head)
            if err != nil {
               panic(err)
            }
         } else {
            detail, err := f.do_details(head)
            if err != nil {
               panic(err)
            }
            fmt.Println(string(detail))
         }
      } else {
         flag.Usage()
      }
   }
}
