package gplayapi

import (
   "154.pages.dev/google/gplayapi/gpproto"
   "errors"
   "net/http"
   "net/url"
   "strconv"
)

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

type GooglePlayClient struct {
   AuthData   *AuthData
   DeviceInfo *DeviceInfo

   // SessionFile if SessionFile is set then session will be saved to it after modification
   SessionFile string
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
