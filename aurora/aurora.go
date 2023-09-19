package aurora

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
)

type Aurora_OSS struct {
   Auth_Token string `json:"authToken"`
   GSF_ID string `json:"gsfId"`
}

func (Aurora_OSS) Marshal() ([]byte, error) {
   body, err := func() ([]byte, error) {
      m := map[string]string{
         "Build.BOOTLOADER": "unknown",
         "Build.BRAND": "generic_x86",
         "Build.DEVICE": "generic_x86",
         "Build.FINGERPRINT": "generic_x86/sdk_google_phone_x86/generic_x86:5.0.2/LSY66K/6695550:eng/test-keys",
         "Build.HARDWARE": "ranchu",
         "Build.MANUFACTURER": "unknown",
         "Build.MODEL": "Android SDK built for x86",
         "Build.PRODUCT": "sdk_google_phone_x86",
         "Build.RADIO": "",
         "Build.VERSION.SDK_INT": "21",
         "CellOperator": "310",
         "Client": "android-google",
         "Features": "android.software.print",
         "GL.Extensions": "",
         "GL.Version": "196609",
         "GSF.version": "203615037",
         "HasFiveWayNavigation": "true",
         "HasHardKeyboard": "false",
         "Keyboard": "1",
         "Locales": "en",
         "Navigation": "2",
         "Platforms": "x86",
         "Roaming": "mobile-notroaming",
         "Screen.Density": "400",
         "Screen.Height": "2040",
         "Screen.Width": "1080",
         "ScreenLayout": "2",
         "SharedLibraries": "",
         "SimOperator": "38",
         "TimeZone": "UTC-10",
         "TouchScreen": "3",
         "Vending.versionString": "22.0.17-21 [0] [PR] 332555730",
      }
      return json.Marshal(m)
   }()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://auroraoss.com/api/auth", bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("User-Agent", "com.aurora.store-4.3.1-49")
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   return io.ReadAll(res.Body)
}

func (a *Aurora_OSS) Unmarshal(text []byte) error {
   return json.Unmarshal(text, a)
}
