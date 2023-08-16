package play

import (
   "154.pages.dev/http/option"
   "fmt"
   "net/http/httputil"
   "testing"
)

func Test_Password(t *testing.T) {
   option.No_Location()
   option.Trace()
   setup, err := new_embedded_setup()
   if err != nil {
      t.Fatal(err)
   }
   sign, err := setup.signup()
   if err != nil {
      t.Fatal(err)
   }
   account, err := setup.account_lookup(sign)
   if err != nil {
      t.Fatal(err)
   }
   res, err := account.password()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   text, err := httputil.DumpResponse(res, true)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(string(text))
}
