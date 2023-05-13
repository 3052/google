package googleplay

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "os"
   "strings"
   "testing"
   "time"
)

func sign_in(name string) ([]string, error) {
   data, err := os.ReadFile(name)
   if err != nil {
      return nil, err
   }
   return strings.Split(string(data), "\n"), nil
}

func Test_Auth(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   account, err := sign_in(home + "/Documents/gmail.txt")
   if err != nil {
      t.Fatal(err)
   }
   http.Default_Client.Log_Level = 2
   res, err := New_Auth(account[0], account[1])
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   if err := res.Write_File(home + "/Documents/googleplay.txt"); err != nil {
      t.Fatal(err)
   }
}

func Test_Header(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   var head Header
   head.Read_Auth(home + "/Documents/googleplay.txt")
   for i := 0; i < 9; i++ {
      if head.Auth.Get_Auth() == "" {
         t.Fatalf("%+v", head)
      }
      time.Sleep(time.Second)
   }
}
