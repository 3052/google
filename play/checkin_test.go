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
   for _, each := range Abi {
      fmt.Println(each)
      Device.Abi = each
      data, err := Device.GoogleCheckin()
      if err != nil {
         t.Fatal(err)
      }
      os.WriteFile(home+"/"+each+".bin", data, 0666)
      time.Sleep(time.Second)
   }
}
