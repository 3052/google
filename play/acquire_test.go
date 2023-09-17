package play

import (
   "154.pages.dev/http/option"
   "fmt"
   "os"
   "time"
   "testing"
)

const (
   acquire_doc = "com.chess24.application"
   acquire_version = 10000957
)

func Test_Acquire(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   home += "/google/play"
   var head Header
   head.Set_Agent(false)
   option.No_Location()
   option.Trace()
   {
      b, err := os.ReadFile(home + "/token.txt")
      if err != nil {
         t.Fatal(err)
      }
      if err := head.Set_Authorization(b); err != nil {
         t.Fatal(err)
      }
   }
   fmt.Println("sleep")
   time.Sleep(9 * time.Second)
   if err := head.Acquire(acquire_doc); err != nil {
      t.Fatal(err)
   }
   fmt.Println("sleep")
   time.Sleep(9 * time.Second)
   if err := New_Delivery(head, acquire_doc, acquire_version); err != nil {
      t.Fatal(err)
   }
}
