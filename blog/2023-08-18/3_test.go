package play

import (
   "154.pages.dev/http/option"
   "encoding/json"
   "fmt"
   "net/http/httputil"
   "os"
   "testing"
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

func Test_Signin(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   u, err := user_info(home + "/gmail.json")
   if err != nil {
      t.Fatal(err)
   }
   option.No_Location()
   option.Trace()
   setup, err := new_embedded_setup()
   if err != nil {
      t.Fatal(err)
   }
   account, err := setup.account_lookup(u["username"])
   if err != nil {
      t.Fatal(err)
   }
   res, err := account.signin(u["password"])
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   {
      b, err := httputil.DumpResponse(res, true)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(string(b))
   }
}
