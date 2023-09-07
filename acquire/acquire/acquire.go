package main

import (
   "154.pages.dev/google/play"
   "154.pages.dev/http/option"
   "acquire"
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
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   home += "/google/play"
   head := new(play.Header)
   head.Set_Agent(false)
   option.No_Location()
   option.Trace()
   {
      b, err := os.ReadFile(home + "/token.txt")
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
   return
   fmt.Println("sleep")
   time.Sleep(9 * time.Second)
   if err := acquire.New_Delivery(head, doc, version_code); err != nil {
      panic(err)
   }
}
