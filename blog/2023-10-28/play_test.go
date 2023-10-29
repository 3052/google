package play

import (
   "os"
   "testing"
)

func Test_Details(t *testing.T) {
   var d Details
   var err error
   d.Checkin.Raw, err = os.ReadFile("checkin.txt")
   if err != nil {
      t.Fatal(err)
   }
   d.Checkin.Unmarshal()
   var r Refresh_Token
   r.Raw, err = os.ReadFile("token.txt")
   if err != nil {
      t.Fatal(err)
   }
   r.Unmarshal()
   d.Access_Token.Refresh(r)
   if err := d.Details("hello", false); err != nil {
      t.Fatal(err)
   }
}

func Test_Exchange(t *testing.T) {
   r, err := Exchange("hello")
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("token.txt", r.Raw, 0666)
}

func Test_Sync(t *testing.T) {
   c, err := Phone.Checkin()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("checkin.txt", c.Raw, 0666)
   c.Unmarshal()
   Phone.Sync(c)
}
