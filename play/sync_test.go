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
      fmt.Println(platform)
      Pixel_6.Platform = platform
      if err := check.Checkin(Pixel_6); err != nil {
         t.Fatal(err)
      }
      err := os.WriteFile(home+platform+".bin", check.Raw, 0666)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
      if err := check.Unmarshal(); err != nil {
         t.Fatal(err)
      }
      if err := check.Sync(Pixel_6); err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}
