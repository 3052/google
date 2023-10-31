package play

import "strconv"

var apps = []app_type{
{date:"2023-02-01",requires:24,platform:1,id:"com.miui.weather2"},
{date:"2023-02-20",requires:17,platform:0,id:"org.videolan.vlc"},
{date:"2023-04-12",requires:21,platform:0,id:"com.amctve.amcfullepisodes"},
{date:"2023-10-26",requires:23,platform:0,id:"br.com.rodrigokolb.realdrum"},
{date:"2023-08-04",requires:23,platform:2,id:"com.app.xt"},
{date:"2023-10-25",requires:24,platform:1,id:"com.supercell.brawlstars"},
{date:"2023-09-14",requires:23,platform:1,id:"com.sygic.aura"},
{date:"2023-09-16",requires:23,platform:1,id:"com.axis.drawingdesk.v3"},
{date:"2023-10-19",requires:21,platform:2,id:"com.kakaogames.twodin"},
{date:"2023-09-21",requires:26,platform:0,id:"app.source.getcontact"},
{date:"2023-10-24",requires:21,platform:0,id:"com.cabify.rider"},
{date:"2023-10-25",requires:23,platform:1,id:"com.madhead.tos.zh"},
{date:"2023-10-23",requires:23,platform:0,id:"com.clearchannel.iheartradio.controller"},
{date:"2023-10-23",requires:24,platform:0,id:"com.pinterest"},
{date:"2023-10-16",requires:21,platform:0,id:"kr.sira.metal"},
{date:"2023-10-24",requires:24,platform:0,id:"com.instagram.android"},
{date:"2023-10-20",requires:21,platform:0,id:"org.thoughtcrime.securesms"},
{date:"2023-10-17",requires:24,platform:0,id:"com.google.android.apps.walletnfcrel"},
{date:"2023-10-19",requires:26,platform:0,id:"com.google.android.youtube"},
{date:"2023-10-27",requires:21,platform:1,id:"com.binance.dev"},
{date:"2023-10-25",requires:21,platform:1,id:"com.xiaomi.smarthome"},
{date:"2023-10-23",requires:28,platform:0,id:"com.busuu.android.enc"},
}

func (a app_type) GoString() string {
   var b []byte
   b = append(b, "{date:"...)
   b = strconv.AppendQuote(b, a.date)
   b = append(b, ",requires:"...)
   b = strconv.AppendInt(b, a.requires, 10)
   b = append(b, ",platform:"...)
   b = strconv.AppendInt(b, a.platform, 10)
   b = append(b, ",id:"...)
   b = strconv.AppendQuote(b, a.id)
   b = append(b, '}')
   return string(b)
}

type app_type struct {
   date     string
   requires int64
   platform int64 // X-DFE-Device-ID
   id      string
}
