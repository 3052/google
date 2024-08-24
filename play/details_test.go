package play

import (
   "fmt"
   "os"
   "path/filepath"
   "testing"
   "time"
)

func TestDetails(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   home = filepath.ToSlash(home) + "/google-play"
   var token GoogleToken
   token.Raw, err = os.ReadFile(home + "/token.txt")
   if err != nil {
      t.Fatal(err)
   }
   err = token.Unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   auth, err := token.Auth()
   if err != nil {
      t.Fatal(err)
   }
   var max_app struct {
      downloads uint64
      id string
   }
   for _, app := range apps {
      var checkin GoogleCheckin
      checkin.Raw, err = os.ReadFile(home + "/" + app.abi + ".txt")
      if err != nil {
         t.Fatal(err)
      }
      err = checkin.Unmarshal()
      if err != nil {
         t.Fatal(err)
      }
      detail, err := auth.Details(&checkin, app.id, false)
      if err != nil {
         t.Fatal(err)
      }
      if v, ok := detail.Downloads(); ok {
         if v > max_app.downloads {
            max_app.downloads = v
            max_app.id = app.id
         }
      } else {
         t.Fatal("downloads")
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
      if _, ok := detail.field_13_1_12(); !ok {
         t.Fatal("field 13 1 12")
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
      if _, ok := detail.field_13_1_17()(); !ok {
         t.Fatal("field 13 1 17")
      }
      if _, ok := detail.field_13_1_82_1_1(); !ok {
         t.Fatal("field 13 1 82 1 1")
      }
      if _, ok := detail.Name(); !ok {
         t.Fatal("name")
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
   fmt.Println("max(downloads) = ", max_app.id)
}
