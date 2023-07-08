package google_play

import (
   "encoding/json"
   "os"
   "testing"
   "time"
)

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

func Test_Auth(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   u, err := user_info(home + "/gmail.json")
   if err != nil {
      t.Fatal(err)
   }
   res, err := New_Auth(u["username"], u["password"])
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   if err := res.Write_File(home + "/google-play/auth.txt"); err != nil {
      t.Fatal(err)
   }
}

func Test_Header(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   var head Header
   head.Read_Auth(home + "/google-play/auth.txt")
   for i := 0; i < 9; i++ {
      if head.Auth.Get_Auth() == "" {
         t.Fatalf("%+v", head)
      }
      time.Sleep(time.Second)
   }
}
