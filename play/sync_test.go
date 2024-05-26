package play

import (
   "fmt"
   "os"
   "testing"
   "time"
)

func TestSync(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   home += "/google-play"
   for _, each := range ABI {
      var checkin GoogleCheckin
      fmt.Println(each)
      Device.ABI = each
      if err := checkin.Checkin(Device); err != nil {
         t.Fatal(err)
      }
      err := os.WriteFile(home+"/"+each+".bin", checkin.Data, 0666)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
      if err := checkin.Unmarshal(); err != nil {
         t.Fatal(err)
      }
      if err := checkin.Sync(Device); err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}
