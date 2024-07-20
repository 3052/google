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
   text, err := os.ReadFile(home + "/token.txt")
   if err != nil {
      t.Fatal(err)
   }
   var token GoogleToken
   err = token.Unmarshal(text)
   if err != nil {
      t.Fatal(err)
   }
   auth, err := token.Auth()
   if err != nil {
      t.Fatal(err)
   }
   for _, app := range apps {
      data, err := os.ReadFile(fmt.Sprint(home, "/", Abi[app.abi], ".bin"))
      if err != nil {
         t.Fatal(err)
      }
      var checkin GoogleCheckin
      err = checkin.Unmarshal(data)
      if err != nil {
         t.Fatal(err)
      }
      detail, err := auth.Details(checkin, app.id, false)
      if err != nil {
         t.Fatal(err)
      }
      if _, ok := detail.Downloads(); !ok {
         t.Fatal("downloads")
      }
      if _, ok := detail.Name(); !ok {
         t.Fatal("name")
      }
      if _, ok := detail.field_6(); !ok {
         t.Fatal("field 6")
      }
      if _, ok := detail.field_8_1(); !ok {
         t.Fatal("field 8 1")
      }
      if _, ok := detail.field_8_2(); !ok {
         t.Fatal("field 8 2")
      }
      if _, ok := detail.field_13_1_4(); !ok {
         t.Fatal("field 13 1 4")
      }
      app.date = func() string {
         u, ok := detail.field_13_1_16()
         if !ok {
            t.Fatal("field 13 1 16")
         }
         p, err := time.Parse("Jan 2, 2006", u)
         if err != nil {
            t.Fatal(err)
         }
         return p.Format("2006-01-02")
      }()
      if _, ok := <-detail.field_13_1_17(); !ok {
         t.Fatal("field 13 1 17")
      }
      if _, ok := detail.field_13_1_82_1_1(); !ok {
         t.Fatal("field 13 1 82 1 1")
      }
      if _, ok := detail.size(); !ok {
         t.Fatal("size")
      }
      if _, ok := detail.version_code(); !ok {
         t.Fatal("version code")
      }
      fmt.Printf("%#v,\n", app)
      time.Sleep(99 * time.Millisecond)
   }
}
