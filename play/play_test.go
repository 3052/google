package play

import "fmt"

var apps = []app_type{
{date:"2023-10-31",requires:21,abi:0,id:"com.roku.web.trc"},
{date:"2024-01-08",requires:21,abi:0,id:"kr.sira.metal"},
{date:"2024-01-29",requires:21,abi:0,id:"com.amctve.amcfullepisodes"},
{date:"2024-03-25",requires:21,abi:1,id:"com.xiaomi.smarthome"},
{date:"2024-04-03",requires:23,abi:0,id:"com.clearchannel.iheartradio.controller"},
{date:"2024-04-11",requires:23,abi:1,id:"com.sygic.aura"},
{date:"2024-04-16",requires:24,abi:0,id:"com.google.android.apps.walletnfcrel"},
{date:"2024-04-18",requires:21,abi:2,id:"com.kakaogames.twodin"},
{date:"2024-04-19",requires:26,abi:0,id:"app.source.getcontact"},
{date:"2024-04-22",requires:23,abi:1,id:"com.madhead.tos.zh"},
{date:"2024-04-23",requires:21,abi:0,id:"org.thoughtcrime.securesms"},
{date:"2024-04-23",requires:28,abi:0,id:"com.busuu.android.enc"},
{date:"2024-04-24",requires:23,abi:2,id:"com.app.xt"},
{date:"2024-04-26",requires:26,abi:0,id:"com.google.android.youtube"},
{date:"2024-04-29",requires:21,abi:0,id:"com.cabify.rider"},
{date:"2024-04-29",requires:21,abi:1,id:"com.binance.dev"},
{date:"2024-04-29",requires:24,abi:0,id:"com.instagram.android"},
{date:"2024-04-29",requires:24,abi:0,id:"com.pinterest"},
{date:"2024-05-01",requires:23,abi:0,id:"br.com.rodrigokolb.realdrum"},
{date:"2024-05-03",requires:23,abi:1,id:"com.axis.drawingdesk.v3"},
}

type app_type struct {
   date     string
   requires int
   abi int
   id      string
}

func (a app_type) GoString() string {
   var b []byte
   b = fmt.Appendf(b, "{date:%q", a.date)
   b = fmt.Append(b, ",requires:", a.requires)
   b = fmt.Appendf(b, ",abi:%d", a.abi)
   b = fmt.Appendf(b, ",id:%q", a.id)
   b = append(b, '}')
   return string(b)
}
