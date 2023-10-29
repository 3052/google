package play

import (
   "os"
   "testing"
)

func Test_Details(t *testing.T) {
   var detail Details
   var err error
   detail.Checkin.Raw, err = os.ReadFile("checkin.txt")
   if err != nil {
      t.Fatal(err)
   }
   detail.Checkin.Unmarshal()
   var token Refresh_Token
   token.Raw, err = os.ReadFile("token.txt")
   if err != nil {
      t.Fatal(err)
   }
   token.Unmarshal()
   detail.Access_Token.Refresh(token)
   if err := detail.Details("hello", false); err != nil {
      t.Fatal(err)
   }
}

func Test_Exchange(t *testing.T) {
   token, err := Exchange("hello")
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("token.txt", token.Raw, 0666)
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
