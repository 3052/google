package gplayapi

import (
   "154.pages.dev/google/gplayapi/gpproto"
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "os"
   "strconv"
)

var ErrMissingAppDeliveryData = errors.New("buy response is missing AppDeliveryData")

func (client *GooglePlayClient) GetBuyResponse(packageName string, version int) (*gpproto.BuyResponse, error) {
   params := &url.Values{}
   params.Set("ot", "1")
   params.Set("doc", packageName)
   params.Set("vc", strconv.Itoa(version))

   r, _ := http.NewRequest("POST", UrlPurchase+"?"+params.Encode(), nil)
   payload, err := client.doAuthedReq(r)
   if err != nil {
      return nil, err
   }
   if payload == nil {
      return nil, ErrNilPayload
   }
   return payload.BuyResponse, nil
}

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

func (client *GooglePlayClient) Purchase(packageName string, version int) (*gpproto.AndroidAppDeliveryData, error) {
   if version == 0 {
      app, err := client.GetAppDetails(packageName)
      if err != nil {
         return nil, err
      }
      version = app.VersionCode
   }

   _, _ = client.GetBuyResponse(packageName, version)

   res, err := client.GetDeliveryResponse(packageName, version)
   if err != nil {
      return nil, err
   }

   appDeliveryData := res.AppDeliveryData
   if appDeliveryData == nil {
      return nil, ErrMissingAppDeliveryData
   }
   return appDeliveryData, nil
}
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

// NewClient makes new client with Pixel3a config
func NewClient(email, aasToken string) (*GooglePlayClient, error) {
   return NewClientWithDeviceInfo(email, aasToken, Pixel3a)
}

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

func LoadSession(file string) (*GooglePlayClient, error) {
   return LoadSessionWithDeviceInfo(file, Pixel3a)
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
