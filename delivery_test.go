package googleplay

import (
   "2a.pages.dev/rosso/http"
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
   head.Open_Auth(home + "/googleplay/auth.txt")
   head.Open_Device(home + "/googleplay/x86.bin")
   deliver, err := head.Delivery(
      http.Default_Client, "com.google.android.youtube", 1524221376,
   )
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", deliver)
}
