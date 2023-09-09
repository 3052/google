package play

import (
   "154.pages.dev/http/option"
   "os"
   "time"
   "testing"
)

const (
   acquire_device = "38e1364d1c418da6"
   acquire_doc = "com.chess24.application"
   acquire_version = 10000957
)

func Test_Acquire(t *testing.T) {
   option.No_Location()
   option.Trace()
   var head Header
   // agent
   head.Set_Agent(false)
   // authorization
   {
      s, err := os.UserHomeDir()
      if err != nil {
         t.Fatal(err)
      }
      b, err := os.ReadFile(s + "/google/play/token.txt")
      if err != nil {
         t.Fatal(err)
      }
      if err := head.Set_Authorization(b); err != nil {
         t.Fatal(err)
      }
   }
   // device
   head.Device = func() (string, string) {
      return "X-DFE-Device-ID", acquire_device
   }
   if err := head.Acquire(acquire_doc); err != nil {
      t.Fatal(err)
   }
   time.Sleep(9 * time.Second)
   if _, err := head.Delivery(acquire_doc, acquire_version); err != nil {
      t.Fatal(err)
   }
}
