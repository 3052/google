package play

import "fmt"

var apps = []app_type{
{date:"2024-01-08",requires:21,platform:0,id:"kr.sira.metal"},
{date:"2024-01-29",requires:21,platform:0,id:"com.amctve.amcfullepisodes"},
{date:"2024-01-30",requires:21,platform:1,id:"com.xiaomi.smarthome"},
{date:"2024-02-08",requires:23,platform:0,id:"br.com.rodrigokolb.realdrum"},
{date:"2024-02-08",requires:26,platform:0,id:"app.source.getcontact"},
{date:"2024-02-20",requires:24,platform:0,id:"com.google.android.apps.walletnfcrel"},
{date:"2024-02-22",requires:21,platform:2,id:"com.kakaogames.twodin"},
{date:"2024-02-22",requires:23,platform:1,id:"com.sygic.aura"},
{date:"2024-02-23",requires:21,platform:0,id:"org.thoughtcrime.securesms"},
{date:"2024-02-28",requires:23,platform:0,id:"com.clearchannel.iheartradio.controller"},
{date:"2024-02-28",requires:23,platform:1,id:"com.axis.drawingdesk.v3"},
{date:"2024-03-01",requires:21,platform:1,id:"com.binance.dev"},
{date:"2024-03-01",requires:23,platform:1,id:"com.madhead.tos.zh"},
{date:"2024-03-04",requires:24,platform:0,id:"com.instagram.android"},
{date:"2024-03-04",requires:26,platform:0,id:"com.google.android.youtube"},
{date:"2024-03-05",requires:24,platform:0,id:"com.pinterest"},
{date:"2024-03-06",requires:23,platform:2,id:"com.app.xt"},
{date:"2024-03-06",requires:28,platform:0,id:"com.busuu.android.enc"},
{date:"2024-03-10",requires:21,platform:0,id:"com.cabify.rider"},
}

type app_type struct {
   date     string
   requires int
   platform Platform
   id      string
}

func (a app_type) GoString() string {
   var b []byte
   b = fmt.Appendf(b, "{date:%q", a.date)
   b = fmt.Append(b, ",requires:", a.requires)
   b = fmt.Appendf(b, ",platform:%d", a.platform)
   b = fmt.Appendf(b, ",id:%q", a.id)
   b = append(b, '}')
   return string(b)
}
