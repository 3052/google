package play

var Native_Platforms = map[string]string{
   // com.google.android.youtube
   "0": "x86",
   // com.miui.weather2
   "1": "armeabi-v7a",
   // com.kakaogames.twodin
   "2": "arm64-v8a",
}

// These can use default values, but they must all be included
type Config struct {
   Touch_Screen uint64 // 1
   Keyboard uint64 // 2
   Navigation uint64 // 3
   Screen_Layout uint64 // 4
   Has_Hard_Keyboard bool // 5
   Has_Five_Way_Navigation bool // 6
   Screen_Density uint64 // 7
   GL_ES_Version uint64 // 8
   System_Shared_Library []string // 9
   Native_Platform string // 11
   GL_Extension []string // 15
   System_Available_Feature []string // 26
}

var Phone = Config{
   System_Available_Feature: []string{
      // app.source.getcontact
      "android.hardware.location.gps",
      // br.com.rodrigokolb.realdrum
      "android.software.midi",
      // com.app.xt
      "android.hardware.camera.front",
      // com.cabify.rider
      "android.hardware.camera.flash",
      // com.clearchannel.iheartradio.controller
      "android.hardware.microphone",
      // com.google.android.apps.walletnfcrel
      "android.software.device_admin",
      // com.google.android.youtube
      "android.hardware.touchscreen",
      "android.hardware.wifi",
      // com.madhead.tos.zh
      "android.hardware.sensor.accelerometer",
      // com.miHoYo.GenshinImpact
      "android.hardware.opengles.aep",
      // com.pinterest
      "android.hardware.camera",
      "android.hardware.location",
      "android.hardware.screen.portrait",
      // com.supercell.brawlstars
      "android.hardware.touchscreen.multitouch",
      // com.sygic.aura
      "android.hardware.location.network",
      // com.xiaomi.smarthome
      "android.hardware.bluetooth",
      "android.hardware.bluetooth_le",
      "android.hardware.camera.autofocus",
      "android.hardware.usb.host",
      // kr.sira.metal
      "android.hardware.sensor.compass",
      // org.thoughtcrime.securesms
      "android.hardware.telephony",
      // org.videolan.vlc
      "android.hardware.screen.landscape",
   },
   System_Shared_Library: []string{
      // com.amctve.amcfullepisodes
      "org.apache.http.legacy",
      // com.binance.dev
      "android.test.runner",
      // com.miui.weather2
      "global-miui11-empty.jar",
   },
   GL_Extension: []string{
      // com.instagram.android
      "GL_OES_compressed_ETC1_RGB8_texture",
      // com.kakaogames.twodin
      "GL_KHR_texture_compression_astc_ldr",
   },
   // com.axis.drawingdesk.v3
   // valid range 0x03_00_01 - math.MaxInt32
   GL_ES_Version: 0xFF_FF_FF,
}
