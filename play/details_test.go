package play

import (
   "fmt"
   "os"
   "testing"
   "time"
)

func TestDetails(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   home += "/google-play"
   var token GoogleToken
   token.Data, err = os.ReadFile(home + "/token.txt")
   if err != nil {
      t.Fatal(err)
   }
   if err := token.Unmarshal(); err != nil {
      t.Fatal(err)
   }
   var auth GoogleAuth
   if err := auth.Auth(token); err != nil {
      t.Fatal(err)
   }
   for _, app := range apps {
      name := fmt.Sprint(home, "/", app.platform, ".bin")
      var checkin GoogleCheckin
      checkin.Data, err = os.ReadFile(name)
      if err != nil {
         t.Fatal(err)
      }
      if err := checkin.Unmarshal(); err != nil {
         t.Fatal(err)
      }
      detail, err := checkin.Details(auth, app.id, false)
      if err != nil {
         t.Fatal(err)
      }
      if _, ok := detail.field_13_1_70(); !ok {
         t.Fatal(err)
      }
      if _, ok := detail.field_5(); !ok {
         t.Fatal("Details.Name")
      }
      if _, ok := detail.field_6(); !ok {
         t.Fatal("Details.OfferedBy")
      }
      if _, ok := detail.field_8_1(); !ok {
         t.Fatal("Details.Price")
      }
      if _, ok := detail.field_8_2(); !ok {
         t.Fatal("Details.PriceCurrency")
      }
      if _, ok := detail.field_13_1_82_1_1(); !ok {
         t.Fatal("Details.Requires")
      }
      if _, ok := detail.field_13_1_9(); !ok {
         t.Fatal("Details.Size")
      }
      if _, ok := detail.field_13_1_3(); !ok {
         t.Fatal("Details.VersionCode")
      }
      if _, ok := detail.field_13_1_4(); !ok {
         t.Fatal("Details.VersionName")
      }
      app.date = func() string {
         u, ok := detail.field_13_1_16()
         if !ok {
            t.Fatal("Details.UpdatedOn")
         }
         p, err := time.Parse("Jan 2, 2006", u)
         if err != nil {
            t.Fatal(err)
         }
         return p.Format("2006-01-02")
      }()
      if _, ok := <-detail.field_13_1_17(); !ok {
         t.Fatal("Details.File")
      }
      fmt.Printf("%#v,\n", app)
      time.Sleep(99 * time.Millisecond)
   }
}
