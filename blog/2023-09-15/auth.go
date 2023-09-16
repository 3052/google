package main

import (
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "auroraoss.com"
   req.URL.Path = "/api/auth"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(req_body)
   req.Header["User-Agent"] = []string{"com.aurora.store-4.3.1-49"}
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res_body, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(res_body)
}

var req_body = strings.NewReader(`
{
   "TouchScreen": "3",
   "Keyboard": "1",
   "Navigation": "2",
   "ScreenLayout": "2",
   "HasHardKeyboard": "false",
   "HasFiveWayNavigation": "true",
   "Screen.Density": "400",
   "GL.Version": "196609",
   "SharedLibraries": "android.test.runner,com.android.future.usb.accessory,com.android.media.remotedisplay,com.android.location.provider,javax.obex,com.google.android.media.effects,com.google.android.maps,com.google.android.gms,com.android.mediadrm.signer",
   "Platforms": "x86",
   "GL.Extensions": ",ANDROID_EMU_CHECKSUM_HELPER_v1,ANDROID_EMU_dma_v1,ANDROID_EMU_gles_max_version_3_1,ANDROID_EMU_host_side_tracing,ANDROID_EMU_sync_buffer_data,GL_APPLE_texture_format_BGRA8888,GL_EXT_color_buffer_float,GL_EXT_color_buffer_half_float,GL_EXT_debug_marker,GL_EXT_texture_format_BGRA8888,GL_KHR_texture_compression_astc_ldr,GL_OES_EGL_image,GL_OES_EGL_image_external,GL_OES_EGL_image_external_essl3,GL_OES_EGL_sync,GL_OES_blend_equation_separate,GL_OES_blend_func_separate,GL_OES_blend_subtract,GL_OES_byte_coordinates,GL_OES_compressed_ETC1_RGB8_texture,GL_OES_compressed_paletted_texture,GL_OES_depth24,GL_OES_depth32,GL_OES_depth_texture,GL_OES_draw_texture,GL_OES_element_index_uint,GL_OES_fbo_render_mipmap,GL_OES_framebuffer_object,GL_OES_packed_depth_stencil,GL_OES_point_size_array,GL_OES_point_sprite,GL_OES_rgb8_rgba8,GL_OES_single_precision,GL_OES_stencil1,GL_OES_stencil4,GL_OES_stencil8,GL_OES_stencil_wrap,GL_OES_texture_cube_map,GL_OES_texture_env_crossbar,GL_OES_texture_float,GL_OES_texture_float_linear,GL_OES_texture_half_float,GL_OES_texture_half_float_linear,GL_OES_texture_mirored_repeat,GL_OES_texture_npot,GL_OES_vertex_array_object,GL_OES_vertex_half_float",
   "Features": "android.software.print,android.hardware.touchscreen.multitouch,android.hardware.sensor.accelerometer,android.software.app_widgets,android.software.device_admin,android.hardware.camera.any,android.hardware.camera,android.hardware.usb.accessory,android.hardware.touchscreen.multitouch.distinct,android.software.voice_recognizers,android.hardware.faketouch,android.hardware.sensor.compass,android.software.input_methods,android.software.connectionservice,android.hardware.location,android.hardware.screen.landscape,android.software.webview,android.hardware.location.network,android.software.backup,android.software.managed_users,android.hardware.bluetooth,android.hardware.screen.portrait,android.hardware.touchscreen,android.hardware.audio.output,android.software.live_wallpaper,android.software.home_screen,android.hardware.touchscreen.multitouch.jazzhand,android.hardware.microphone,android.hardware.camera.autofocus",
   "Build.BOOTLOADER": "unknown",
   "Build.BRAND": "generic_x86",
   "Build.DEVICE": "generic_x86",
   "Build.FINGERPRINT": "generic_x86/sdk_google_phone_x86/generic_x86:5.0.2/LSY66K/6695550:eng/test-keys",
   "Build.HARDWARE": "ranchu",
   
   "Build.MANUFACTURER": "unknown",
   "Build.MODEL": "Android SDK built for x86",
   "Build.PRODUCT": "sdk_google_phone_x86",
   "Build.RADIO": "",
   "Build.VERSION.RELEASE": "5.0.2",
   "Build.VERSION.SDK_INT": "21",
   "CellOperator": "310",
   "Client": "android-google",
   "GSF.version": "203615037",
   "Locales": "ca,da,ja,nb,de,bg,th,fi,hi,vi,sk,uk,el,nl,pl,sl,fil,ko,ro,ar,fr,hr,sr,tr,cs,es,it,lt,pt,hu,ru,lv,sv,en_CA,fr_CA,uk_UA,en_ZA,en_GB,en_IE,ar_EG,en_SG,th_TH,fi_FI,sl_SI,sk_SK,zh_CN,hi_IN,en_IN,vi_VN,ro_RO,hr_HR,sr_RS,en_US,es_US,lt_LT,pt_PT,en_AU,hu_HU,lv_LV,zh_TW,en_NZ,nl_BE,fr_BE,de_DE,sv_SE,bg_BG,de_CH,fr_CH,it_CH,fil_PH,de_LI,da_DK,ar_IL,nl_NL,pl_PL,nb_NO,ja_JP,pt_BR,fr_FR,el_GR,ko_KR,tr_TR,ca_ES,es_ES,de_AT,it_IT,ru_RU,cs_CZ,en,id,fa,ka,pa,ta,be,ne,te,af,si,kk,mk,gl,ml,am,km,bn,in,kn,mn,lo,sq,mr,or,sr_Latn,ur,as,bs,is,ms,et,eu,gu,zu,iw,sw,hy,ky,my,az,uz,es_419,zh_HK,kab,ast,ckb,ia,sc,he,eo,so,ht,kmr,yue,en_XC",
   "Roaming": "mobile-notroaming",
   "Screen.Height": "2040",
   "Screen.Width": "1080",
   "SimOperator": "38",
   "TimeZone": "UTC-10",
   "UserReadableName": "generic_x86-default",
   "Vending.version": "82201710",
   "Vending.versionString": "22.0.17-21 [0] [PR] 332555730"
}
`)
