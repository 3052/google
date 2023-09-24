package play

import (
   "fmt"
   "os"
   "slices"
   "testing"
   "time"
)

func Test_Details(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   home += "/google/play/"
   var head Header
   {
      b, err := os.ReadFile(home + "token.txt")
      if err != nil {
         t.Fatal(err)
      }
      head.Set_Authorization(b)
   }
   head.Set_Agent(false)
   for i, app := range apps {
      {
         b, err := os.ReadFile(home + Platforms[app.platform] + ".bin")
         if err != nil {
            t.Fatal(err)
         }
         head.Set_Device(b)
      }
      d, err := head.Details(app.doc)
      if err != nil {
         t.Fatal(err)
      }
      if _, err := d.Version(); err != nil {
         t.Fatal(err)
      }
      if _, err := d.Version_Code(); err != nil {
         t.Fatal(err)
      }
      if _, err := d.Title(); err != nil {
         t.Fatal(err)
      }
      apps[i].size, err = d.Installation_Size()
      if err != nil {
         t.Fatal(err)
      }
      if _, err := d.Currency_Code(); err != nil {
         t.Fatal(err)
      }
      {
         s, err := d.Upload_Date()
         if err != nil {
            t.Fatal(err)
         }
         d, err := time.Parse("Jan 2, 2006", s)
         if err != nil {
            t.Fatal(err)
         }
         apps[i].date = d.Format("2006-01-02")
      }
      {
         d, err := d.Num_Downloads()
         if err != nil {
            t.Fatal(err)
         }
         apps[i].downloads = d
      }
      fmt.Printf("%#v,\n", app)
      time.Sleep(99 * time.Millisecond)
   }
   min_size := slices.MinFunc(apps, func(a, b app_type) int {
      return int(a.size - b.size)
   })
   fmt.Println("min size", min_size)
   max_downloads := slices.MaxFunc(apps, func(a, b app_type) int {
      return int(a.downloads - b.downloads)
   })
   fmt.Println("max downloads", max_downloads)
}
