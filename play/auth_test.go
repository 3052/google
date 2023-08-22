package play

import (
   "encoding/json"
   "os"
   "testing"
   "time"
)

func Test_Auth(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   u, err := user_info(home + "/gmail.json")
   if err != nil {
      t.Fatal(err)
   }
   a, err := New_Auth(u["username"], u["password"])
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

func user_info(name string) (map[string]string, error) {
   b, err := os.ReadFile(name)
   if err != nil {
      return nil, err
   }
   var m map[string]string
   if err := json.Unmarshal(b, &m); err != nil {
      return nil, err
   }
   return m, nil
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

