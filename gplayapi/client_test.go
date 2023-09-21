package gplayapi

import (
   "154.pages.dev/http/option"
   "encoding/json"
   "net/url"
   "os"
   "strconv"
   "testing"
)

const query = "doc=com.zzkko&vc=818"

func Test_Client(t *testing.T) {
   client, err := NewClient("srpen6@gmail.com", "")
   if err != nil {
      t.Fatal(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   enc.Encode(client)
}

func Test_Acquire(t *testing.T) {
   text, err := os.ReadFile("client.json")
   if err != nil {
      t.Fatal(err)
   }
   var client GooglePlayClient
   if err := json.Unmarshal(text, &client); err != nil {
      t.Fatal(err)
   }
   doc, version, err := parse(query)
   if err != nil {
      t.Fatal(err)
   }
   option.No_Location()
   option.Trace()
   if err := client.Acquire(doc, uint64(version)); err != nil {
      t.Fatal(err)
   }
}

func parse(s string) (string, int, error) {
   q, err := url.ParseQuery(s)
   if err != nil {
      return "", 0, err
   }
   v, err := strconv.Atoi(q.Get("vc"))
   if err != nil {
      return "", 0, err
   }
   return q.Get("doc"), v, nil
}

func Test_Delivery(t *testing.T) {
   text, err := os.ReadFile("client.json")
   if err != nil {
      t.Fatal(err)
   }
   var client GooglePlayClient
   if err := json.Unmarshal(text, &client); err != nil {
      t.Fatal(err)
   }
   doc, version, err := parse(query)
   if err != nil {
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
