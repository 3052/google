package acquire

import (
   "154.pages.dev/encoding/protobuf"
   "bytes"
   "errors"
   "fmt"
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "strconv"
   "strings"
)

func (c checkin) android_ID() (string, error) {
   v, ok := c.m.Fixed64(7)
   if ok {
      return strconv.FormatUint(v, 16), nil
   }
   return "", errors.New("androidId_")
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

func (client *_GooglePlayClient) checkIn() (*checkin, error) {
   c := Phone
   c.Native_Platform = "x86"
   // Checkin$AndroidCheckinRequest
   var m protobuf.Message
   // Logs$AndroidCheckinProto checkin_
   m.Add(4, func(m *protobuf.Message) {
      // Logs$AndroidBuildProto build_
      m.Add(1, func(m *protobuf.Message) {
         // int sdkVersion_
         // single APK valid range 14 - 28
         // multiple APK valid range 14 - math.MaxInt32
         m.Add_Varint(10, 28)
      })
      m.Add_Varint(18, 1)
   })
   // int version
   // valid range 2 - 3
   m.Add_Varint(14, 3)
   // Config$DeviceConfigurationProto deviceConfiguration_
   m.Add(18, func(m *protobuf.Message) {
      // int touchScreen
      m.Add_Varint(1, c.Touch_Screen)
      // int keyboard
      m.Add_Varint(2, c.Keyboard)
      // int navigation
      m.Add_Varint(3, c.Navigation)
      // int screenLayout
      m.Add_Varint(4, c.Screen_Layout)
      // boolean hasHardKeyboard
      m.Add_Bool(5, c.Has_Hard_Keyboard)
      // boolean hasFiveWayNavigation
      m.Add_Bool(6, c.Has_Five_Way_Navigation)
      // int screenDensity
      m.Add_Varint(7, c.Screen_Density)
      // int glEsVersion
      m.Add_Varint(8, c.GL_ES_Version)
      for _, library := range c.System_Shared_Library {
         // String[] systemSharedLibrary
         m.Add_String(9, library)
      }
      // String[] nativePlatform
      m.Add_String(11, c.Native_Platform)
      for _, extension := range c.GL_Extension {
         // String[] glExtension
         m.Add_String(15, extension)
      }
      // you cannot swap the next two lines:
      for _, feature := range c.System_Available_Feature {
         m.Add(26, func(m *protobuf.Message) {
            // String[] systemAvailableFeature
            m.Add_String(1, feature)
         })
      }
   })
   b = m.Append(nil)
   r, _ := http.NewRequest("POST", _UrlCheckIn, bytes.NewReader(b))
   r.Header.Set("User-Agent", "GoogleAuth/1.4 sargo PQ3B.190705.003")
   r.Header.Set("Content-Type", "application/x-protobuffer")
   r.Header.Set("Host", "android.clients.google.com")
   b, _, err := doReq(r)
   if err != nil {
      return nil, err
   }
   var check checkin
   check.m, err = protobuf.Consume(b)
   if err != nil {
      return nil, err
   }
   return &check, nil
}

func NewClientWithDeviceInfo(email, aasToken string) (*_GooglePlayClient, error) {
   authData := &_AuthData{
      _Email:    email,
      _AASToken: aasToken,
      _Locale:   "en_GB",
   }
   client := _GooglePlayClient{AuthData: authData}
   checkInResp, err := client.checkIn()
   if err != nil {
      return nil, err
   }
   client.AuthData.GsfID, err = checkInResp.android_ID()
   if err != nil {
      return nil, err
   }
   err = client.uploadDeviceConfig()
   if err != nil {
      return nil, err
   }
   token, err := client._GenerateGPToken()
   if err != nil {
      return nil, err
   }
   authData._AuthToken = token
   return &client, nil
}

type checkin struct {
   m protobuf.Message
}

type _AuthData struct {
   GsfID                          string
   _AASToken                      string
   _AuthToken                     string
   _Email                         string
   _Locale                        string
}

type _GooglePlayClient struct {
   AuthData    *_AuthData
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
   if ok {
      return token, nil
   }
   return "", errors.New("authentication failed: could not generate oauth token")
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
