package aurora

import (
   "154.pages.dev/http/option"
   "fmt"
   "testing"
)

func Test_Aurora(t *testing.T) {
   option.No_Location()
   option.Verbose()
   aurora, err := new_aurora_OSS()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", aurora)
}
