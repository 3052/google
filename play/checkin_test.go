package play

import (
   "fmt"
   "os"
   "testing"
   "time"
)

func TestCheckin(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   home += "/google-play"
   for _, each := range ABI {
      fmt.Println(each)
      Device.ABI = each
      var checkin GoogleCheckin
      err := checkin.Checkin(Device)
      if err != nil {
         t.Fatal(err)
      }
      os.WriteFile(home + "/" + each + ".bin", checkin.Data, 0666)
      time.Sleep(time.Second)
   }
}
