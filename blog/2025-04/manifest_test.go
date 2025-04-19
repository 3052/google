package manifest

import (
   "encoding/xml"
   "fmt"
   "os"
   "testing"
)

func Test(t *testing.T) {
   data, err := os.ReadFile("AndroidManifest.xml")
   if err != nil {
      t.Fatal(err)
   }
   var manifest1 manifest
   err = xml.Unmarshal(data, &manifest1)
   if err != nil {
      t.Fatal(err)
   }
   for intent := range manifest1.intent_filter() {
      fmt.Print(&intent, "\n\n")
   }
}
