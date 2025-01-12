package play

import (
   "fmt"
   "os"
   "path/filepath"
   "testing"
   "time"
)

type app_test struct {
   date     string
   abi      string
   id       string
}

var apps = []app_test{
{date:"2024-06-15", abi:"x86", id:"kr.sira.metal"},
{date:"2024-07-12", abi:"x86", id:"com.amctve.amcfullepisodes"},
{date:"2024-07-31", abi:"x86", id:"br.com.rodrigokolb.realdrum"},
{date:"2024-08-30", abi:"arm64-v8a", id:"com.kakaogames.twodin"},
{date:"2024-09-03", abi:"x86", id:"com.clearchannel.iheartradio.controller"},
{date:"2024-09-10", abi:"x86", id:"app.source.getcontact"},
{date:"2024-09-19", abi:"armeabi-v7a", id:"com.madhead.tos.zh"},
{date:"2024-09-19", abi:"armeabi-v7a", id:"com.xiaomi.smarthome"},
{date:"2024-09-20", abi:"armeabi-v7a", id:"com.sygic.aura"},
{date:"2024-09-23", abi:"x86", id:"com.busuu.android.enc"},
{date:"2024-09-24", abi:"x86", id:"com.google.android.youtube"},
{date:"2024-09-24", abi:"x86", id:"com.pinterest"},
{date:"2024-09-25", abi:"x86", id:"com.google.android.apps.walletnfcrel"},
{date:"2024-09-25", abi:"x86", id:"org.thoughtcrime.securesms"},
{date:"2024-09-26", abi:"arm64-v8a", id:"com.app.xt"},
{date:"2024-09-27", abi:"armeabi-v7a", id:"com.binance.dev"},
{date:"2024-09-27", abi:"x86", id:"com.cabify.rider"},
{date:"2024-09-30", abi:"x86", id:"com.instagram.android"},
{date:"2024-10-01", abi:"armeabi-v7a", id:"com.axis.drawingdesk.v3"},
{date:"2024-11-21", abi:"armeabi-v7a-leanback", id:"com.netflix.ninja"},
{date:"2024-12-26", abi:"x86-leanback", id:"com.roku.web.trc"},
}

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
      value, _ := detail.Message.Get(8)()
      _, ok := value.GetVarint(1)()
      if !ok {
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
