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
   for _, abi := range Abis {
      fmt.Println(abi)
      Device.Abi = abi
      var checkin GoogleCheckin
      data, err := Device.Checkin(&checkin)
      if err != nil {
         t.Fatal(err)
      }
      err = os.WriteFile(home+"/"+abi+".txt", data, os.ModePerm)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
      err = Device.Sync(&checkin)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}
