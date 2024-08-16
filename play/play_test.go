package play

var apps = []app_type{
{date:"2023-10-31", abi:"x86-leanback", id:"com.roku.web.trc"},
{date:"2024-06-15", abi:"x86", id:"kr.sira.metal"},
{date:"2024-07-12", abi:"x86", id:"com.amctve.amcfullepisodes"},
{date:"2024-07-22", abi:"armeabi-v7a", id:"com.sygic.aura"},
{date:"2024-07-23", abi:"armeabi-v7a", id:"com.axis.drawingdesk.v3"},
{date:"2024-07-31", abi:"x86", id:"br.com.rodrigokolb.realdrum"},
{date:"2024-08-04", abi:"armeabi-v7a", id:"com.xiaomi.smarthome"},
{date:"2024-08-06", abi:"x86", id:"com.cabify.rider"},
{date:"2024-08-07", abi:"x86", id:"app.source.getcontact"},
{date:"2024-08-08", abi:"arm64-v8a", id:"com.kakaogames.twodin"},
{date:"2024-08-08", abi:"x86", id:"com.busuu.android.enc"},
{date:"2024-08-12", abi:"armeabi-v7a", id:"com.binance.dev"},
{date:"2024-08-12", abi:"x86", id:"com.google.android.apps.walletnfcrel"},
{date:"2024-08-13", abi:"arm64-v8a", id:"com.app.xt"},
{date:"2024-08-13", abi:"armeabi-v7a", id:"com.madhead.tos.zh"},
{date:"2024-08-13", abi:"x86", id:"com.google.android.youtube"},
{date:"2024-08-13", abi:"x86", id:"org.thoughtcrime.securesms"},
{date:"2024-08-14", abi:"x86", id:"com.instagram.android"},
{date:"2024-08-14", abi:"x86", id:"com.pinterest"},
{date:"2024-08-15", abi:"x86", id:"com.clearchannel.iheartradio.controller"},
}

type app_type struct {
   date     string
   abi      string
   id       string
}
