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
   for _, each := range Abi {
      fmt.Println(each)
      Device.Abi = each
      data, err := Device.GoogleCheckin()
      if err != nil {
         t.Fatal(err)
      }
      err = os.WriteFile(home+"/"+each+".bin", data, 0666)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
      var checkin GoogleCheckin
      err = checkin.Unmarshal(data)
      if err != nil {
         t.Fatal(err)
      }
      err = Device.Sync(checkin)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}
