package gplayapi

import (
   "154.pages.dev/google/gplayapi/gpproto"
   "encoding/json"
   "errors"
   "google.golang.org/protobuf/proto"
   "io"
   "net/http"
   "net/url"
   "os"
   "strconv"
   "strings"
)

func NewClientWithDeviceInfo(email, aasToken string, deviceInfo *DeviceInfo) (client *GooglePlayClient, err error) {
   authData := &AuthData{
      Email:    email,
      AASToken: aasToken,
      Locale:   "en_GB",
   }
   client = &GooglePlayClient{AuthData: authData, DeviceInfo: deviceInfo}
   _, err = client.GenerateGsfID()
   if err != nil {
      return
   }
   deviceConfigRes, err := client.uploadDeviceConfig()
   if err != nil {
      return
   }
   authData.DeviceConfigToken = deviceConfigRes.GetUploadDeviceConfigToken()
   token, err := client.GenerateGPToken()
   if err != nil {
      return
   }
   authData.AuthToken = token
   return
}

func doReq(r *http.Request) ([]byte, int, error) {
   res, err := http.DefaultClient.Do(r)
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

type (
   App struct {
      PackageName        string
      AppInfo            *AppInfo
      CategoryImage      *gpproto.Image
      CategoryID         int
      CategoryName       string
      Changes            string
      ContainsAds        bool
      CoverImage         *gpproto.Image
      Description        string
      DeveloperName      string
      DisplayName        string
      DownloadString     string
      EarlyAccess        bool
      IconImage          *gpproto.Image
      InstantAppLink     string
      IsFree             bool
      IsSystem           bool
      LiveStreamUrl      string
      OfferDetails       map[string]string
      OfferType          int32
      Price              string
      PromotionStreamUrl string
      Screenshots        []*gpproto.Image
      ShareUrl           string
      ShortDescription   string
      Size               int64
      TargetSdk          int
      TestingProgram     *TestingProgram
      UpdatedOn          string
      VersionCode        int
      VersionName        string
      Video              *gpproto.Image
   }

   AppInfo struct {
      AppInfoMap map[string]string
   }

   TestingProgram struct {
      Image                    *gpproto.Image
      DisplayName              string
      Email                    string
      IsAvailable              bool
      IsSubscribedAndInstalled bool
   }

   TestingProgramStatus struct {
      Subscribed   bool
      Unsubscribed bool
   }
)
type GooglePlayClient struct {
   AuthData   *AuthData
   DeviceInfo *DeviceInfo

   // SessionFile if SessionFile is set then session will be saved to it after modification
   SessionFile string
}

var err_GPTokenExpired = errors.New("unauthorized, gp token expired")

func (client *GooglePlayClient) SaveSession(file string) error {
   f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
   if err != nil {
      return err
   }
   return json.NewEncoder(f).Encode(client.AuthData)
}

func LoadSessionWithDeviceInfo(file string, deviceInfo *DeviceInfo) (client *GooglePlayClient, err error) {
   f, err := os.Open(file)
   if err != nil {
      return
   }
   client = &GooglePlayClient{DeviceInfo: deviceInfo}
   err = json.NewDecoder(f).Decode(&client.AuthData)
   return
}
var ErrMissingAppDeliveryData = errors.New("buy response is missing AppDeliveryData")

func (client *GooglePlayClient) GetDeliveryResponse(packageName string, version int) (*gpproto.DeliveryResponse, error) {
   params := &url.Values{}
   params.Set("ot", "1")
   params.Set("doc", packageName)
   params.Set("vc", strconv.Itoa(version))

   r, _ := http.NewRequest("GET", UrlDelivery+"?"+params.Encode(), nil)
   payload, err := client.doAuthedReq(r)
   if err != nil {
      return nil, err
   }
   if payload == nil {
      return nil, ErrNilPayload
   }
   return payload.DeliveryResponse, nil
}

//goland:noinspection GoUnusedConst
const (
   ImageTypeAppScreenshot = iota + 1
   ImageTypePlayStorePageBackground
   ImageTypeYoutubeVideoLink
   ImageTypeAppIcon
   ImageTypeCategoryIcon
   ImageTypeYoutubeVideoThumbnail = 13

   UrlBase               = "https://android.clients.google.com"
   UrlFdfe               = UrlBase + "/fdfe"
   UrlAuth               = UrlBase + "/auth"
   UrlCheckIn            = UrlBase + "/checkin"
   UrlDetails            = UrlFdfe + "/details"
   UrlDelivery           = UrlFdfe + "/delivery"
   UrlPurchase           = UrlFdfe + "/purchase"
   UrlToc                = UrlFdfe + "/toc"
   UrlTosAccept          = UrlFdfe + "/acceptTos"
   UrlUploadDeviceConfig = UrlFdfe + "/uploadDeviceConfig"
)

var ErrNilPayload = errors.New("got nil payload from google play")
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

func (client *GooglePlayClient) _doAuthedReq(r *http.Request) (_ *gpproto.Payload, err error) {
   data := client.AuthData
   r.Header.Set("Authorization", "Bearer "+data.AuthToken)
   r.Header.Set("User-Agent", client.DeviceInfo.GetUserAgent())
   r.Header.Set("X-DFE-Device-Id", data.GsfID)
   r.Header.Set("Accept-Language", "en-GB")
   r.Header.Set(
      "X-DFE-Encoded-Targets",
      "CAESN/qigQYC2AMBFfUbyA7SM5Ij/CvfBoIDgxHqGP8R3xzIBvoQtBKFDZ4HAY4FrwSVMasHBO0O2Q8akgYRAQECAQO7AQEpKZ0CnwECAwRrAQYBr9PPAoK7sQMBAQMCBAkIDAgBAwEDBAICBAUZEgMEBAMLAQEBBQEBAcYBARYED+cBfS8CHQEKkAEMMxcBIQoUDwYHIjd3DQ4MFk0JWGYZEREYAQOLAYEBFDMIEYMBAgICAgICOxkCD18LGQKEAcgDBIQBAgGLARkYCy8oBTJlBCUocxQn0QUBDkkGxgNZQq0BZSbeAmIDgAEBOgGtAaMCDAOQAZ4BBIEBKUtQUYYBQscDDxPSARA1oAEHAWmnAsMB2wFyywGLAxol+wImlwOOA80CtwN26A0WjwJVbQEJPAH+BRDeAfkHK/ABASEBCSAaHQemAzkaRiu2Ad8BdXeiAwEBGBUBBN4LEIABK4gB2AFLfwECAdoENq0CkQGMBsIBiQEtiwGgA1zyAUQ4uwS8AwhsvgPyAcEDF27vApsBHaICGhl3GSKxAR8MC6cBAgItmQYG9QIeywLvAeYBDArLAh8HASI4ELICDVmVBgsY/gHWARtcAsMBpALiAdsBA7QBpAJmIArpByn0AyAKBwHTARIHAX8D+AMBcRIBBbEDmwUBMacCHAciNp0BAQF0OgQLJDuSAh54kwFSP0eeAQQ4M5EBQgMEmwFXywFo0gFyWwMcapQBBugBPUW2AVgBKmy3AR6PAbMBGQxrUJECvQR+8gFoWDsYgQNwRSczBRXQAgtRswEW0ALMAREYAUEBIG6yATYCRE8OxgER8gMBvQEDRkwLc8MBTwHZAUOnAXiiBakDIbYBNNcCIUmuArIBSakBrgFHKs0EgwV/G3AD0wE6LgECtQJ4xQFwFbUCjQPkBS6vAQqEAUZF3QIM9wEhCoYCQhXsBCyZArQDugIziALWAdIBlQHwBdUErQE6qQaSA4EEIvYBHir9AQVLmgMCApsCKAwHuwgrENsBAjNYswEVmgIt7QJnN4wDEnta+wGfAcUBxgEtEFXQAQWdAUAeBcwBAQM7rAEJATJ0LENrdh73A6UBhAE+qwEeASxLZUMhDREuH0CGARbd7K0GlQo",
   )
   r.Header.Set(
      "X-DFE-Phenotype",
      "H4sIAAAAAAAAAB3OO3KjMAAA0KRNuWXukBkBQkAJ2MhgAZb5u2GCwQZbCH_EJ77QHmgvtDtbv-Z9_H63zXXU0NVPB1odlyGy7751Q3CitlPDvFd8lxhz3tpNmz7P92CFw73zdHU2Ie0Ad2kmR8lxhiErTFLt3RPGfJQHSDy7Clw10bg8kqf2owLokN4SecJTLoSwBnzQSd652_MOf2d1vKBNVedzg4ciPoLz2mQ8efGAgYeLou-l-PXn_7Sna1MfhHuySxt-4esulEDp8Sbq54CPPKjpANW-lkU2IZ0F92LBI-ukCKSptqeq1eXU96LD9nZfhKHdtjSWwJqUm_2r6pMHOxk01saVanmNopjX3YxQafC4iC6T55aRbC8nTI98AF_kItIQAJb5EQxnKTO7TZDWnr01HVPxelb9A2OWX6poidMWl16K54kcu_jhXw-JSBQkVcD_fPsLSZu6joIBAAA",
   )
   r.Header.Set("X-DFE-Client-Id", "am-android-google")
   r.Header.Set("X-DFE-Network-Type", "4")
   r.Header.Set("X-DFE-Content-Filters", "")
   r.Header.Set("X-Limit-Ad-Tracking-Enabled", "false")
   r.Header.Set("X-Ad-Id", "LawadaMera")
   r.Header.Set("X-DFE-UserLanguages", "en_GB")
   r.Header.Set("X-DFE-Request-Params", "timeoutMs=4000")
   if data.DeviceCheckInConsistencyToken != "" {
      r.Header.Set("X-DFE-Device-Checkin-Consistency-Token", data.DeviceCheckInConsistencyToken)
   }
   if data.DeviceConfigToken != "" {
      r.Header.Set("X-DFE-Device-Config-Token", data.DeviceConfigToken)
   }
   if data.DFECookie != "" {
      r.Header.Set("X-DFE-Cookie", data.DFECookie)
   }
   if client.DeviceInfo.SimOperator != "" {
      r.Header.Set("X-DFE-MCCMNC", client.DeviceInfo.SimOperator)
   }
   b, status, err := doReq(r)
   if err != nil {
      return
   }
   if status == 401 {
      return nil, err_GPTokenExpired
   }
   resp := &gpproto.ResponseWrapper{}
   err = proto.Unmarshal(b, resp)
   if err != nil {
      return
   }
   return resp.Payload, nil
}

func (client *GooglePlayClient) doAuthedReq(r *http.Request) (res *gpproto.Payload, err error) {
   res, err = client._doAuthedReq(r)
   if err == err_GPTokenExpired {
      err = client.RegenerateGPToken()
      if err != nil {
         return
      }
      if client.SessionFile != "" {
         client.SaveSession(client.SessionFile)
      }
      res, err = client._doAuthedReq(r)
   }
   return
}

func (client *GooglePlayClient) RegenerateGPToken() (err error) {
   client.AuthData.AuthToken, err = client.GenerateGPToken()
   return
}
