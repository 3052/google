package gplayapi

import (
   "154.pages.dev/encoding/protobuf"
   "acquire/gplayapi/gpproto"
   "bytes"
   "errors"
   "fmt"
   "google.golang.org/protobuf/proto"
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "strings"
)

type _AuthData struct {
   GsfID                          string
   _AASToken                      string
   _AuthToken                     string
   _DeviceCheckInConsistencyToken string
   _DeviceConfigToken             string
   _Email                         string
   _Locale                        string
}

type _GooglePlayClient struct {
   AuthData    *_AuthData
   _DeviceInfo *_DeviceInfo
}

func doReq(r *http.Request) ([]byte, int, error) {
   {
      b, err := httputil.DumpRequest(r, true)
      if err != nil {
         return nil, 0, err
      }
      fmt.Printf("%q\n\n", b)
   }
   res, err := httpClient.Do(r)
   if err != nil {
      return nil, 0, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, 0, errors.New(res.Status)
   }
   b, err := io.ReadAll(res.Body)
   return b, res.StatusCode, err
}

func ptrBool(b bool) *bool {
   return &b
}

func ptrStr(str string) *string {
   return &str
}

func ptrInt32(i int32) *int32 {
   return &i
}

func parseResponse(res string) map[string]string {
   ret := map[string]string{}
   for _, ln := range strings.Split(res, "\n") {
      keyVal := strings.SplitN(ln, "=", 2)
      if len(keyVal) >= 2 {
         ret[keyVal[0]] = keyVal[1]
      }
   }
   return ret
}

func (client *_GooglePlayClient) _GenerateGPToken() (string, error) {
   params := &url.Values{}
   params.Set("Token", client.AuthData._AASToken)
   params.Set("app", "com.android.vending")
   params.Set("client_sig", "38918a453d07199354f8b19af05ec6562ced5788")
   params.Set("service", "oauth2:https://www.googleapis.com/auth/googleplay")
   r, _ := http.NewRequest("POST", _UrlAuth+"?"+params.Encode(), nil)
   b, _, err := doReq(r)
   if err != nil {
      return "", nil
   }
   resp := parseResponse(string(b))
   token, ok := resp["Auth"]
   if !ok {
      return "", errors.New("authentication failed: could not generate oauth token")
   }
   return token, nil
}

func NewClientWithDeviceInfo(email, aasToken string, deviceInfo *_DeviceInfo) (client *_GooglePlayClient, err error) {
   authData := &_AuthData{
      _Email:    email,
      _AASToken: aasToken,
      _Locale:   "en_GB",
   }
   client = &_GooglePlayClient{AuthData: authData, _DeviceInfo: deviceInfo}
   checkInResp, err := client.checkIn()
   if err != nil {
      return
   }
   client.AuthData.GsfID = fmt.Sprintf("%x", checkInResp.GetAndroidId())
   client.AuthData._DeviceCheckInConsistencyToken = checkInResp.GetDeviceCheckinConsistencyToken()
   deviceConfigRes, err := client.uploadDeviceConfig()
   if err != nil {
      return
   }
   authData._DeviceConfigToken = deviceConfigRes.GetUploadDeviceConfigToken()
   token, err := client._GenerateGPToken()
   if err != nil {
      return
   }
   authData._AuthToken = token
   return
}

const _UrlBase = "https://android.clients.google.com"
const _UrlFdfe = _UrlBase + "/fdfe"
const _UrlAuth = _UrlBase + "/auth"
const _UrlCheckIn = _UrlBase + "/checkin"
const _UrlUploadDeviceConfig = _UrlFdfe + "/uploadDeviceConfig"

var (
   err_GPTokenExpired = errors.New("unauthorized, gp token expired")

   httpClient = &http.Client{}
)

func (client *_GooglePlayClient) checkIn() (resp *gpproto.AndroidCheckinResponse, err error) {
   b := checkin_body.Append(nil)
   r, _ := http.NewRequest("POST", _UrlCheckIn, bytes.NewReader(b))
   r.Header.Set("User-Agent", client._DeviceInfo._GetAuthUserAgent())
   r.Header.Set("Content-Type", "application/x-protobuffer")
   r.Header.Set("Host", "android.clients.google.com")
   b, _, err = doReq(r)
   if err != nil {
      return
   }
   resp = &gpproto.AndroidCheckinResponse{}
   err = proto.Unmarshal(b, resp)
   return
}

var checkin_body = protobuf.Message{
   protobuf.Field{Number: 4, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("PQ3B.190705.003")},
         protobuf.Field{Number: 2, Type: 2, Value: protobuf.Bytes("sargo")},
         protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("google")},
         protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("g670-00011-190411-B-5457439")},
         protobuf.Field{Number: 5, Type: 2, Value: protobuf.Bytes("b4s4-0.1-5613380")},
         protobuf.Field{Number: 6, Type: 2, Value: protobuf.Bytes("android-google")},
         protobuf.Field{Number: 7, Type: 0, Value: protobuf.Varint(1694054582)},
         protobuf.Field{Number: 8, Type: 0, Value: protobuf.Varint(203615028)},
         protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("sargo")},
         protobuf.Field{Number: 10, Type: 0, Value: protobuf.Varint(28)},
         protobuf.Field{Number: 11, Type: 2, Value: protobuf.Bytes("Pixel 3a")},
         protobuf.Field{Number: 12, Type: 2, Value: protobuf.Bytes("google")},
         protobuf.Field{Number: 13, Type: 2, Value: protobuf.Bytes("sargo")},
         protobuf.Field{Number: 14, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      protobuf.Field{Number: 6, Type: 2, Value: protobuf.Bytes("334050")},
      protobuf.Field{Number: 7, Type: 2, Value: protobuf.Bytes("20815")},
      protobuf.Field{Number: 8, Type: 2, Value: protobuf.Bytes("mobile-notroaming")},
      protobuf.Field{Number: 9, Type: 0, Value: protobuf.Varint(0)},
   }},
   protobuf.Field{Number: 14, Type: 0, Value: protobuf.Varint(3)},
   protobuf.Field{Number: 18, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(3)},
      protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1)},
      protobuf.Field{Number: 3, Type: 0, Value: protobuf.Varint(1)},
      protobuf.Field{Number: 4, Type: 0, Value: protobuf.Varint(2)},
      protobuf.Field{Number: 5, Type: 0, Value: protobuf.Varint(0)},
      protobuf.Field{Number: 6, Type: 0, Value: protobuf.Varint(0)},
      protobuf.Field{Number: 7, Type: 0, Value: protobuf.Varint(490)},
      protobuf.Field{Number: 8, Type: 0, Value: protobuf.Varint(196610)},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("android.ext.services")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("android.ext.shared")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("android.test.base")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("android.test.mock")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("android.test.runner")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.android.future.usb.accessory")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.android.ims.rcsmanager")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.android.location.provider")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.android.media.remotedisplay")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.android.mediadrm.signer")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.google.android.camera.experimental2018")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.google.android.dialer.support")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.google.android.gms")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.google.android.hardwareinfo")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.google.android.lowpowermonitordevicefactory")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.google.android.lowpowermonitordeviceinterface")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.google.android.maps")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.google.android.poweranomalydatafactory")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.google.android.poweranomalydatamodeminterface")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.qti.snapdragon.sdk.display")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.qualcomm.embmslibrary")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.qualcomm.qcrilhook")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.qualcomm.qti.QtiTelephonyServicelibrary")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.qualcomm.qti.imscmservice@1.0-java")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.qualcomm.qti.lpa.uimlpalibrary")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.qualcomm.qti.ltedirectdiscoverylibrary")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.qualcomm.qti.remoteSimlock.uimremotesimlocklibrary")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.qualcomm.qti.uim.uimservicelibrary")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.quicinc.cne")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.quicinc.cneapiclient")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.verizon.embms")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.verizon.provider")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("com.vzw.apnlib")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("javax.obex")},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Bytes("org.apache.http.legacy")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.audio.low_latency")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.audio.output")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.audio.pro")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.bluetooth")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.bluetooth_le")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.camera")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.camera.any")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.camera.autofocus")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.camera.capability.manual_post_processing")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.camera.capability.manual_sensor")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.camera.capability.raw")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.camera.flash")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.camera.front")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.camera.level.full")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.faketouch")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.fingerprint")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.location")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.location.gps")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.location.network")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.microphone")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.nfc")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.nfc.any")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.nfc.hce")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.nfc.hcef")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.opengles.aep")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.ram.normal")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.screen.landscape")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.screen.portrait")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.accelerometer")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.assist")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.barometer")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.compass")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.gyroscope")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.hifi_sensors")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.light")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.proximity")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.stepcounter")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.stepdetector")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.strongbox_keystore")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.telephony")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.telephony.carrierlock")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.telephony.cdma")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.telephony.euicc")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.telephony.gsm")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen.multitouch")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen.multitouch.distinct")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen.multitouch.jazzhand")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.usb.accessory")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.usb.host")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.vulkan.compute")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.vulkan.level")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.vulkan.version")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.wifi")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.wifi.aware")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.wifi.direct")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.wifi.passpoint")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.hardware.wifi.rtt")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.activities_on_secondary_displays")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.app_widgets")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.autofill")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.backup")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.cant_save_state")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.companion_device_setup")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.connectionservice")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.cts")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.device_admin")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.device_id_attestation")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.file_based_encryption")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.home_screen")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.input_methods")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.live_wallpaper")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.managed_users")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.midi")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.picture_in_picture")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.print")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.securely_removes_users")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.sip")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.sip.voip")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.verified_boot")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.voice_recognizers")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("android.software.webview")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("com.google.android.apps.dialer.SUPPORTED")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("com.google.android.apps.photos.PIXEL_2019_MIDYEAR_PRELOAD")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("com.google.android.feature.EXCHANGE_6_2")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("com.google.android.feature.GOOGLE_BUILD")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("com.google.android.feature.GOOGLE_EXPERIENCE")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("com.google.android.feature.PIXEL_2017_EXPERIENCE")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("com.google.android.feature.PIXEL_2018_EXPERIENCE")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("com.google.android.feature.PIXEL_2019_MIDYEAR_EXPERIENCE")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("com.google.android.feature.PIXEL_EXPERIENCE")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("com.google.android.feature.TURBO_PRELOAD")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("com.google.android.feature.WELLBEING")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("com.google.android.feature.ZERO_TOUCH")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("com.google.hardware.camera.easel_2018")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("com.verizon.hardware.telephony.ehrpd")},
      protobuf.Field{Number: 10, Type: 2, Value: protobuf.Bytes("com.verizon.hardware.telephony.lte")},
      protobuf.Field{Number: 11, Type: 2, Value: protobuf.Bytes("arm64-v8a")},
      protobuf.Field{Number: 11, Type: 2, Value: protobuf.Bytes("armeabi-v7a")},
      protobuf.Field{Number: 11, Type: 2, Value: protobuf.Bytes("armeabi")},
      protobuf.Field{Number: 12, Type: 0, Value: protobuf.Varint(1080)},
      protobuf.Field{Number: 13, Type: 0, Value: protobuf.Varint(2073)},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("af")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("af_ZA")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("am")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("am_ET")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ar")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ar_EG")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ar_SA")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ar_XB")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("as")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ast")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("az")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("be")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("be_BY")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("bg")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("bg_BG")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("bh_IN")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("bn")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("bs")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ca")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ca_ES")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("cs")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("cs_CZ")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("cy_GB")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("da")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("da_DK")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("de")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("de_DE")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("el")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("el_GR")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("en")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("en_AU")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("en_CA")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("en_GB")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("en_IN")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("en_US")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("en_XA")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("es")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("es_ES")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("es_US")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("et")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("et_EE")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("eu")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("fa")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("fa_IR")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("fi")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("fi_FI")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("fil")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("fil_PH")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("fr")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("fr_CA")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("fr_FR")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("gl")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("gl_ES")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("gu")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("hi")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("hi_IN")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("hr")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("hr_HR")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("hu")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("hu_HU")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("hy")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("in")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("in_ID")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("is")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("it")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("it_IT")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("iw")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("iw_IL")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ja")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ja_JP")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ka")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("kab_DZ")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("kk")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("km")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("kn")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ko")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ko_KR")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ky")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("lo")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("lt")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("lt_LT")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("lv")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("lv_LV")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("mk")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ml")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("mn")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("mr")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ms")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ms_MY")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("my")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("nb")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("nb_NO")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ne")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("nl")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("nl_NL")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("or")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("pa")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("pa_IN")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("pl")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("pl_PL")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("pt")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("pt_BR")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("pt_PT")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ro")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ro_RO")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ru")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ru_RU")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sc_IT")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("si")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sk")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sk_SK")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sl")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sl_SI")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sq")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sr")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sr_Latn")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sr_RS")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sv")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sv_SE")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sw")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("sw_TZ")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ta")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("te")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("th")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("th_TH")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("tr")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("tr_TR")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("uk")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("uk_UA")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("ur")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("uz")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("vi")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("vi_VN")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("zh_CN")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("zh_HK")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("zh_TW")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("zu")},
      protobuf.Field{Number: 14, Type: 2, Value: protobuf.Bytes("zu_ZA")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_AMD_compressed_ATC_texture")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_AMD_performance_monitor")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_ANDROID_extension_pack_es31a")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_APPLE_texture_2D_limited_npot")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_ARB_vertex_buffer_object")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_ARM_shader_framebuffer_fetch_depth_stencil")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_EGL_image_array")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_EGL_image_external_wrap_modes")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_EGL_image_storage")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_YUV_target")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_blend_func_extended")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_blit_framebuffer_params")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_buffer_storage")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_clip_control")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_clip_cull_distance")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_color_buffer_float")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_color_buffer_half_float")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_copy_image")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_debug_label")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_debug_marker")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_discard_framebuffer")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_disjoint_timer_query")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_draw_buffers_indexed")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_external_buffer")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_geometry_shader")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_gpu_shader5")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_memory_object")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_memory_object_fd")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_multisampled_render_to_texture")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_multisampled_render_to_texture2")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_primitive_bounding_box")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_protected_textures")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_robustness")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_sRGB")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_sRGB_write_control")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_shader_framebuffer_fetch")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_shader_io_blocks")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_shader_non_constant_global_initializers")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_tessellation_shader")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_texture_border_clamp")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_texture_buffer")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_texture_cube_map_array")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_texture_filter_anisotropic")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_texture_format_BGRA8888")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_texture_format_sRGB_override")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_texture_norm16")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_texture_sRGB_R8")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_texture_sRGB_decode")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_EXT_texture_type_2_10_10_10_REV")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_KHR_blend_equation_advanced")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_KHR_blend_equation_advanced_coherent")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_KHR_debug")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_KHR_no_error")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_KHR_robust_buffer_access_behavior")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_KHR_texture_compression_astc_hdr")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_KHR_texture_compression_astc_ldr")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_NV_shader_noperspective_interpolation")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_EGL_image")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_EGL_image_external")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_EGL_image_external_essl3")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_EGL_sync")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_blend_equation_separate")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_blend_func_separate")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_blend_subtract")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_compressed_ETC1_RGB8_texture")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_compressed_paletted_texture")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_depth24")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_depth_texture")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_depth_texture_cube_map")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_draw_texture")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_element_index_uint")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_framebuffer_object")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_get_program_binary")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_matrix_palette")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_packed_depth_stencil")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_point_size_array")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_point_sprite")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_read_format")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_rgb8_rgba8")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_sample_shading")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_sample_variables")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_shader_image_atomic")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_shader_multisample_interpolation")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_standard_derivatives")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_stencil_wrap")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_surfaceless_context")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_texture_3D")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_texture_compression_astc")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_texture_cube_map")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_texture_env_crossbar")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_texture_float")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_texture_float_linear")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_texture_half_float")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_texture_half_float_linear")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_texture_mirrored_repeat")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_texture_npot")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_texture_stencil8")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_texture_storage_multisample_2d_array")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_vertex_array_object")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OES_vertex_half_float")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OVR_multiview")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OVR_multiview2")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_OVR_multiview_multisampled_render_to_texture")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_QCOM_alpha_test")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_QCOM_extended_get")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_QCOM_shader_framebuffer_fetch_noncoherent")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_QCOM_texture_foveated")},
      protobuf.Field{Number: 15, Type: 2, Value: protobuf.Bytes("GL_QCOM_tiled_rendering")},
      protobuf.Field{Number: 16, Type: 0, Value: protobuf.Varint(0)},
      protobuf.Field{Number: 19, Type: 0, Value: protobuf.Varint(0)},
      protobuf.Field{Number: 20, Type: 0, Value: protobuf.Varint(8589935000)},
      protobuf.Field{Number: 21, Type: 0, Value: protobuf.Varint(8)},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.audio.low_latency")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.audio.output")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.audio.pro")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.bluetooth")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.bluetooth_le")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera.any")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera.autofocus")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera.capability.manual_post_processing")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera.capability.manual_sensor")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera.capability.raw")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera.flash")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera.front")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.camera.level.full")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.faketouch")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.fingerprint")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.location")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.location.gps")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.location.network")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.microphone")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.nfc")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.nfc.any")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.nfc.hce")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.nfc.hcef")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.opengles.aep")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.ram.normal")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.screen.landscape")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.screen.portrait")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.accelerometer")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.assist")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.barometer")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.compass")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.gyroscope")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.hifi_sensors")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.light")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.proximity")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.stepcounter")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.sensor.stepdetector")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.strongbox_keystore")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.telephony")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.telephony.carrierlock")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.telephony.cdma")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.telephony.euicc")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.telephony.gsm")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen.multitouch")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen.multitouch.distinct")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.touchscreen.multitouch.jazzhand")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.usb.accessory")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.usb.host")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.vulkan.compute")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.vulkan.level")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.vulkan.version")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.wifi")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.wifi.aware")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.wifi.direct")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.wifi.passpoint")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.hardware.wifi.rtt")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.activities_on_secondary_displays")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.app_widgets")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.autofill")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.backup")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.cant_save_state")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.companion_device_setup")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.connectionservice")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.cts")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.device_admin")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.device_id_attestation")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.file_based_encryption")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.home_screen")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.input_methods")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.live_wallpaper")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.managed_users")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.midi")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.picture_in_picture")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.print")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.securely_removes_users")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.sip")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.sip.voip")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.verified_boot")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.voice_recognizers")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("android.software.webview")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.android.apps.dialer.SUPPORTED")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.android.apps.photos.PIXEL_2019_MIDYEAR_PRELOAD")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.android.feature.EXCHANGE_6_2")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.android.feature.GOOGLE_BUILD")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.android.feature.GOOGLE_EXPERIENCE")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.android.feature.PIXEL_2017_EXPERIENCE")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.android.feature.PIXEL_2018_EXPERIENCE")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.android.feature.PIXEL_2019_MIDYEAR_EXPERIENCE")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.android.feature.PIXEL_EXPERIENCE")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.android.feature.TURBO_PRELOAD")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.android.feature.WELLBEING")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.android.feature.ZERO_TOUCH")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.google.hardware.camera.easel_2018")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.verizon.hardware.telephony.ehrpd")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 26, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("com.verizon.hardware.telephony.lte")},
         protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
      }},
   }},
}
