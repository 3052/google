package play

import (
   "154.pages.dev/http/option"
   "testing"
   "time"
)

func Test_Toc(t *testing.T) {
   option.No_Location()
   option.Verbose()
   check, err := new_checkin()
   if err != nil {
      t.Fatal(err)
   }
   time.Sleep(9*time.Second)
   if err := check.upload(); err != nil {
      t.Fatal(err)
   }
   option.Trace()
   time.Sleep(9*time.Second)
   auth, err := check.auth()
   if err != nil {
      t.Fatal(err)
   }
   time.Sleep(9*time.Second)
   if err := check.toc(auth); err != nil {
      t.Fatal(err)
   }
}
