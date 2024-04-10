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
   token.Data, err = os.ReadFile(home + "/token.txt")
   if err != nil {
      t.Fatal(err)
   }
   if err := token.Unmarshal(); err != nil {
      t.Fatal(err)
   }
   var auth GoogleAuth
   if err := auth.Auth(token); err != nil {
      t.Fatal(err)
   }
   var checkin GoogleCheckin
   checkin.Data, err = os.ReadFile(home + "/x86.bin")
   if err != nil {
      t.Fatal(err)
   }
   if err := checkin.Unmarshal(); err != nil {
      t.Fatal(err)
   }
   app := AndroidApp{"com.google.android.youtube", 1524221376}
   var deliver Delivery
   if err := deliver.Delivery(auth, checkin, app, false); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%#v\n", d.m)
}
