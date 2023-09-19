package acquire

import (
   "154.pages.dev/google/play"
   "fmt"
   "os"
   "testing"
)

func Test_Delivery(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   var head play.Header
   head.Set_Agent(true)
   {
      b, err := os.ReadFile(home + "/google/play/token.txt")
      if err != nil {
         t.Fatal(err)
      }
      head.Set_Authorization(b)
   }
   {
      b, err := os.ReadFile(home + "/google/play/x86.bin")
      if err != nil {
         t.Fatal(err)
      }
      head.Set_Device(b)
   }
   del, err := head.Delivery("com.google.android.youtube", 1524221376)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", del)
}
