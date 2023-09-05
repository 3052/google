package main

import (
   "154.pages.dev/encoding/protobuf"
   "bytes"
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Header["Authorization"] = []string{"GoogleLogin auth=aghNhuY8mfnldimVPcpmITL5Q2e6-HLpqLBU08wN8MqBroyW0EzpdCgO58px5BrayQ6rvQ."}
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Header["Host"] = []string{"android.clients.google.com"}
   req.Header["User-Agent"] = []string{"Android-Finsky/11.9.30-all%20%5B0%5D%20%5BPR%5D%20215272473 (api=3,versionCode=81193000,sdk=21,device=generic_x86,hardware=ranchu,product=sdk_google_phone_x86,platformVersionRelease=5.0.2,model=Android%20SDK%20built%20for%20x86,buildId=LSY66K,isWideScreen=0,supportedAbis=x86)"}
   req.Header["X-Dfe-Device-Id"] = []string{"337ff139c042f382"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "android.clients.google.com"
   req.URL.Path = "/fdfe/uploadDeviceConfig"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(bytes.NewReader(device_config.Append(nil)))
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

var device_config = protobuf.Message{
   protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\b\x03\x10\x01\x18\x02 \x02(\x000\x018\x90\x03@\x81\x80\fJ\x13android.test.runnerJ com.android.future.usb.accessoryJ\x1fcom.android.media.remotedisplayJ\x1dcom.android.location.providerJ\njavax.obexJ com.google.android.media.effectsJ\x17com.google.android.mapsJ\x16com.google.android.gmsJ\x1bcom.android.mediadrm.signerR\x16android.software.printR'android.hardware.touchscreen.multitouchR%android.hardware.sensor.accelerometerR\x1candroid.software.app_widgetsR\x1dandroid.software.device_adminR\x1bandroid.hardware.camera.anyR\x17android.hardware.cameraR\x1eandroid.hardware.usb.accessoryR0android.hardware.touchscreen.multitouch.distinctR\"android.software.voice_recognizersR\x1aandroid.hardware.faketouchR\x1fandroid.hardware.sensor.compassR\x1eandroid.software.input_methodsR\"android.software.connectionserviceR\x19android.hardware.locationR!android.hardware.screen.landscapeR\x18android.software.webviewR!android.hardware.location.networkR\x17android.software.backupR\x1eandroid.software.managed_usersR\x1aandroid.hardware.bluetoothR android.hardware.screen.portraitR\x1candroid.hardware.touchscreenR\x1dandroid.hardware.audio.outputR\x1fandroid.software.live_wallpaperR\x1candroid.software.home_screenR0android.hardware.touchscreen.multitouch.jazzhandR\x1bandroid.hardware.microphoneR!android.hardware.camera.autofocusZ\x03x86`\xb8\bh\xf0\x10r\x00r\x02car\x02dar\x02jar\x02nbr\x02der\x02bgr\x02thr\x02fir\x02hir\x02vir\x02skr\x02ukr\x02elr\x02nlr\x02plr\x02slr\x03filr\x02kor\x02ror\x02arr\x02frr\x02hrr\x02srr\x02trr\x02csr\x02esr\x02itr\x02ltr\x02ptr\x02hur\x02rur\x02lvr\x02svr\x05en-CAr\x05fr-CAr\x05uk-UAr\x05en-ZAr\x05en-GBr\x05en-IEr\x05ar-EGr\x05en-SGr\x05th-THr\x05fi-FIr\x05sl-SIr\x05sk-SKr\x05zh-CNr\x05hi-INr\x05en-INr\x05vi-VNr\x05ro-ROr\x05hr-HRr\x05sr-RSr\x05en-USr\x05es-USr\x05lt-LTr\x05pt-PTr\x05en-AUr\x05hu-HUr\x05lv-LVr\x05zh-TWr\x05en-NZr\x05nl-BEr\x05fr-BEr\x05de-DEr\x05sv-SEr\x05bg-BGr\x05de-CHr\x05fr-CHr\x05it-CHr\x06fil-PHr\x05de-LIr\x05da-DKr\x05ar-ILr\x05nl-NLr\x05pl-PLr\x05nb-NOr\x05ja-JPr\x05pt-BRr\x05fr-FRr\x05el-GRr\x05ko-KRr\x05tr-TRr\x05ca-ESr\x05es-ESr\x05de-ATr\x05it-ITr\x05ru-RUr\x05cs-CZr\x02enr\x02idr\x02far\x02kar\x02par\x02tar\x02ber\x02ner\x02ter\x02afr\x02sir\x02kkr\x02mkr\x02glr\x02mlr\x02amr\x02kmr\x02bnr\x02inr\x02knr\x02mnr\x02lor\x02sqr\x02mrr\asr-Latnr\x02urr\x02isr\x02msr\x02etr\x02eur\x02gur\x02zur\x02iwr\x02swr\x02hyr\x02kyr\x02myr\x02azr\x02uzr\x05zh-HKz\x1eANDROID_EMU_CHECKSUM_HELPER_v1z\x12ANDROID_EMU_dma_v1z ANDROID_EMU_gles_max_version_3_1z\x1dANDROID_EMU_host_side_tracingz\x1cANDROID_EMU_sync_buffer_dataz GL_APPLE_texture_format_BGRA8888z\x19GL_EXT_color_buffer_floatz\x1eGL_EXT_color_buffer_half_floatz\x13GL_EXT_debug_markerz\x1eGL_EXT_texture_format_BGRA8888z#GL_KHR_texture_compression_astc_ldrz\x10GL_OES_EGL_imagez\x19GL_OES_EGL_image_externalz\x1fGL_OES_EGL_image_external_essl3z\x0fGL_OES_EGL_syncz\x1eGL_OES_blend_equation_separatez\x1aGL_OES_blend_func_separatez\x15GL_OES_blend_subtractz\x17GL_OES_byte_coordinatesz#GL_OES_compressed_ETC1_RGB8_texturez\"GL_OES_compressed_paletted_texturez\x0eGL_OES_depth24z\x0eGL_OES_depth32z\x14GL_OES_depth_texturez\x13GL_OES_draw_texturez\x19GL_OES_element_index_uintz\x18GL_OES_fbo_render_mipmapz\x19GL_OES_framebuffer_objectz\x1bGL_OES_packed_depth_stencilz\x17GL_OES_point_size_arrayz\x13GL_OES_point_spritez\x11GL_OES_rgb8_rgba8z\x17GL_OES_single_precisionz\x0fGL_OES_stencil1z\x0fGL_OES_stencil4z\x0fGL_OES_stencil8z\x13GL_OES_stencil_wrapz\x17GL_OES_texture_cube_mapz\x1bGL_OES_texture_env_crossbarz\x14GL_OES_texture_floatz\x1bGL_OES_texture_float_linearz\x19GL_OES_texture_half_floatz GL_OES_texture_half_float_linearz\x1dGL_OES_texture_mirored_repeatz\x13GL_OES_texture_npotz\x1aGL_OES_vertex_array_objectz\x18GL_OES_vertex_half_float\x90\x01\xb0\x03\x98\x01\x00\xa0\x01\x80\xa0\xca\xf5\x05\xa8\x01\x04\xd2\x01\x18\n\x16android.software.print\xd2\x01)\n'android.hardware.touchscreen.multitouch\xd2\x01'\n%android.hardware.sensor.accelerometer\xd2\x01\x1e\n\x1candroid.software.app_widgets\xd2\x01\x1f\n\x1dandroid.software.device_admin\xd2\x01\x1d\n\x1bandroid.hardware.camera.any\xd2\x01\x19\n\x17android.hardware.camera\xd2\x01 \n\x1eandroid.hardware.usb.accessory\xd2\x012\n0android.hardware.touchscreen.multitouch.distinct\xd2\x01$\n\"android.software.voice_recognizers\xd2\x01\x1c\n\x1aandroid.hardware.faketouch\xd2\x01!\n\x1fandroid.hardware.sensor.compass\xd2\x01 \n\x1eandroid.software.input_methods\xd2\x01$\n\"android.software.connectionservice\xd2\x01\x1b\n\x19android.hardware.location\xd2\x01#\n!android.hardware.screen.landscape\xd2\x01\x1a\n\x18android.software.webview\xd2\x01#\n!android.hardware.location.network\xd2\x01\x19\n\x17android.software.backup\xd2\x01 \n\x1eandroid.software.managed_users\xd2\x01\x1c\n\x1aandroid.hardware.bluetooth\xd2\x01\"\n android.hardware.screen.portrait\xd2\x01\x1e\n\x1candroid.hardware.touchscreen\xd2\x01\x1f\n\x1dandroid.hardware.audio.output\xd2\x01!\n\x1fandroid.software.live_wallpaper\xd2\x01\x1e\n\x1candroid.software.home_screen\xd2\x012\n0android.hardware.touchscreen.multitouch.jazzhand\xd2\x01\x1d\n\x1bandroid.hardware.microphone\xd2\x01#\n!android.hardware.camera.autofocus\xd8\x01\x02")},
   protobuf.Field{Number: 1, Type: 2, Value: protobuf.Maybe{
      protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(3)},
      protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1)},
      protobuf.Field{Number: 3, Type: 0, Value: protobuf.Varint(2)},
      protobuf.Field{Number: 4, Type: 0, Value: protobuf.Varint(2)},
      protobuf.Field{Number: 5, Type: 0, Value: protobuf.Varint(0)},
      protobuf.Field{Number: 6, Type: 0, Value: protobuf.Varint(1)},
      protobuf.Field{Number: 7, Type: 0, Value: protobuf.Varint(400)},
      protobuf.Field{Number: 8, Type: 0, Value: protobuf.Varint(196609)},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("android.test.runner")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 12, Type: 1, Value: protobuf.Fixed64(8371739161332442222)},
         protobuf.Field{Number: 12, Type: 5, Value: protobuf.Fixed32(1915647091)},
         protobuf.Field{Number: 14, Type: 5, Value: protobuf.Fixed32(1919250030)},
      }},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.android.future.usb.accessory")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.android.media.remotedisplay")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.android.location.provider")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("javax.obex")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.google.android.media.effects")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.google.android.maps")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.google.android.gms")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.android.mediadrm.signer")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.print")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen.multitouch")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.accelerometer")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.app_widgets")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.device_admin")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.camera.any")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 12, Type: 1, Value: protobuf.Fixed64(7507048032877306990)},
         protobuf.Field{Number: 12, Type: 1, Value: protobuf.Fixed64(7146761200619447410)},
         protobuf.Field{Number: 12, Type: 1, Value: protobuf.Fixed64(8750037977858729325)},
      }},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.camera")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.usb.accessory")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen.multitouch.distinct")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.voice_recognizers")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.faketouch")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.compass")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.input_methods")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.connectionservice")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.location")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.screen.landscape")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.webview")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.location.network")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.backup")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.managed_users")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.bluetooth")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.screen.portrait")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.audio.output")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.live_wallpaper")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.home_screen")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen.multitouch.jazzhand")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.microphone")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 12, Type: 1, Value: protobuf.Fixed64(7507048032877306990)},
         protobuf.Field{Number: 12, Type: 1, Value: protobuf.Fixed64(7867337140998726770)},
         protobuf.Field{Number: 13, Type: 1, Value: protobuf.Fixed64(7308901739622527587)},
      }},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.camera.autofocus")},
      protobuf.Field{Number: 11, Type: 2, Value: protobuf.Bytes("x86")},
      protobuf.Field{Number: 12, Type: 0, Value: protobuf.Varint(1080)},
      protobuf.Field{Number: 13, Type: 0, Value: protobuf.Varint(2160)},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ca")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("da")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ja")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("nb")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("de")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("bg")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("th")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("fi")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("hi")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 13, Type: 0, Value: protobuf.Varint(105)},
      }},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("vi")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sk")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("uk")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("el")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("nl")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("pl")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 14, Type: 0, Value: protobuf.Varint(108)},
      }},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sl")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("fil")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ko")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ro")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ar")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("fr")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("hr")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 13, Type: 0, Value: protobuf.Varint(114)},
      }},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sr")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("tr")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("cs")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("es")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("it")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("lt")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("pt")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 14, Type: 0, Value: protobuf.Varint(116)},
      }},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("hu")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 13, Type: 0, Value: protobuf.Varint(117)},
      }},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ru")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("lv")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sv")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("en-CA")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 12, Type: 5, Value: protobuf.Fixed32(1094921582)},
      }},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("fr-CA")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("uk-UA")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 14, Type: 5, Value: protobuf.Fixed32(1096101227)},
      }},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("en-ZA")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 12, Type: 5, Value: protobuf.Fixed32(1096428910)},
      }},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("en-GB")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 12, Type: 5, Value: protobuf.Fixed32(1111960942)},
      }},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("en-IE")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 12, Type: 5, Value: protobuf.Fixed32(1162423662)},
      }},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ar-EG")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("en-SG")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 12, Type: 5, Value: protobuf.Fixed32(1196633454)},
      }},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("th-TH")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("fi-FI")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sl-SI")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sk-SK")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("zh-CN")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("hi-IN")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("en-IN")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 12, Type: 5, Value: protobuf.Fixed32(1313418606)},
      }},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("vi-VN")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ro-RO")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("hr-HR")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sr-RS")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("en-US")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 12, Type: 5, Value: protobuf.Fixed32(1398091118)},
      }},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("es-US")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 12, Type: 5, Value: protobuf.Fixed32(1398091123)},
      }},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("lt-LT")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("pt-PT")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("en-AU")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 12, Type: 5, Value: protobuf.Fixed32(1430334830)},
      }},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("hu-HU")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("lv-LV")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("zh-TW")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("en-NZ")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 12, Type: 5, Value: protobuf.Fixed32(1515072878)},
      }},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("nl-BE")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("fr-BE")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("de-DE")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sv-SE")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("bg-BG")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("de-CH")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("fr-CH")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("it-CH")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("fil-PH")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("de-LI")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("da-DK")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ar-IL")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("nl-NL")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("pl-PL")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("nb-NO")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ja-JP")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("pt-BR")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("fr-FR")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("el-GR")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 12, Type: 5, Value: protobuf.Fixed32(1380396396)},
      }},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ko-KR")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("tr-TR")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ca-ES")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("es-ES")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 12, Type: 5, Value: protobuf.Fixed32(1397042547)},
      }},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("de-AT")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("it-IT")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ru-RU")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("cs-CZ")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("en")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("id")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("fa")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ka")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("pa")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 14, Type: 0, Value: protobuf.Varint(97)},
      }},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ta")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("be")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ne")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("te")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("af")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("si")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("kk")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("mk")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("gl")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ml")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("am")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("km")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("bn")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("in")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("kn")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("mn")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("lo")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sq")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("mr")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sr-Latn")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ur")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("is")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ms")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("et")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("eu")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("gu")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("zu")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("iw")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sw")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("hy")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 13, Type: 0, Value: protobuf.Varint(121)},
      }},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ky")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("my")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("az")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("uz")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("zh-HK")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("ANDROID_EMU_CHECKSUM_HELPER_v1")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("ANDROID_EMU_dma_v1")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("ANDROID_EMU_gles_max_version_3_1")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("ANDROID_EMU_host_side_tracing")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("ANDROID_EMU_sync_buffer_data")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_APPLE_texture_format_BGRA8888")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_color_buffer_float")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_color_buffer_half_float")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_debug_marker")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_texture_format_BGRA8888")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_KHR_texture_compression_astc_ldr")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_EGL_image")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_EGL_image_external")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_EGL_image_external_essl3")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_EGL_sync")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_blend_equation_separate")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_blend_func_separate")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_blend_subtract")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_byte_coordinates")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_compressed_ETC1_RGB8_texture")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_compressed_paletted_texture")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_depth24")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_depth32")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_depth_texture")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_draw_texture")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_element_index_uint")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_fbo_render_mipmap")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_framebuffer_object")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_packed_depth_stencil")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_point_size_array")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_point_sprite")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_rgb8_rgba8")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_single_precision")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_stencil1")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_stencil4")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_stencil8")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_stencil_wrap")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_texture_cube_map")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_texture_env_crossbar")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_texture_float")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_texture_float_linear")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_texture_half_float")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_texture_half_float_linear")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_texture_mirored_repeat")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_texture_npot")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_vertex_array_object")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_vertex_half_float")},
      protobuf.Field{Number: 18, Type: 0, Value: protobuf.Varint(432)},
      protobuf.Field{Number: 19, Type: 0, Value: protobuf.Varint(0)},
      protobuf.Field{Number: 20, Type: 0, Value: protobuf.Varint(1588760576)},
      protobuf.Field{Number: 21, Type: 0, Value: protobuf.Varint(4)},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Bytes("\n\x16android.software.print")},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.print")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Bytes("\n'android.hardware.touchscreen.multitouch")},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen.multitouch")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Bytes("\n%android.hardware.sensor.accelerometer")},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.accelerometer")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Bytes("\n\x1candroid.software.app_widgets")},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.app_widgets")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Bytes("\n\x1dandroid.software.device_admin")},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.device_admin")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Bytes("\n\x1bandroid.hardware.camera.any")},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera.any")},
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Maybe{
            protobuf.Field{Number: 12, Type: 1, Value: protobuf.Fixed64(7507048032877306990)},
            protobuf.Field{Number: 12, Type: 1, Value: protobuf.Fixed64(7146761200619447410)},
            protobuf.Field{Number: 12, Type: 1, Value: protobuf.Fixed64(8750037977858729325)},
         }},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Bytes("\n\x17android.hardware.camera")},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Bytes("\n\x1eandroid.hardware.usb.accessory")},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.usb.accessory")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Bytes("\n0android.hardware.touchscreen.multitouch.distinct")},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen.multitouch.distinct")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Bytes("\n\"android.software.voice_recognizers")},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.voice_recognizers")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Bytes("\n\x1aandroid.hardware.faketouch")},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.faketouch")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Bytes("\n\x1fandroid.hardware.sensor.compass")},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.compass")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Bytes("\n\x1eandroid.software.input_methods")},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.input_methods")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Bytes("\n\"android.software.connectionservice")},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.connectionservice")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Bytes("\n\x19android.hardware.location")},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.location")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Bytes("\n!android.hardware.screen.landscape")},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.screen.landscape")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Bytes("\n\x18android.software.webview")},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.webview")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Bytes("\n!android.hardware.location.network")},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.location.network")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Bytes("\n\x17android.software.backup")},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.backup")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Bytes("\n\x1eandroid.software.managed_users")},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.managed_users")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Bytes("\n\x1aandroid.hardware.bluetooth")},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.bluetooth")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Bytes("\n android.hardware.screen.portrait")},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.screen.portrait")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Bytes("\n\x1candroid.hardware.touchscreen")},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Bytes("\n\x1dandroid.hardware.audio.output")},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.audio.output")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Bytes("\n\x1fandroid.software.live_wallpaper")},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.live_wallpaper")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Bytes("\n\x1candroid.software.home_screen")},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.home_screen")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Bytes("\n0android.hardware.touchscreen.multitouch.jazzhand")},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen.multitouch.jazzhand")},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Bytes("\n\x1bandroid.hardware.microphone")},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.microphone")},
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Maybe{
            protobuf.Field{Number: 12, Type: 1, Value: protobuf.Fixed64(7507048032877306990)},
            protobuf.Field{Number: 12, Type: 1, Value: protobuf.Fixed64(7867337140998726770)},
            protobuf.Field{Number: 13, Type: 1, Value: protobuf.Fixed64(7308901739622527587)},
         }},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Bytes("\n!android.hardware.camera.autofocus")},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera.autofocus")},
      }},
      protobuf.Field{Number: 27, Type: 0, Value: protobuf.Varint(2)},
   }},
   protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("\n\"\b\x80\x90\xda\xc8\xe0\xc5F\x12\aAndroid\"\x0fepc.tmobile.com")},
   protobuf.Field{Number: 4, Type: 2, Value: protobuf.Maybe{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\b\x80\x90\xda\xc8\xe0\xc5F\x12\aAndroid\"\x0fepc.tmobile.com")},
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Maybe{
         protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(310260000000000)},
         protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("Android")},
         protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("epc.tmobile.com")},
         protobuf.Field{Number: 4, Type: 2, Value: protobuf.Maybe{
            protobuf.Field{Number: 12, Type: 5, Value: protobuf.Fixed32(1949197168)},
            protobuf.Field{Number: 13, Type: 5, Value: protobuf.Fixed32(1818845807)},
            protobuf.Field{Number: 12, Type: 5, Value: protobuf.Fixed32(1836016430)},
         }},
      }},
   }},
   protobuf.Field{Number: 6, Type: 2, Value: protobuf.Bytes("\x12\vgeneric_x86\x1aOgeneric_x86/sdk_google_phone_x86/generic_x86:5.0.2/LSY66K/6695550:eng/test-keys!Ƥ\xfdB\xd1E\x01\x10(\x012\aunknown")},
   protobuf.Field{Number: 6, Type: 2, Value: protobuf.Maybe{
      protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("generic_x86")},
      protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("generic_x86/sdk_google_phone_x86/generic_x86:5.0.2/LSY66K/6695550:eng/test-keys")},
      protobuf.Field{Number: 4, Type: 1, Value: protobuf.Fixed64(1153279744657958086)},
      protobuf.Field{Number: 5, Type: 0, Value: protobuf.Varint(1)},
      protobuf.Field{Number: 6, Type: 2, Value: protobuf.Bytes("unknown")},
   }},
   protobuf.Field{Number: 7, Type: 2, Value: protobuf.Bytes("\n+5FrW36Wo1fTkY-lbBq_OlRgTTgKhyxKe6k3oUjoZTKs")},
   protobuf.Field{Number: 7, Type: 2, Value: protobuf.Maybe{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("5FrW36Wo1fTkY-lbBq_OlRgTTgKhyxKe6k3oUjoZTKs")},
   }},
}
