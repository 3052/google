package play

import (
   "154.pages.dev/http"
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
   text, err := os.ReadFile(home + "token.txt")
   if err != nil {
      t.Fatal(err)
   }
   http.No_Location()
   head.Set_Authorization(text)
   time.Sleep(time.Second)
   for _, app := range apps {
      data, err := os.ReadFile(home + Platforms[app.platform] + ".bin")
      if err != nil {
         t.Fatal(err)
      }
      head.Set_Device(data)
      fmt.Println(app.doc)
      if err := head.Acquire(app.doc); err != nil {
         t.Fatal(err)
      }
      time.Sleep(99 * time.Millisecond)
   }
}
