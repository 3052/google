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

type GooglePlayClient struct {
   AuthData   *AuthData
   DeviceInfo *DeviceInfo

   // SessionFile if SessionFile is set then session will be saved to it after modification
   SessionFile string
}

var (
   err_GPTokenExpired = errors.New("unauthorized, gp token expired")

   httpClient = &http.Client{}
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

   _, err = client.toc()
   return
}

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

func doReq(r *http.Request) ([]byte, int, error) {
   res, err := httpClient.Do(r)
   if err != nil {
      return nil, 0, err
   }
   defer res.Body.Close()
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

func (client *GooglePlayClient) _doAuthedReq(r *http.Request) (_ *gpproto.Payload, err error) {
   client.setDefaultHeaders(r)
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
