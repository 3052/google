package main

import (
   "154.pages.dev/google/play"
   "154.pages.dev/http/option"
   "flag"
   "os"
)

type flags struct {
   device_acquire bool
   doc string
   platform int64
   trace bool
   vc uint64
}

func main() {
   var f flags
   flag.StringVar(&f.doc, "d", "", "doc")
   flag.BoolVar(&f.device_acquire, "device_acquire", false, "device acquire (experimental)")
   flag.Int64Var(&f.platform, "p", 0, play.Platforms.String())
   flag.BoolVar(&f.trace, "t", false, "print full HTTP requests")
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
   option.No_Location()
   if f.trace {
      option.Trace()
   } else {
      option.Verbose()
   }
   platform := play.Platforms[f.platform]
   switch {
   case f.device_acquire:
      err := f.do_device_acquire(dir, platform)
      if err != nil {
         panic(err)
      }
   case f.doc != "":
      head, err := f.do_header(dir, platform)
      if err != nil {
         panic(err)
      }
      if err := head.Acquire(f.doc); err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}
