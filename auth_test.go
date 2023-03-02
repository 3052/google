package googleplay

import (
   "os"
   "testing"
   "time"
)

func Test_Auth(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   res, err := New_Auth(email, password)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   if err := res.Create(home + "/googleplay/auth.txt"); err != nil {
      t.Fatal(err)
   }
}

func Test_Header(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   var head Header
   head.Open_Auth(home + "/googleplay/auth.txt")
   for i := 0; i < 9; i++ {
      if head.Auth.Get_Auth() == "" {
         t.Fatalf("%+v", head)
      }
      time.Sleep(time.Second)
   }
}
