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
   head.Auth = make(Auth)
   {
      b, err := os.ReadFile(home + "/google/play/auth.txt")
      if err != nil {
         t.Fatal(err)
      }
      head.Auth.UnmarshalText(b)
   }
   {
      b, err := os.ReadFile(home + "/google/play/x86.bin")
      if err != nil {
         t.Fatal(err)
      }
      head.Device.UnmarshalBinary(b)
   }
   del, err := head.Delivery("com.google.android.youtube", 1524221376)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", del)
}
