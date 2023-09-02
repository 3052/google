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
   flag.Uint64Var(&version_code, "v", 0, "version code")
   flag.Parse()
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   home += "/google/play"
   head := new(play.Header)
   head.Set_Agent(false)
   option.No_Location()
   option.Verbose()
   {
      b, err := os.ReadFile(home + "/token.txt")
      if err != nil {
         panic(err)
      }
      if err := head.Set_Authorization(b); err != nil {
         panic(err)
      }
   }
   if err := acquire.Acquire_2(head, doc, version_code); err != nil {
      panic(err)
   }
   fmt.Println("sleep")
   //pass
   //time.Sleep(69*time.Second)
   
   time.Sleep(59*time.Second)
   
   if err := acquire.New_Delivery(head, doc, version_code); err != nil {
      panic(err)
   }
}
