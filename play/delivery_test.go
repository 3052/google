package play

import (
   "fmt"
   "os"
   "testing"
)

func TestDelivery(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   home += "/google/play"
   var token GoogleToken
   token.Raw, err = os.ReadFile(home + "/token.txt")
   if err != nil {
      t.Fatal(err)
   }
   err = token.Unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   auth, err := token.Auth()
   if err != nil {
      t.Fatal(err)
   }
   var checkin GoogleCheckin
   checkin.Raw, err = os.ReadFile(home + "/x86.txt")
   if err != nil {
      t.Fatal(err)
   }
   err = checkin.Unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   app := StoreApp{"com.google.android.youtube", 1524221376}
   deliver, err := auth.Delivery(&checkin, &app, false)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%#v\n", deliver.Message)
}
