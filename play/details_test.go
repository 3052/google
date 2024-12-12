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
   data, err := os.ReadFile(home + "/token.txt")
   if err != nil {
      t.Fatal(err)
   }
   var token GoogleToken
   err = token.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   auth, err := token.Auth()
   if err != nil {
      t.Fatal(err)
   }
   for _, app := range apps {
      data, err = os.ReadFile(home + "/" + app.abi + ".txt")
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
      if detail.Downloads() == 0 {
         t.Fatal("downloads")
      }
      if detail.Name() == "" {
         t.Fatal("name")
      }
      if detail.field_8_1() == 0 {
         t.Fatal("field 8 1")
      }
      if detail.field_8_2() == "" {
         t.Fatal("field 8 2")
      }
      if detail.field_13_1_4() == "" {
         t.Fatal("field 13 1 4")
      }
      app.date = func() string {
         date, err := time.Parse("Jan 2, 2006", detail.field_13_1_16())
         if err != nil {
            t.Fatal(err)
         }
         return date.Format("2006-01-02")
      }()
      if _, ok := detail.field_13_1_17()(); !ok {
         t.Fatal("field 13 1 17")
      }
      if detail.field_13_1_82_1_1() == "" {
         t.Fatal("field 13 1 82 1 1")
      }
      if detail.field_15_18() == "" {
         t.Fatal("field_15_18")
      }
      if detail.size() == 0 {
         t.Fatal("size")
      }
      if detail.version_code() == 0 {
         t.Fatal("version code")
      }
      fmt.Printf("%#v,\n", app)
      time.Sleep(99 * time.Millisecond)
   }
}
