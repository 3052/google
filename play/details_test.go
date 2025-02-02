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
   data, err := os.ReadFile(home + "/token.txt")
   if err != nil {
      t.Fatal(err)
   }
   var token0 Token
   err = token0.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   auth0, err := token0.Auth()
   if err != nil {
      t.Fatal(err)
   }
   for _, app0 := range apps {
      data, err = os.ReadFile(home + "/" + app0.abi + ".txt")
      if err != nil {
         t.Fatal(err)
      }
      var check Checkin
      err = check.Unmarshal(data)
      if err != nil {
         t.Fatal(err)
      }
      detail, err := auth0.Details(check, app0.id, false)
      if err != nil {
         t.Fatal(err)
      }
      if detail.Downloads() == 0 {
         t.Fatal("downloads")
      }
      if detail.Name() == "" {
         t.Fatal("name")
      }
      {
         message, _ := detail.Message.Get(8)()
         _, ok := message.GetVarint(1)()
         if !ok {
            t.Fatal("field 8 1")
         }
      }
      if detail.field_8_2() == "" {
         t.Fatal("field 8 2")
      }
      if detail.field_13_1_4() == "" {
         t.Fatal("field 13 1 4")
      }
      app0.date = func() string {
         time0, err := time.Parse("Jan 2, 2006", detail.field_13_1_16())
         if err != nil {
            t.Fatal(err)
         }
         return time0.Format("2006-01-02")
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
      fmt.Printf("%#v,\n", app0)
      time.Sleep(99 * time.Millisecond)
   }
}

type app_test struct {
   date string
   abi  string
   id   string
}

var apps = []app_test{
   {date: "2024-08-02", abi: "x86", id: "com.wakdev.wdnfc"},
   {date: "2024-12-06", abi: "armeabi-v7a", id: "com.sygic.aura"},
   {date: "2024-12-10", abi: "x86", id: "com.clearchannel.iheartradio.controller"},
   {date: "2024-12-16", abi: "x86", id: "com.amctve.amcfullepisodes"},
   {date: "2024-12-24", abi: "armeabi-v7a", id: "com.xiaomi.smarthome"},
   {date: "2025-01-10", abi: "x86", id: "app.source.getcontact"},
   {date: "2025-01-13", abi: "armeabi-v7a", id: "com.binance.dev"},
   {date: "2025-01-15", abi: "armeabi-v7a-leanback", id: "com.netflix.ninja"},
   {date: "2025-01-16", abi: "arm64-v8a", id: "com.kakaogames.twodin"},
   {date: "2025-01-16", abi: "armeabi-v7a", id: "com.axis.drawingdesk.v3"},
   {date: "2025-01-17", abi: "x86", id: "com.google.android.apps.walletnfcrel"},
   {date: "2025-01-17", abi: "x86", id: "com.pinterest"},
   {date: "2025-01-17", abi: "x86-leanback", id: "com.roku.web.trc"},
   {date: "2025-01-18", abi: "x86", id: "kr.sira.metal"},
   {date: "2025-01-20", abi: "x86", id: "com.instagram.android"},
   {date: "2025-01-21", abi: "arm64-v8a", id: "com.app.xt"},
   {date: "2025-01-21", abi: "x86", id: "com.busuu.android.enc"},
   {date: "2025-01-21", abi: "x86", id: "org.thoughtcrime.securesms"},
   {date: "2025-01-22", abi: "armeabi-v7a", id: "com.madhead.tos.zh"},
   {date: "2025-01-22", abi: "x86", id: "br.com.rodrigokolb.realdrum"},
   {date: "2025-01-22", abi: "x86", id: "com.cabify.rider"},
   {date: "2025-01-22", abi: "x86", id: "com.google.android.youtube"},
}
