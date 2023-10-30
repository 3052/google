package play

import (
   "fmt"
   "os"
   "testing"
   "time"
)

func Test_Sync(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   home += "/google/play/"
   for _, platform := range Platforms {
      var check Checkin
      Phone.Platform = platform
      fmt.Println(platform)
      if err := check.Checkin(Phone); err != nil {
         t.Fatal(err)
      }
      if err := os.WriteFile(home+platform+".bin", check.Raw, 0666); err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
      if err := check.Unmarshal(); err != nil {
         t.Fatal(err)
      }
      if err := check.Sync(Phone); err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}
