package googleplay

import (
   "fmt"
   "os"
   "sort"
   "strconv"
   "testing"
   "time"
)

var apps = []app_type{
   {"2023-06-15",2,295924,"com.kakaogames.twodin"},
   {"2023-06-02",1,524826,"com.app.xt"},
   {"2023-04-12",0,5083467,"com.amctve.amcfullepisodes"},
   {"2023-06-14",1,15386744,"com.madhead.tos.zh"},
   {"2023-06-18",1,17146720,"com.axis.drawingdesk.v3"},
   {"2023-03-23",0,22080563,"kr.sira.metal"},
   {"2023-05-15",1,47920365,"com.xiaomi.smarthome"},
   {"2023-05-12",2,80187885,"com.miHoYo.GenshinImpact"},
   {"2023-06-22",1,83263469,"com.binance.dev"},
   {"2023-05-25",0,89650576,"com.clearchannel.iheartradio.controller"},
   {"2023-06-07",1,89823699,"com.sygic.aura"},
   {"2023-06-15",0,124259171,"org.thoughtcrime.securesms"},
   {"2023-05-08",0,142129205,"br.com.rodrigokolb.realdrum"},
   {"2023-05-10",0,187860132,"app.source.getcontact"},
   {"2023-05-29",1,326475842,"com.supercell.brawlstars"},
   {"2023-02-20",0,341682282,"org.videolan.vlc"},
   {"2023-06-20",0,655061044,"com.google.android.apps.walletnfcrel"},
   {"2023-06-13",0,895249472,"com.pinterest"},
   {"2023-02-01",1,1513248093,"com.miui.weather2"},
   {"2023-06-22",0,4890507365,"com.instagram.android"},
   {"2023-06-16",0,14207300870,"com.google.android.youtube"},
}

func Test_Details(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   home += "/2a/googleplay"
   var head Header
   if err := head.Read_Auth(home + "/auth.txt"); err != nil {
      t.Fatal(err)
   }
   head.Auth.Exchange()
   for _, app := range apps {
      platform := Platforms[app.platform]
      err := head.Read_Device(home + "/" + platform + ".bin")
      if err != nil {
         t.Fatal(err)
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
      if _, err := d.Installation_Size(); err != nil {
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
         app.date = d.String()
      }
      {
         d, err := d.Num_Downloads()
         if err != nil {
            t.Fatal(err)
         }
         app.downloads = d
      }
      time.Sleep(99 * time.Millisecond)
   }
   sort.Slice(apps, func(i, j int) bool {
      return apps[i].downloads < apps[j].downloads
   })
   for _, app := range apps {
      fmt.Print(app, ",\n")
   }
}

type app_type struct {
   date string
   platform int64 // X-DFE-Device-ID
   downloads uint64
   doc string
}

func (a app_type) String() string {
   var b []byte
   b = append(b, '{')
   b = strconv.AppendQuote(b, a.date)
   b = append(b, ',')
   b = strconv.AppendInt(b, a.platform, 10)
   b = append(b, ',')
   b = strconv.AppendUint(b, a.downloads, 10)
   b = append(b, ',')
   b = strconv.AppendQuote(b, a.doc)
   b = append(b, '}')
   return string(b)
}
