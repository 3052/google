package main

import (
   "154.pages.dev/google/play"
   "fmt"
   "os"
   "time"
)

func (f flags) do_device_acquire(dir, platform string) error {
   data, err := play.Checkin_Acquire(platform)
   if err != nil {
      return err
   }
   err = os.WriteFile(dir + "/" + platform + ".bin", data, 0666)
   if err != nil {
      return err
   }
   fmt.Printf("Sleeping %v for server to process\n", play.Sleep)
   time.Sleep(play.Sleep)
   return nil
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
