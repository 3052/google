package play

import (
   "154.pages.dev/http/option"
   "fmt"
   "net/http/httputil"
   "testing"
)

func Test_Signin(t *testing.T) {
   option.No_Location()
   option.Trace()
   setup, err := new_embedded_setup()
   if err != nil {
      t.Fatal(err)
   }
   account, err := setup.account_lookup()
   if err != nil {
      t.Fatal(err)
   }
   res, err := account.signin()
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
