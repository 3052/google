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
      var checkin GoogleCheckin
      Device.Abi = abi
      fmt.Println(abi)
      time.Sleep(time.Second)
      data, err := checkin.Marshal(&Device)
      if err != nil {
         t.Fatal(err)
      }
      err = os.WriteFile(home+"/"+abi+".txt", data, os.ModePerm)
      if err != nil {
         t.Fatal(err)
      }
      err = checkin.Unmarshal(data)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
      err = Device.Sync(&checkin)
      if err != nil {
         t.Fatal(err)
      }
   }
}
