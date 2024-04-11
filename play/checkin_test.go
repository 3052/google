package play

import (
   "fmt"
   "testing"
   "time"
)

func TestCheckin(t *testing.T) {
   for _, platform := range ABIs {
      fmt.Println(platform)
      Phone.ABI = platform
      var checkin GoogleCheckin
      err := checkin.Checkin(Phone)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}
