package play

import "strconv"

var apps = []app_type{
{date:"2023-02-01",requires:24,platform:1,doc:"com.miui.weather2"},
{date:"2023-02-20",requires:17,platform:0,doc:"org.videolan.vlc"},
{date:"2023-04-12",requires:21,platform:0,doc:"com.amctve.amcfullepisodes"},
{date:"2023-05-08",requires:23,platform:0,doc:"br.com.rodrigokolb.realdrum"},
{date:"2023-06-23",requires:23,platform:1,doc:"com.sygic.aura"},
{date:"2023-06-30",requires:24,platform:1,doc:"com.supercell.brawlstars"},
{date:"2023-07-28",requires:21,platform:0,doc:"kr.sira.metal"},
{date:"2023-08-04",requires:21,platform:2,doc:"com.miHoYo.GenshinImpact"},
{date:"2023-08-04",requires:23,platform:1,doc:"com.app.xt"},
{date:"2023-08-08",requires:24,platform:0,doc:"com.google.android.apps.walletnfcrel"},
{date:"2023-08-09",requires:23,platform:1,doc:"com.madhead.tos.zh"},
{date:"2023-08-10",requires:21,platform:2,doc:"com.kakaogames.twodin"},
{date:"2023-08-10",requires:23,platform:0,doc:"com.clearchannel.iheartradio.controller"},
{date:"2023-08-15",requires:26,platform:0,doc:"app.source.getcontact"},
{date:"2023-08-18",requires:21,platform:1,doc:"com.xiaomi.smarthome"},
{date:"2023-08-21",requires:21,platform:0,doc:"com.cabify.rider"},
{date:"2023-08-21",requires:21,platform:1,doc:"com.binance.dev"},
{date:"2023-08-21",requires:24,platform:0,doc:"com.instagram.android"},
{date:"2023-08-21",requires:26,platform:0,doc:"com.google.android.youtube"},
{date:"2023-08-22",requires:21,platform:0,doc:"org.thoughtcrime.securesms"},
{date:"2023-08-22",requires:24,platform:0,doc:"com.pinterest"},
{date:"2023-08-24",requires:23,platform:1,doc:"com.axis.drawingdesk.v3"},
}

func (a app_type) GoString() string {
   var b []byte
   b = append(b, "{date:"...)
   b = strconv.AppendQuote(b, a.date)
   b = append(b, ",requires:"...)
   b = strconv.AppendInt(b, a.requires, 10)
   b = append(b, ",platform:"...)
   b = strconv.AppendInt(b, a.platform, 10)
   b = append(b, ",doc:"...)
   b = strconv.AppendQuote(b, a.doc)
   b = append(b, '}')
   return string(b)
}

type app_type struct {
   date string
   requires int64
   platform int64 // X-DFE-Device-ID
   doc string
}
