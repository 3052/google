package acquire

import (
   "154.pages.dev/google/acquire"
   "154.pages.dev/google/play"
   "154.pages.dev/http/option"
   "flag"
   "fmt"
   "os"
   "time"
   "testing"
)

type flags struct {
   device bool
   doc string
   platform int64
   trace bool
   vc uint64
}

func Test_Acquire(t *testing.T) {
   var f flags
   flag.StringVar(&f.doc, "d", "", "doc")
   flag.BoolVar(&f.device, "device", false, "device acquire (experimental)")
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
   case f.device:
      err := f.do_device_acquire(dir, platform)
      if err != nil {
         panic(err)
      }
   case f.doc != "":
      head, err := f.do_header(dir, platform)
      if err != nil {
         panic(err)
      }
      if err := acquire.Acquire(head, f.doc); err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}

func (f flags) do_header(dir, platform string) (*play.Header, error) {
   var head play.Header
   head.Set_Agent(false)
   {
      b, err := os.ReadFile(dir + "/token.txt")
      if err != nil {
         return nil, err
      }
      if err := head.Set_Authorization(b); err != nil {
         return nil, err
      }
   }
   {
      b, err := os.ReadFile(dir + "/" + platform + ".bin")
      if err != nil {
         return nil, err
      }
      if err := head.Set_Device(b); err != nil {
         return nil, err
      }
   }
   return &head, nil
}
