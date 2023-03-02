package googleplay

import (
   "os"
   "testing"
)

func Test_Get_Items(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   var head Header
   if err := head.Open_Device(home + "/googleplay/x86.bin"); err != nil {
      t.Fatal(err)
   }
   res, err := head.Get_Items("com.watchfacestudio.md307digital")
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   if err := res.Create("ignore.txt"); err != nil {
      t.Fatal(err)
   }
   item, err := Open_Items("ignore.txt")
   if err != nil {
      t.Fatal(err)
   }
   if v, err := item.Category(); err != nil {
      t.Fatal(err)
   } else if v != "Personalization" {
      t.Fatal(v)
   }
   if v, err := item.Creator(); err != nil {
      t.Fatal(err)
   } else if v != "Matteo Dini MD ®" {
      t.Fatal(v)
   }
   if v, err := item.Offer(); err != nil {
      t.Fatal(err)
   } else if v != "$0.99" {
      t.Fatal(v)
   }
   if v, err := item.Title(); err != nil {
      t.Fatal(err)
   } else if v != "MD307 Digital watch face" {
      t.Fatal(v)
   }
   if v, err := item.Upload_Date(); err != nil {
      t.Fatal(err)
   } else if v != "Nov 3, 2022" {
      t.Fatal(v)
   }
   if v, err := item.Version(); err != nil {
      t.Fatal(err)
   } else if v != "3.0.0" {
      t.Fatal(v)
   }
   if v, err := item.Version_Code(); err != nil {
      t.Fatal(err)
   } else if v != 7 {
      t.Fatal(v)
   }
   if v, err := item.Num_Downloads(); err != nil {
      t.Fatal(err)
   } else if v < 5023 {
      t.Fatal(v)
   }
}
