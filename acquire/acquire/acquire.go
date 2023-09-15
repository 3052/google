package main

import (
   "154.pages.dev/google/acquire"
   "154.pages.dev/google/play"
   "154.pages.dev/http/option"
   "flag"
   "fmt"
   "os"
   "time"
)

func main() {
   var doc string
   flag.StringVar(&doc, "d", "", "doc")
   var version_code uint64
   flag.StringVar(&acquire.Device_ID, "id", "", "device ID")
   flag.Uint64Var(&version_code, "v", 0, "version code")
   flag.Parse()
   if doc == "" {
      flag.Usage()
      return
   }
   if acquire.Device_ID == "" {
      flag.Usage()
      return
   }
   head := new(play.Header)
   head.Set_Agent(false)
   option.No_Location()
   option.Trace()
   {
      s, err := os.UserHomeDir()
      if err != nil {
         panic(err)
      }
      b, err := os.ReadFile(s + "/google/play/token.txt")
      if err != nil {
         panic(err)
      }
      if err := head.Set_Authorization(b); err != nil {
         panic(err)
      }
   }
   fmt.Println("sleep")
   time.Sleep(9 * time.Second)
   if err := acquire.Acquire(head, doc); err != nil {
      panic(err)
   }
   fmt.Println("sleep")
   time.Sleep(9 * time.Second)
   if err := acquire.New_Delivery(head, doc, version_code); err != nil {
      panic(err)
   }
}
