package acquire

import (
   "154.pages.dev/http/option"
   "testing"
)

func Test_Checkin(t *testing.T) {
   option.No_Location()
   option.Trace()
   check, err := new_checkin()
   if err != nil {
      t.Fatal(err)
   }
   if err := check.upload_device(); err != nil {
      t.Fatal(err)
   }
}
