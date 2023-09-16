package aurora

import (
   "154.pages.dev/http/option"
   "fmt"
   "os"
   "testing"
)

func Test_Marshal(t *testing.T) {
   option.No_Location()
   option.Verbose()
   var aurora aurora_OSS
   text, err := aurora.Marshal()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("aurora.json", text, 0666)
}

func Test_Unmarshal(t *testing.T) {
   var aurora aurora_OSS
   text, err := os.ReadFile("aurora.json")
   if err != nil {
      t.Fatal(err)
   }
   if err := aurora.Unmarshal(text); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", aurora)
}
