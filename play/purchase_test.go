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
      r, err := Raw_Token.Refresh(b)
      if err != nil {
         t.Fatal(err)
      }
      head.Token, err = r.Access()
      if err != nil {
         t.Fatal(err)
      }
   }
   {
      b, err := os.ReadFile(home + "/x86.bin")
      if err != nil {
         t.Fatal(err)
      }
      d, err := Raw_Device.Device(b)
      if err != nil {
         t.Fatal(err)
      }
      head.Device_ID, err = d.ID()
      if err != nil {
         t.Fatal(err)
      }
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
