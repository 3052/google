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
   home += "/google-play/"
   var token RefreshToken
   token.Raw, err = os.ReadFile(home + "token.txt")
   if err != nil {
      t.Fatal(err)
   }
   if err := token.Unmarshal(); err != nil {
      t.Fatal(err)
   }
   var d Details
   if err := d.Token.Refresh(token); err != nil {
      t.Fatal(err)
   }
   for _, app := range apps {
      name := fmt.Sprint(home, app.platform, ".bin")
      d.Checkin.Raw, err = os.ReadFile(name)
      if err != nil {
         t.Fatal(err)
      }
      if err := d.Checkin.Unmarshal(); err != nil {
         t.Fatal(err)
      }
      if err := d.Details(app.id, false); err != nil {
         t.Fatal(err)
      }
      if _, ok := d.Downloads(); !ok {
         t.Fatal(err)
      }
      if _, ok := d.Name(); !ok {
         t.Fatal("Details.Name")
      }
      if _, ok := d.OfferedBy(); !ok {
         t.Fatal("Details.OfferedBy")
      }
      if _, ok := d.Price(); !ok {
         t.Fatal("Details.Price")
      }
      if _, ok := d.PriceCurrency(); !ok {
         t.Fatal("Details.PriceCurrency")
      }
      if _, ok := d.Requires(); !ok {
         t.Fatal("Details.Requires")
      }
      if _, ok := d.Size(); !ok {
         t.Fatal("Details.Size")
      }
      if _, ok := d.VersionCode(); !ok {
         t.Fatal("Details.VersionCode")
      }
      if _, ok := d.VersionName(); !ok {
         t.Fatal("Details.VersionName")
      }
      app.date = func() string {
         u, ok := d.UpdatedOn()
         if !ok {
            t.Fatal("Details.UpdatedOn")
         }
         p, err := time.Parse("Jan 2, 2006", u)
         if err != nil {
            t.Fatal(err)
         }
         return p.Format("2006-01-02")
      }()
      var ok bool
      d.Files(func(uint64) {
         ok = true
      })
      if !ok {
         t.Fatal("Details.Files")
      }
      fmt.Printf("%#v,\n", app)
      time.Sleep(99 * time.Millisecond)
   }
}
