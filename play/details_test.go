package play

import (
   "fmt"
   "os"
   "slices"
   "strconv"
   "testing"
   "time"
)

var apps = []app_type{
{date:"2023-02-01",platform:"1",size:18708178,downloads:1563645747,doc:"com.miui.weather2"},
{date:"2023-02-20",platform:"0",size:37564028,downloads:347934672,doc:"org.videolan.vlc"},
{date:"2023-04-12",platform:"0",size:13591550,downloads:5094063,doc:"com.amctve.amcfullepisodes"},
{date:"2023-05-08",platform:"0",size:72540261,downloads:143756560,doc:"br.com.rodrigokolb.realdrum"},
{date:"2023-06-23",platform:"1",size:96751752,downloads:90246658,doc:"com.sygic.aura"},
{date:"2023-06-30",platform:"1",size:620149250,downloads:331766132,doc:"com.supercell.brawlstars"},
{date:"2023-07-28",platform:"0",size:5785991,downloads:22453676,doc:"kr.sira.metal"},
{date:"2023-08-04",platform:"1",size:73687367,downloads:536391,doc:"com.app.xt"},
{date:"2023-08-04",platform:"2",size:870402375,downloads:84214957,doc:"com.miHoYo.GenshinImpact"},
{date:"2023-08-08",platform:"0",size:9035872,downloads:691956417,doc:"com.google.android.apps.walletnfcrel"},
{date:"2023-08-09",platform:"1",size:95833969,downloads:15433116,doc:"com.madhead.tos.zh"},
{date:"2023-08-10",platform:"0",size:82128315,downloads:90235324,doc:"com.clearchannel.iheartradio.controller"},
{date:"2023-08-10",platform:"2",size:147816312,downloads:297401,doc:"com.kakaogames.twodin"},
{date:"2023-08-15",platform:"0",size:157677501,downloads:195745539,doc:"app.source.getcontact"},
{date:"2023-08-18",platform:"1",size:98435268,downloads:49600879,doc:"com.xiaomi.smarthome"},
{date:"2023-08-21",platform:"0",size:88255721,downloads:14592454008,doc:"com.google.android.youtube"},
{date:"2023-08-21",platform:"0",size:86931226,downloads:34567828,doc:"com.cabify.rider"},
{date:"2023-08-21",platform:"0",size:70172136,downloads:5008322663,doc:"com.instagram.android"},
{date:"2023-08-21",platform:"1",size:146603067,downloads:88030401,doc:"com.binance.dev"},
{date:"2023-08-22",platform:"0",size:61406761,downloads:127162129,doc:"org.thoughtcrime.securesms"},
{date:"2023-08-22",platform:"0",size:32784749,downloads:921326697,doc:"com.pinterest"},
{date:"2023-08-24",platform:"1",size:110846862,downloads:17990819,doc:"com.axis.drawingdesk.v3"},
}

type app_type struct {
   date string
   platform string // X-DFE-Device-ID
   size uint64
   downloads uint64
   doc string
}

func (a app_type) GoString() string {
   var b []byte
   b = append(b, "{date:"...)
   b = strconv.AppendQuote(b, a.date)
   b = append(b, ",platform:"...)
   b = append(b, a.platform...)
   b = append(b, ",size:"...)
   b = strconv.AppendUint(b, a.size, 10)
   b = append(b, ",downloads:"...)
   b = strconv.AppendUint(b, a.downloads, 10)
   b = append(b, ",doc:"...)
   b = strconv.AppendQuote(b, a.doc)
   b = append(b, '}')
   return string(b)
}

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
         b, err := os.ReadFile(home + Native_Platforms[app.platform] + ".bin")
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
