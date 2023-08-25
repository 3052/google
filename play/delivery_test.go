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
   token, err := os.ReadFile(home + "/google/play/token.txt")
   if err != nil {
      t.Fatal(err)
   }
   device, err := os.ReadFile(home + "/google/play/x86.bin")
   if err != nil {
      t.Fatal(err)
   }
   head, err := New_Header(token, device, true)
   if err != nil {
      t.Fatal(err)
   }
   del, err := head.Delivery("com.google.android.youtube", 1524221376)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", del)
}
