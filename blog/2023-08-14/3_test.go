package play

import (
   "154.pages.dev/http/option"
   "fmt"
   "testing"
)

func Test_Eligible(t *testing.T) {
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
   fmt.Println(cookies)
}
