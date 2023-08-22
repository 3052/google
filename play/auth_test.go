package play

import (
   "os"
   "testing"
   "time"
)

const code = "oauth2_4/0Adeu5BVtOS_6vKpJRrBcF7xoa5V-J8XfKlMG3J1JbIj5bcaEb5IOX..."

func Test_Auth(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   a, err := New_Auth(code)
   if err != nil {
      t.Fatal(err)
   }
   {
      b, err := a.MarshalText()
      if err != nil {
         t.Fatal(err)
      }
      os.WriteFile(home + "/google/play/auth.txt", b, 0666)
   }
}

func Test_Header(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   var head Header
   head.Auth = make(Auth)
   {
      b, err := os.ReadFile(home + "/google/play/auth.txt")
      if err != nil {
         t.Fatal(err)
      }
      head.Auth.UnmarshalText(b)
   }
   for i := 0; i < 9; i++ {
      if head.Auth.Auth() == "" {
         t.Fatalf("%+v", head)
      }
      time.Sleep(time.Second)
   }
}

