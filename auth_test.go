package googleplay

import (
   "encoding/json"
   "os"
   "testing"
   "time"
)

func credential(name string) (map[string]string, error) {
   data, err := os.ReadFile(name)
   if err != nil {
      return nil, err
   }
   var cred map[string]string
   if err := json.Unmarshal(data, &cred); err != nil {
      return nil, err
   }
   return cred, nil
}

func Test_Auth(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   cred, err := credential(home + "/Document/gmail.json")
   if err != nil {
      t.Fatal(err)
   }
   res, err := New_Auth(cred["username"], cred["password"])
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   if err := res.Create(home + "/Documents/googleplay.txt"); err != nil {
      t.Fatal(err)
   }
}

func Test_Header(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   var head Header
   head.Open_Auth(home + "/Documents/googleplay.txt")
   for i := 0; i < 9; i++ {
      if head.Auth.Get_Auth() == "" {
         t.Fatalf("%+v", head)
      }
      time.Sleep(time.Second)
   }
}
