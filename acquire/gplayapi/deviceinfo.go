package gplayapi

import (
	"acquire/gplayapi/gpproto"
	"fmt"
	"strconv"
	"strings"
)

var (
	androidGoogle         = ptrStr("android-google")
	falsePtr              = ptrBool(false)
	roaming               = "mobile-notroaming"
	glVersion       int32 = 196610
	glExtensions          = []string{"GL_AMD_compressed_ATC_texture", "GL_AMD_performance_monitor", "GL_ANDROID_extension_pack_es31a", "GL_APPLE_texture_2D_limited_npot", "GL_ARB_vertex_buffer_object", "GL_ARM_shader_framebuffer_fetch_depth_stencil", "GL_EXT_EGL_image_array", "GL_EXT_EGL_image_external_wrap_modes", "GL_EXT_EGL_image_storage", "GL_EXT_YUV_target", "GL_EXT_blend_func_extended", "GL_EXT_blit_framebuffer_params", "GL_EXT_buffer_storage", "GL_EXT_clip_control", "GL_EXT_clip_cull_distance", "GL_EXT_color_buffer_float", "GL_EXT_color_buffer_half_float", "GL_EXT_copy_image", "GL_EXT_debug_label", "GL_EXT_debug_marker", "GL_EXT_discard_framebuffer", "GL_EXT_disjoint_timer_query", "GL_EXT_draw_buffers_indexed", "GL_EXT_external_buffer", "GL_EXT_geometry_shader", "GL_EXT_gpu_shader5", "GL_EXT_memory_object", "GL_EXT_memory_object_fd", "GL_EXT_multisampled_render_to_texture", "GL_EXT_multisampled_render_to_texture2", "GL_EXT_primitive_bounding_box", "GL_EXT_protected_textures", "GL_EXT_robustness", "GL_EXT_sRGB", "GL_EXT_sRGB_write_control", "GL_EXT_shader_framebuffer_fetch", "GL_EXT_shader_io_blocks", "GL_EXT_shader_non_constant_global_initializers", "GL_EXT_tessellation_shader", "GL_EXT_texture_border_clamp", "GL_EXT_texture_buffer", "GL_EXT_texture_cube_map_array", "GL_EXT_texture_filter_anisotropic", "GL_EXT_texture_format_BGRA8888", "GL_EXT_texture_format_sRGB_override", "GL_EXT_texture_norm16", "GL_EXT_texture_sRGB_R8", "GL_EXT_texture_sRGB_decode", "GL_EXT_texture_type_2_10_10_10_REV", "GL_KHR_blend_equation_advanced", "GL_KHR_blend_equation_advanced_coherent", "GL_KHR_debug", "GL_KHR_no_error", "GL_KHR_robust_buffer_access_behavior", "GL_KHR_texture_compression_astc_hdr", "GL_KHR_texture_compression_astc_ldr", "GL_NV_shader_noperspective_interpolation", "GL_OES_EGL_image", "GL_OES_EGL_image_external", "GL_OES_EGL_image_external_essl3", "GL_OES_EGL_sync", "GL_OES_blend_equation_separate", "GL_OES_blend_func_separate", "GL_OES_blend_subtract", "GL_OES_compressed_ETC1_RGB8_texture", "GL_OES_compressed_paletted_texture", "GL_OES_depth24", "GL_OES_depth_texture", "GL_OES_depth_texture_cube_map", "GL_OES_draw_texture", "GL_OES_element_index_uint", "GL_OES_framebuffer_object", "GL_OES_get_program_binary", "GL_OES_matrix_palette", "GL_OES_packed_depth_stencil", "GL_OES_point_size_array", "GL_OES_point_sprite", "GL_OES_read_format", "GL_OES_rgb8_rgba8", "GL_OES_sample_shading", "GL_OES_sample_variables", "GL_OES_shader_image_atomic", "GL_OES_shader_multisample_interpolation", "GL_OES_standard_derivatives", "GL_OES_stencil_wrap", "GL_OES_surfaceless_context", "GL_OES_texture_3D", "GL_OES_texture_compression_astc", "GL_OES_texture_cube_map", "GL_OES_texture_env_crossbar", "GL_OES_texture_float", "GL_OES_texture_float_linear", "GL_OES_texture_half_float", "GL_OES_texture_half_float_linear", "GL_OES_texture_mirrored_repeat", "GL_OES_texture_npot", "GL_OES_texture_stencil8", "GL_OES_texture_storage_multisample_2d_array", "GL_OES_vertex_array_object", "GL_OES_vertex_half_float", "GL_OVR_multiview", "GL_OVR_multiview2", "GL_OVR_multiview_multisampled_render_to_texture", "GL_QCOM_alpha_test", "GL_QCOM_extended_get", "GL_QCOM_shader_framebuffer_fetch_noncoherent", "GL_QCOM_texture_foveated", "GL_QCOM_tiled_rendering"}
	sharedLibraries       = []string{"android.ext.services", "android.ext.shared", "android.test.base", "android.test.mock", "android.test.runner", "com.android.future.usb.accessory", "com.android.ims.rcsmanager", "com.android.location.provider", "com.android.media.remotedisplay", "com.android.mediadrm.signer", "com.google.android.camera.experimental2018", "com.google.android.dialer.support", "com.google.android.gms", "com.google.android.hardwareinfo", "com.google.android.lowpowermonitordevicefactory", "com.google.android.lowpowermonitordeviceinterface", "com.google.android.maps", "com.google.android.poweranomalydatafactory", "com.google.android.poweranomalydatamodeminterface", "com.qti.snapdragon.sdk.display", "com.qualcomm.embmslibrary", "com.qualcomm.qcrilhook", "com.qualcomm.qti.QtiTelephonyServicelibrary", "com.qualcomm.qti.imscmservice@1.0-java", "com.qualcomm.qti.lpa.uimlpalibrary", "com.qualcomm.qti.ltedirectdiscoverylibrary", "com.qualcomm.qti.remoteSimlock.uimremotesimlocklibrary", "com.qualcomm.qti.uim.uimservicelibrary", "com.quicinc.cne", "com.quicinc.cneapiclient", "com.verizon.embms", "com.verizon.provider", "com.vzw.apnlib", "javax.obex", "org.apache.http.legacy"}
	features              = []string{"android.hardware.audio.low_latency", "android.hardware.audio.output", "android.hardware.audio.pro", "android.hardware.bluetooth", "android.hardware.bluetooth_le", "android.hardware.camera", "android.hardware.camera.any", "android.hardware.camera.autofocus", "android.hardware.camera.capability.manual_post_processing", "android.hardware.camera.capability.manual_sensor", "android.hardware.camera.capability.raw", "android.hardware.camera.flash", "android.hardware.camera.front", "android.hardware.camera.level.full", "android.hardware.faketouch", "android.hardware.fingerprint", "android.hardware.location", "android.hardware.location.gps", "android.hardware.location.network", "android.hardware.microphone", "android.hardware.nfc", "android.hardware.nfc.any", "android.hardware.nfc.hce", "android.hardware.nfc.hcef", "android.hardware.opengles.aep", "android.hardware.ram.normal", "android.hardware.screen.landscape", "android.hardware.screen.portrait", "android.hardware.sensor.accelerometer", "android.hardware.sensor.assist", "android.hardware.sensor.barometer", "android.hardware.sensor.compass", "android.hardware.sensor.gyroscope", "android.hardware.sensor.hifi_sensors", "android.hardware.sensor.light", "android.hardware.sensor.proximity", "android.hardware.sensor.stepcounter", "android.hardware.sensor.stepdetector", "android.hardware.strongbox_keystore", "android.hardware.telephony", "android.hardware.telephony.carrierlock", "android.hardware.telephony.cdma", "android.hardware.telephony.euicc", "android.hardware.telephony.gsm", "android.hardware.touchscreen", "android.hardware.touchscreen.multitouch", "android.hardware.touchscreen.multitouch.distinct", "android.hardware.touchscreen.multitouch.jazzhand", "android.hardware.usb.accessory", "android.hardware.usb.host", "android.hardware.vulkan.compute", "android.hardware.vulkan.level", "android.hardware.vulkan.version", "android.hardware.wifi", "android.hardware.wifi.aware", "android.hardware.wifi.direct", "android.hardware.wifi.passpoint", "android.hardware.wifi.rtt", "android.software.activities_on_secondary_displays", "android.software.app_widgets", "android.software.autofill", "android.software.backup", "android.software.cant_save_state", "android.software.companion_device_setup", "android.software.connectionservice", "android.software.cts", "android.software.device_admin", "android.software.device_id_attestation", "android.software.file_based_encryption", "android.software.home_screen", "android.software.input_methods", "android.software.live_wallpaper", "android.software.managed_users", "android.software.midi", "android.software.picture_in_picture", "android.software.print", "android.software.securely_removes_users", "android.software.sip", "android.software.sip.voip", "android.software.verified_boot", "android.software.voice_recognizers", "android.software.webview", "com.google.android.apps.dialer.SUPPORTED", "com.google.android.apps.photos.PIXEL_2019_MIDYEAR_PRELOAD", "com.google.android.feature.EXCHANGE_6_2", "com.google.android.feature.GOOGLE_BUILD", "com.google.android.feature.GOOGLE_EXPERIENCE", "com.google.android.feature.PIXEL_2017_EXPERIENCE", "com.google.android.feature.PIXEL_2018_EXPERIENCE", "com.google.android.feature.PIXEL_2019_MIDYEAR_EXPERIENCE", "com.google.android.feature.PIXEL_EXPERIENCE", "com.google.android.feature.TURBO_PRELOAD", "com.google.android.feature.WELLBEING", "com.google.android.feature.ZERO_TOUCH", "com.google.hardware.camera.easel_2018", "com.verizon.hardware.telephony.ehrpd", "com.verizon.hardware.telephony.lte"}
	locales               = []string{"af", "af_ZA", "am", "am_ET", "ar", "ar_EG", "ar_SA", "ar_XB", "as", "ast", "az", "be", "be_BY", "bg", "bg_BG", "bh_IN", "bn", "bs", "ca", "ca_ES", "cs", "cs_CZ", "cy_GB", "da", "da_DK", "de", "de_DE", "el", "el_GR", "en", "en_AU", "en_CA", "en_GB", "en_IN", "en_US", "en_XA", "es", "es_ES", "es_US", "et", "et_EE", "eu", "fa", "fa_IR", "fi", "fi_FI", "fil", "fil_PH", "fr", "fr_CA", "fr_FR", "gl", "gl_ES", "gu", "hi", "hi_IN", "hr", "hr_HR", "hu", "hu_HU", "hy", "in", "in_ID", "is", "it", "it_IT", "iw", "iw_IL", "ja", "ja_JP", "ka", "kab_DZ", "kk", "km", "kn", "ko", "ko_KR", "ky", "lo", "lt", "lt_LT", "lv", "lv_LV", "mk", "ml", "mn", "mr", "ms", "ms_MY", "my", "nb", "nb_NO", "ne", "nl", "nl_NL", "or", "pa", "pa_IN", "pl", "pl_PL", "pt", "pt_BR", "pt_PT", "ro", "ro_RO", "ru", "ru_RU", "sc_IT", "si", "sk", "sk_SK", "sl", "sl_SI", "sq", "sr", "sr_Latn", "sr_RS", "sv", "sv_SE", "sw", "sw_TZ", "ta", "te", "th", "th_TH", "tr", "tr_TR", "uk", "uk_UA", "ur", "uz", "vi", "vi_VN", "zh_CN", "zh_HK", "zh_TW", "zu", "zu_ZA"}

	sargo = ptrStr("sargo")

	// _Pixel3a is default device which uses arm64 and Android 9 / SDK 28
	Pixel3a = &_DeviceInfo{
		_Build: &_DeviceBuildInfo{
			AndroidBuildProto: &gpproto.AndroidBuildProto{
				BuildProduct:   sargo,
				Radio:          ptrStr("g670-00011-190411-B-5457439"),
				Bootloader:     ptrStr("b4s4-0.1-5613380"),
				Carrier:        ptrStr("google"),
				Device:         sargo,
				Model:          ptrStr("Pixel 3a"),
				Manufacturer:   ptrStr("google"),
				Product:        sargo,
				Id:             ptrStr("PQ3B.190705.003"),
				SdkVersion:     ptrInt32(28),
				Client:         androidGoogle,
				GoogleServices: ptrInt32(203615028),
				OtaInstalled:   falsePtr,
			},
			_VersionRelease: 9,
		},
		_SimOperator:  "20815",
		_Platforms:    []string{"arm64-v8a", "armeabi-v7a", "armeabi"},
		_CellOperator: "334050",
		_Roaming:      roaming,
		_TimeZone:     "America/Mexico_City",
		_TouchScreen:  3,
		_Keyboard:     1,
		_Navigation:   1,
		_ScreenLayout: 2,
		_Screen: &_DeviceInfoScreen{
			_Density: 490,
			_Width:   1080,
			_Height:  2073,
		},
		_GLVersion:       glVersion,
		_GLExtensions:    glExtensions,
		_SharedLibraries: sharedLibraries,
		_Features:        features,
		_Locales:         locales,
	}
)

type (
	_DeviceInfo struct {
		_Build           *_DeviceBuildInfo
		_SimOperator     string
		_Platforms       []string
		_CellOperator    string
		_Roaming         string
		_TimeZone        string
		_TouchScreen     int32
		_Keyboard        int32
		_Navigation      int32
		_ScreenLayout    int32
		_Screen          *_DeviceInfoScreen
		_GLVersion       int32
		_GLExtensions    []string
		_SharedLibraries []string
		_Features        []string
		_Locales         []string
	}

	_DeviceBuildInfo struct {
		*gpproto.AndroidBuildProto
		_VersionRelease int
	}

	_DeviceInfoScreen struct {
		_Density int32
		_Width   int32
		_Height  int32
	}
)

func (i *_DeviceInfo) _GetAuthUserAgent() string {
	return fmt.Sprintf("GoogleAuth/1.4 %s %s", i._Build.GetDevice(), i._Build.GetId())
}

func (i *_DeviceInfo) _GetUserAgent() string {
	params := []string{
		"api=3",
		"versionCode=81582300",
		"sdk=28",
		"device=" + i._Build.GetDevice(),
		"hardware=" + i._Build.GetDevice(),
		"product=" + i._Build.GetProduct(),
		"platformVersionRelease=" + strconv.Itoa(i._Build._VersionRelease),
		"model=" + i._Build.GetModel(),
		"buildId=" + i._Build.GetId(),
		"isWideScreen=0",
		"supportedAbis=" + strings.Join(i._Platforms, ";"),
	}
	return "Android-Finsky/15.8.23-all [0] [PR] 259261889 (" + strings.Join(params, ",") + ")"
}

func (i *_DeviceInfo) _GetDeviceConfigProto() *gpproto.DeviceConfigurationProto {
	var mem int64 = 8589935000
	return &gpproto.DeviceConfigurationProto{
		TouchScreen:            &i._TouchScreen,
		Keyboard:               &i._Keyboard,
		Navigation:             &i._Navigation,
		ScreenLayout:           &i._ScreenLayout,
		HasHardKeyboard:        ptrBool(false),
		HasFiveWayNavigation:   ptrBool(false),
		ScreenDensity:          &i._Screen._Density,
		GlEsVersion:            &i._GLVersion,
		SystemSharedLibrary:    i._SharedLibraries,
		SystemAvailableFeature: i._Features,
		NativePlatform:         i._Platforms,
		ScreenWidth:            &i._Screen._Width,
		ScreenHeight:           &i._Screen._Height,
		SystemSupportedLocale:  i._Locales,
		GlExtension:            i._GLExtensions,
		DeviceClass:            ptrInt32(0),
		LowRamDevice:           ptrInt32(0),
		TotalMemoryBytes:       &mem,
		MaxNumOf_CPUCores:      ptrInt32(8),
		DeviceFeature:          i._GetDeviceFeatures(),
	}
}

func (i *_DeviceInfo) _GetDeviceFeatures() (ret []*gpproto.DeviceFeature) {
	var int0 int32 = 0
	for _, f := range i._Features {
		name := f
		ret = append(ret, &gpproto.DeviceFeature{Name: &name, Value: &int0})
	}
	return ret
}
