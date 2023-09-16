package main

import (
   "154.pages.dev/google/aurora"
   "154.pages.dev/http/option"
   "os"
)

func main() {
   option.No_Location()
   option.Verbose()
   text, err := new(aurora.Aurora_OSS).Marshal()
   if err != nil {
      panic(err)
   }
   if err := os.WriteFile("aurora.json", text, 0666); err != nil {
      panic(err)
   }
}
