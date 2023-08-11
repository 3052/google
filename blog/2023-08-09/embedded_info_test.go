package play

import (
   "154.pages.dev/http/option"
   "fmt"
   "testing"
)

func Test_Info(t *testing.T) {
   option.No_Location()
   option.Verbose()
   s, _, err := get_embedded_info()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(s)
}
