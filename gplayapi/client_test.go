package gplayapi

import (
   "154.pages.dev/http/option"
   "encoding/json"
   "fmt"
   "net/url"
   "os"
   "strconv"
   "testing"
   "time"
)

const query = "doc=com.zzkko&vc=818"

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
   fmt.Println("sleep")
   time.Sleep(9*time.Second)
}

// case 3: return nil, errors.New("acquire")
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
   option.No_Location()
   option.Trace()
   delivery, err := client.GetDeliveryResponse(doc, version)
   if err != nil {
      t.Fatal(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   enc.Encode(delivery)
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

func Test_Client(t *testing.T) {
   token, err := func() (string, error) {
      b, err := os.ReadFile(`C:\Users\Steven\google\play\token.txt`)
      if err != nil {
         return "", err
      }
      return parseResponse(string(b))["Token"], nil
   }()
   if err != nil {
      t.Fatal(err)
   }
   client, err := NewClientWithDeviceInfo("srpen6@gmail.com", token, Pixel3a)
   if err != nil {
      t.Fatal(err)
   }
   file, err := os.Create("client.json")
   if err != nil {
      t.Fatal(err)
   }
   defer file.Close()
   enc := json.NewEncoder(file)
   enc.SetIndent("", " ")
   enc.Encode(client)
}
