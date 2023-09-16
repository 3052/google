package main

import (
   "154.pages.dev/google/acquire"
   "fmt"
)

func main() {
   c, err := acquire.NewClientWithDeviceInfo()
   if err != nil {
      panic(err)
   }
   fmt.Printf("%+v\n", c.AuthData.GsfID)
}
