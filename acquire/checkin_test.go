package acquire

import (
   "fmt"
   "testing"
)

func Test_Checkin(t *testing.T) {
   id, err := checkin()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", id)
}
