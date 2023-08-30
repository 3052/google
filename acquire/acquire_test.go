package acquire

import (
   "154.pages.dev/google/play"
   "154.pages.dev/http/option"
   "fmt"
   "os"
   "testing"
   "time"
)

const (
   doc = "com.google.android.apps.mapslite"
   version_code = 160
)

func Test_Acquire(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   home += "/google/play"
   head := new(play.Header)
   head.Set_Agent(false)
   option.No_Location()
   option.Verbose()
   {
      b, err := os.ReadFile(home + "/token.txt")
      if err != nil {
         t.Fatal(err)
      }
      if err := head.Set_Authorization(b); err != nil {
         t.Fatal(err)
      }
   }
   if err := acquire_2(head, doc, version_code); err != nil {
      t.Fatal(err)
   }
   fmt.Println("sleep")
   time.Sleep(9 * time.Second)
   if err := New_Delivery(head, doc, version_code); err != nil {
      t.Fatal(err)
   }
}
