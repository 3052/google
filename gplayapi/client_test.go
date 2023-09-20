package gplayapi

import (
   "encoding/json"
   "os"
   "testing"
)

const (
   doc = "com.google.android.youtube"
   version = 1537856960
)

func Test_Client(t *testing.T) {
   text, err := os.ReadFile("client.json")
   if err != nil {
      t.Fatal(err)
   }
   var client GooglePlayClient
   if err := json.Unmarshal(text, &client); err != nil {
      t.Fatal(err)
   }
   delivery, err := client.GetDeliveryResponse(doc, version)
   if err != nil {
      t.Fatal(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   enc.Encode(delivery)
}
