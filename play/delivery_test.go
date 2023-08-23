package play

import (
   "fmt"
   "os"
   "testing"
)

func Test_Delivery(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   var head Header
   {
      b, err := os.ReadFile(home + "/google/play/token.txt")
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
      b, err := os.ReadFile(home + "/google/play/x86.bin")
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
   del, err := head.Delivery("com.google.android.youtube", 1524221376)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", del)
}
