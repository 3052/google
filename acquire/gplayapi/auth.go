package gplayapi

import (
   "bytes"
   "errors"
   "fmt"
   "github.com/Juby210/gplayapi-go/gpproto"
   "google.golang.org/protobuf/proto"
   "net/http"
   "net/url"
   "strconv"
)

type _AuthData struct {
   _Email                         string
   _AASToken                      string
   _AuthToken                     string
   GsfID                         string
   _DeviceCheckInConsistencyToken string
   _DeviceConfigToken             string
   _DFECookie                     string
   _Locale                        string
}

func (client *_GooglePlayClient) _GenerateGsfID() (gsfID string, err error) {
   req := client._DeviceInfo._GenerateAndroidCheckInRequest()
   checkInResp, err := client.checkIn(req)
   if err != nil {
      return
   }
   gsfID = fmt.Sprintf("%x", checkInResp.GetAndroidId())
   client.AuthData.GsfID = gsfID
   client.AuthData._DeviceCheckInConsistencyToken = checkInResp.GetDeviceCheckinConsistencyToken()
   return
}

func (client *_GooglePlayClient) checkIn(req *gpproto.AndroidCheckinRequest) (resp *gpproto.AndroidCheckinResponse, err error) {
   b, err := proto.Marshal(req)
   if err != nil {
      return
   }
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

func (client *_GooglePlayClient) uploadDeviceConfig() (*gpproto.UploadDeviceConfigResponse, error) {
   b, err := proto.Marshal(&gpproto.UploadDeviceConfigRequest{DeviceConfiguration: client._DeviceInfo._GetDeviceConfigProto()})
   if err != nil {
      return nil, err
   }
   r, _ := http.NewRequest("POST", _UrlUploadDeviceConfig, bytes.NewReader(b))
   payload, err := client.doAuthedReq(r)
   if err != nil {
      return nil, err
   }
   return payload.UploadDeviceConfigResponse, nil
}

func (client *_GooglePlayClient) _GenerateGPToken() (string, error) {
   params := &url.Values{}
   client.setDefaultAuthParams(params)
   client.setAuthParams(params)
   params.Set("service", "oauth2:https://www.googleapis.com/auth/googleplay")
   r, _ := http.NewRequest("POST", _UrlAuth+"?"+params.Encode(), nil)
   r.Header.Set("User-Agent", client._DeviceInfo._GetAuthUserAgent())
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

func (client *_GooglePlayClient) setDefaultHeaders(r *http.Request) {
   data := client.AuthData
   r.Header.Set("Authorization", "Bearer "+data._AuthToken)
   r.Header.Set("User-Agent", client._DeviceInfo._GetUserAgent())
   r.Header.Set("X-DFE-Device-Id", data.GsfID)
   // seems like we need this, what the fuck:
   r.Header.Set("X-DFE-Client-Id", "am-android-google")
}

///////////////////////////////////////////////////////////////////////////////

func (client *_GooglePlayClient) zsetDefaultAuthParams(params *url.Values) {
   if client.AuthData.GsfID != "" {
      params.Set("androidId", client.AuthData.GsfID)
   }
   params.Set("sdk_version", strconv.Itoa(int(client._DeviceInfo._Build.GetSdkVersion())))
   params.Set("email", client.AuthData._Email)
   params.Set("google_play_services_version", strconv.Itoa(int(client._DeviceInfo._Build.GetGoogleServices())))
   params.Set("callerSig", "38918a453d07199354f8b19af05ec6562ced5788")
}

func (client *_GooglePlayClient) setAuthParams(params *url.Values) {
   params.Set("app", "com.android.vending")
   params.Set("client_sig", "38918a453d07199354f8b19af05ec6562ced5788")
   params.Set("callerPkg", "com.google.android.gms")
   params.Set("Token", client.AuthData._AASToken)
   params.Set("oauth2_foreground", "1")
   params.Set("token_request_options", "CAA4AVAB")
   params.Set("check_email", "1")
   params.Set("system_partition", "1")
}
