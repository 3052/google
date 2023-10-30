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
   home += "/google/play/"
   var token Refresh_Token
   token.Raw, err = os.ReadFile(home + "token.txt")
   if err != nil {
      t.Fatal(err)
   }
   if err := token.Unmarshal(); err != nil {
      t.Fatal(err)
   }
   var d Delivery
   if err := d.Client.Token.Refresh(token); err != nil {
      t.Fatal(err)
   }
   d.Client.Checkin.Raw, err = os.ReadFile(home + "x86.bin")
   if err != nil {
      t.Fatal(err)
   }
   if err := d.Client.Checkin.Unmarshal(); err != nil {
      t.Fatal(err)
   }
   d.App = Application{"com.google.android.youtube", 1524221376}
   if err := d.Delivery(false); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%#v\n", d.m)
}
