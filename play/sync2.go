package play

import (
	"154.pages.dev/protobuf"
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
)

func (h Header) Sync2(Device) error {
	req := new(http.Request)
	req.Header = make(http.Header)
	req.Header["Accept-Language"] = []string{"en-US"}
	req.Header["Connection"] = []string{"Keep-Alive"}
	req.Header["Content-Type"] = []string{"application/x-protobuf"}
	req.Header["Host"] = []string{"android.clients.google.com"}
	req.Header["User-Agent"] = []string{"Android-Finsky/29.4.13-19%20%5B0%5D%20%5BPR%5D%20429618493 (api=3,versionCode=82941300,sdk=26,device=generic_x86,hardware=ranchu,product=sdk_gphone_x86,platformVersionRelease=8.0.0,model=Android%20SDK%20built%20for%20x86,buildId=OSR1.170720.005,isWideScreen=0,supportedAbis=x86)"}
	req.Header["X-Dfe-Build-Fingerprint"] = []string{"google/sdk_gphone_x86/generic_x86:8.0.0/OSR1.170720.005/4205617:userdebug/dev-keys"}
	req.Header["X-Dfe-Client-Id"] = []string{"am-unknown"}
	req.Header["X-Dfe-Encoded-Targets"] = []string{"CAE"}
	req.Header["X-Dfe-Mccmnc"] = []string{"310260"}
	req.Header["X-Dfe-Network-Type"] = []string{"3"}
	req.Header["X-Dfe-Request-Params"] = []string{"timeoutMs=4000"}
	req.Method = "POST"
	req.ProtoMajor = 1
	req.ProtoMinor = 1
	req.URL = new(url.URL)
	req.URL.Host = "android.clients.google.com"
	req.URL.Path = "/fdfe/sync"
	req.URL.Scheme = "https"
	req.Body = io.NopCloser(bytes.NewReader(sync_body.Append(nil)))
	req.Header.Set(h.Device_ID())
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}
	return nil
}

var sync_body = protobuf.Message{
	protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
		protobuf.Field{Number: 9, Type: 2, Value: protobuf.Prefix{
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("310260")},
			protobuf.Field{Number: 2, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
					protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("Android")},
					protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("epc.tmobile.com")},
				}},
			}},
		}},
	}},
	protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
		protobuf.Field{Number: 8, Type: 2, Value: protobuf.Bytes("")},
	}},
	protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
		protobuf.Field{Number: 10, Type: 2, Value: protobuf.Prefix{
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.proximity")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.accelerometer")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.faketouch")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.usb.accessory")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.backup")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen.multitouch")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.print")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.voice_recognizers")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.picture_in_picture")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.fingerprint")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.relative_humidity")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.android.feature.GOOGLE_BUILD")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.telephony.gsm")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.audio.output")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.screen.portrait")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.ambient_temperature")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.home_screen")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.microphone")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.autofill")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.compass")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen.multitouch.jazzhand")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.barometer")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.app_widgets")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.input_methods")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.light")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.device_admin")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.screen.landscape")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.managed_users")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.webview")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera.any")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.connectionservice")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen.multitouch.distinct")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.location.network")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.cts")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.android.apps.dialer.SUPPORTED")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.live_wallpaper")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.android.feature.GOOGLE_EXPERIENCE")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.android.feature.EXCHANGE_6_2")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.location.gps")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.midi")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.location")},
			}},
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
				protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.telephony")},
			}},
			protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("com.google.android.media.effects")},
			protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("com.android.location.provider")},
			protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("com.android.future.usb.accessory")},
			protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("android.ext.shared")},
			protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("javax.obex")},
			protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("com.google.android.gms")},
			protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("android.ext.services")},
			protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("android.test.runner")},
			protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("com.google.android.dialer.support")},
			protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("com.google.android.maps")},
			protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("org.apache.http.legacy")},
			protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("com.android.media.remotedisplay")},
			protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("com.android.mediadrm.signer")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("af")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("am")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("ar")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("ar-EG")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("ar-IL")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("az")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("be")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("bg")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("bg-BG")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("bn")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("ca")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("ca-ES")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("cs")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("cs-CZ")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("da")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("da-DK")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("de")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("de-AT")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("de-CH")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("de-DE")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("de-LI")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("el")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("el-GR")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("en")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("en-AU")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("en-CA")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("en-GB")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("en-IE")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("en-IN")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("en-NZ")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("en-SG")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("en-US")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("en-ZA")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("es")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("es-ES")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("es-US")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("et")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("eu")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("fa")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("fi")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("fi-FI")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("fil")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("fil-PH")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("fr")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("fr-BE")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("fr-CA")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("fr-CH")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("fr-FR")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("gl")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("gu")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("hi")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("hi-IN")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("hr")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("hu")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("hu-HU")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("hy")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("id")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("in")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("is")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("it")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("it-CH")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("it-IT")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("iw")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("ja")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("ja-JP")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("ka")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("kk")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("km")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("kn")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("ko")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("ko-KR")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("ky")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("lo")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("lt")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("lt-LT")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("lv")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("lv-LV")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("mk")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("ml")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("mn")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("mr")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("ms")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("my")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("nb")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("nb-NO")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("ne")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("nl")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("nl-BE")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("nl-NL")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("pa")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("pl")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("pl-PL")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("pt")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("pt-BR")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("pt-PT")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("ro")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("ro-RO")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("ru")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("ru-RU")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("si")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("sk")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("sk-SK")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("sl")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("sl-SI")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("sq")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("sr")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("sr-Latn")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("sr-RS")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("sv")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("sv-SE")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("sw")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("ta")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("te")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("th")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("th-TH")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("tr")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("tr-TR")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("uk")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("uk-UA")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("ur")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("uz")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("vi")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("vi-VN")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("zh-CN")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("zh-HK")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("zh-TW")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("zu")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("ANDROID_EMU_CHECKSUM_HELPER_v1")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("ANDROID_EMU_dma_v1")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("ANDROID_EMU_gles_max_version_3_1")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("ANDROID_EMU_host_side_tracing")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("ANDROID_EMU_native_sync_v2")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("ANDROID_EMU_native_sync_v3")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("ANDROID_EMU_native_sync_v4")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("ANDROID_EMU_sync_buffer_data")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_APPLE_texture_format_BGRA8888")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_EXT_color_buffer_float")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_EXT_color_buffer_half_float")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_EXT_debug_marker")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_EXT_texture_format_BGRA8888")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_KHR_texture_compression_astc_ldr")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_EGL_image")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_EGL_image_external")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_EGL_image_external_essl3")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_EGL_sync")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_blend_equation_separate")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_blend_func_separate")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_blend_subtract")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_byte_coordinates")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_compressed_ETC1_RGB8_texture")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_compressed_paletted_texture")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_depth24")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_depth32")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_depth_texture")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_draw_texture")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_element_index_uint")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_fbo_render_mipmap")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_framebuffer_object")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_packed_depth_stencil")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_point_size_array")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_point_sprite")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_rgb8_rgba8")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_single_precision")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_stencil1")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_stencil4")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_stencil8")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_stencil_wrap")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_texture_cube_map")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_texture_env_crossbar")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_texture_float")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_texture_float_linear")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_texture_half_float")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_texture_half_float_linear")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_texture_mirored_repeat")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_texture_npot")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_vertex_array_object")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("GL_OES_vertex_half_float")},
		}},
	}},
	protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
		protobuf.Field{Number: 11, Type: 2, Value: protobuf.Prefix{
			protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(2)},
			protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1)},
			protobuf.Field{Number: 3, Type: 0, Value: protobuf.Varint(1)},
		}},
	}},
	protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
		protobuf.Field{Number: 12, Type: 2, Value: protobuf.Prefix{
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("Google")},
			protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("Android SDK built for x86")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("generic_x86")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("sdk_gphone_x86")},
			protobuf.Field{Number: 5, Type: 2, Value: protobuf.Bytes("google")},
		}},
	}},
	protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
		protobuf.Field{Number: 13, Type: 2, Value: protobuf.Bytes("")},
	}},
	protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
		protobuf.Field{Number: 14, Type: 2, Value: protobuf.Prefix{
			protobuf.Field{Number: 1, Type: 1, Value: protobuf.Fixed64(0)},
		}},
	}},
	protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
		protobuf.Field{Number: 15, Type: 2, Value: protobuf.Prefix{
			protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(0)},
			protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1586618368)},
			protobuf.Field{Number: 3, Type: 0, Value: protobuf.Varint(4)},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("x86")},
		}},
	}},
	protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
		protobuf.Field{Number: 16, Type: 2, Value: protobuf.Prefix{
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("GMT-07:00")},
		}},
	}},
	protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
		protobuf.Field{Number: 18, Type: 2, Value: protobuf.Prefix{
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("am-unknown")},
			protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("play-ms-android-google")},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("play-ad-ms-android-google")},
		}},
	}},
	protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
		protobuf.Field{Number: 19, Type: 2, Value: protobuf.Prefix{
			protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(82941300)},
		}},
	}},
	protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
		protobuf.Field{Number: 20, Type: 2, Value: protobuf.Prefix{
			protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(3)},
			protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1080)},
			protobuf.Field{Number: 3, Type: 0, Value: protobuf.Varint(1920)},
			protobuf.Field{Number: 4, Type: 0, Value: protobuf.Varint(2)},
			protobuf.Field{Number: 5, Type: 0, Value: protobuf.Varint(420)},
		}},
	}},
	protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
		protobuf.Field{Number: 21, Type: 2, Value: protobuf.Prefix{
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("google/sdk_gphone_x86/generic_x86:8.0.0/OSR1.170720.005/4205617:userdebug/dev-keys")},
			protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(26)},
			protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("")},
			protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("")},
			protobuf.Field{Number: 5, Type: 2, Value: protobuf.Bytes("")},
			protobuf.Field{Number: 6, Type: 0, Value: protobuf.Varint(196609)},
		}},
	}},
}
