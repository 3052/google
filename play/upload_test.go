package play

import (
   "154.pages.dev/http/option"
   "fmt"
   "os"
   "testing"
   "time"
)

func Test_Upload(t *testing.T) {
   option.No_Location()
   option.Verbose()
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   home += "/google/play/"
   var head Header
   // agent
   head.Set_Agent(false)
   // authorization
   {
      b, err := os.ReadFile(home + "token.txt")
      if err != nil {
         t.Fatal(err)
      }
      head.Set_Authorization(b)
   }
   // device
   {
      b, err := Phone.Checkin()
      if err != nil {
         t.Fatal(err)
      }
      head.Set_Device(b)
   }
   Phone.Native_Platform = "x86"
   if err := head.Upload(Phone); err != nil {
      t.Fatal(err)
   }
   fmt.Println(head.Device())
   time.Sleep(9*time.Second)
}
