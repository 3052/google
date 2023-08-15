package play

import (
   "154.pages.dev/http/option"
   "net/http/httputil"
   "os"
   "testing"
)

func Test_Account_Lookup(t *testing.T) {
   option.No_Location()
   option.Trace()
   setup, err := new_embedded_setup()
   if err != nil {
      t.Fatal(err)
   }
   cookies, err := setup.eligible()
   if err != nil {
      t.Fatal(err)
   }
   res, err := setup.account_lookup(cookies)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   text, err := httputil.DumpResponse(res, false)
   if err != nil {
      t.Fatal(err)
   }
   os.Stdout.Write(text)
}
