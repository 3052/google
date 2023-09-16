package aurora

import (
   "154.pages.dev/http/option"
   "fmt"
   "testing"
)

func Test_Aurora(t *testing.T) {
   option.No_Location()
   option.Verbose()
   var aurora aurora_OSS
   text, err := aurora.MarshalText()
   if err != nil {
      t.Fatal(err)
   }
   if err := aurora.UnmarshalText(text); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", aurora)
}
