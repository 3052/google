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
   data, err := os.ReadFile(home + "/token.txt")
   if err != nil {
      t.Fatal(err)
   }
   var token GoogleToken
   err = token.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   auth, err := token.Auth()
   if err != nil {
      t.Fatal(err)
   }
   data, err = os.ReadFile(home + "/x86.bin")
   if err != nil {
      t.Fatal(err)
   }
   var checkin GoogleCheckin
   err = checkin.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   app := StoreApp{"com.google.android.youtube", 1524221376}
   deliver, err := auth.Delivery(checkin, app, false)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%#v\n", deliver.Message)
}
