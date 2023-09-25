package play

import (
   "154.pages.dev/http/option"
   "fmt"
   "os"
   "testing"
   "time"
)

func Test_Acquire(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   home += "/google/play/"
   var head Header
   head.Set_Agent(true)
   option.No_Location()
   {
      b, err := os.ReadFile(home + "token.txt")
      if err != nil {
         t.Fatal(err)
      }
      head.Set_Authorization(b)
   }
   for _, app := range apps {
      {
         b, err := os.ReadFile(home + Platforms[app.platform] + ".bin")
         if err != nil {
            t.Fatal(err)
         }
         head.Set_Device(b)
      }
      fmt.Println(app.doc)
      err := head.Acquire(app.doc)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}
