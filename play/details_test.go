package play

import (
   "fmt"
   "os"
   "sort"
   "strconv"
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
   if err := head.Read_Auth(home + "auth.txt"); err != nil {
      t.Fatal(err)
   }
   head.Auth.Exchange()
   for i, app := range apps {
      err := head.Read_Device(home + Platforms[app.platform] + ".bin")
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
         apps[i].date = d.String()
      }
      {
         d, err := d.Num_Downloads()
         if err != nil {
            t.Fatal(err)
         }
         apps[i].downloads = d
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
var apps = []app_type{
   {"2023-02-01",1,1519802840,"com.miui.weather2"},
   {"2023-02-20",0,342652288,"org.videolan.vlc"},
   {"2023-03-23",0,22129580,"kr.sira.metal"},
   {"2023-04-12",0,5085578,"com.amctve.amcfullepisodes"},
   {"2023-05-08",0,142364173,"br.com.rodrigokolb.realdrum"},
   {"2023-05-10",0,189095178,"app.source.getcontact"},
   {"2023-05-12",2,80820762,"com.miHoYo.GenshinImpact"},
   {"2023-05-25",0,89735370,"com.clearchannel.iheartradio.controller"},
   {"2023-06-15",1,48170136,"com.xiaomi.smarthome"},
   {"2023-06-15",2,296264,"com.kakaogames.twodin"},
   {"2023-06-20",0,659843093,"com.google.android.apps.walletnfcrel"},
   {"2023-06-21",1,15395664,"com.madhead.tos.zh"},
   {"2023-06-22",1,83822542,"com.binance.dev"},
   {"2023-06-23",1,89867479,"com.sygic.aura"},
   {"2023-06-26",0,124636422,"org.thoughtcrime.securesms"},
   {"2023-06-26",0,14257009198,"com.google.android.youtube"},
   {"2023-06-26",0,33784619,"com.cabify.rider"},
   {"2023-06-26",0,4906291283,"com.instagram.android"},
   {"2023-06-27",0,898584734,"com.pinterest"},
   {"2023-06-29",1,526676,"com.app.xt"},
   {"2023-06-30",1,327293582,"com.supercell.brawlstars"},
   {"2023-07-01",1,17597951,"com.axis.drawingdesk.v3"},
}

