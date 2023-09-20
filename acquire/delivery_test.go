package acquire

import (
   "154.pages.dev/google/aurora"
   "154.pages.dev/google/play"
   "154.pages.dev/http/option"
   "fmt"
   "os"
   "testing"
   "time"
)

func Test_Acquire(t *testing.T) {
   var auth aurora.Aurora_OSS
   {
      b, err := os.ReadFile("aurora.json")
      if err != nil {
         t.Fatal(err)
      }
      if err := auth.Unmarshal(b); err != nil {
         t.Fatal(err)
      }
   }
   var head play.Header
   head.Authorization = func() (string, string) {
      return "Authorization", "Bearer " + auth.Auth_Token
   }
   option.No_Location()
   option.Verbose()
   app := x86_apps[2]
   err := Acquire(head, app.doc, app.version)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println("sleep")
   time.Sleep(9 * time.Second)
   if err := Delivery(head, app.doc, app.version); err != nil {
      t.Fatal(err)
   }
}
