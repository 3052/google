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
   text, err := new(aurora_OSS).Marshal()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("aurora.json", text, 0666)
}

func Test_Unmarshal(t *testing.T) {
   text, err := os.ReadFile("aurora.json")
   if err != nil {
      t.Fatal(err)
   }
   var aurora aurora_OSS
   if err := aurora.Unmarshal(text); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", aurora)
}
