package main

import (
   "154.pages.dev/google/play"
   "flag"
   "fmt"
   "os"
)

type flags struct {
   device bool
   doc string
   email string
   file string
   passwd string
   platform int64
   purchase bool
   single bool
   vc uint64
}

func main() {
   var f flags
   flag.StringVar(&f.doc, "d", "", "doc")
   flag.BoolVar(&f.device, "device", false, "create device")
   flag.StringVar(&f.email, "email", "", "your email")
   flag.StringVar(&f.file, "f", "", "password file")
   flag.Int64Var(&f.platform, "p", 0, play.Platforms.String())
   flag.StringVar(&f.passwd, "passwd", "", "password")
   flag.BoolVar(&f.purchase, "purchase", false, "purchase request")
   flag.BoolVar(&f.single, "s", false, "single APK")
   flag.Uint64Var(&f.vc, "v", 0, "version code")
   flag.Parse()
   dir, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   dir += "/google/play"
   if err := os.MkdirAll(dir, os.ModePerm); err != nil {
      panic(err)
   }
   if f.passwd != "" || f.file != "" {
      err := f.do_auth(dir)
      if err != nil {
         panic(err)
      }
   } else {
      platform := play.Platforms[f.platform]
      if f.device {
         err := f.do_device(dir, platform)
         if err != nil {
            panic(err)
         }
      } else if f.doc != "" {
         head, err := f.do_header(dir, platform)
         if err != nil {
            panic(err)
         }
         if f.purchase {
            err := head.Purchase(f.doc)
            if err != nil {
               panic(err)
            }
         } else if f.vc >= 1 {
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
