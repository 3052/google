package acquire

import (
   "fmt"
   "testing"
)

func Test_Checkin(t *testing.T) {
   check, err := new_checkin()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(check)
}
