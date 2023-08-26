package play

import (
   "fmt"
   "os"
   "testing"
   "time"
)

func Test_Purchase(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   home += "/google/play"
   var head Header
   {
      b, err := os.ReadFile(home + "/token.txt")
      if err != nil {
         t.Fatal(err)
      }
      head.Set_Authorization(b)
   }
   {
      b, err := os.ReadFile(home + "/x86.bin")
      if err != nil {
         t.Fatal(err)
      }
      head.Set_Device(b)
   }
   for _, app := range apps {
      fmt.Println(app)
      err := head.Purchase(app.doc)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(99 * time.Millisecond)
   }
}
