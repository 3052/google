package gplayapi

import (
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

func (client *_GooglePlayClient) setDefaultHeaders(r *http.Request) {
   data := client.AuthData
   r.Header.Set("Authorization", "Bearer "+data._AuthToken)
   r.Header.Set("User-Agent", client._DeviceInfo._GetUserAgent())
   r.Header.Set("X-DFE-Device-Id", data.GsfID)
   // seems like we need this, what the fuck:
   r.Header.Set("X-DFE-Client-Id", "am-android-google")
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
type _GooglePlayClient struct {
   AuthData   *_AuthData
   _DeviceInfo *_DeviceInfo
}

func NewClientWithDeviceInfo(email, aasToken string, deviceInfo *_DeviceInfo) (client *_GooglePlayClient, err error) {
   authData := &_AuthData{
      _Email:    email,
      _AASToken: aasToken,
      _Locale:   "en_GB",
   }
   client = &_GooglePlayClient{AuthData: authData, _DeviceInfo: deviceInfo}

   _, err = client._GenerateGsfID()
   if err != nil {
      return
   }
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

func (client *_GooglePlayClient) _doAuthedReq(r *http.Request) (_ *gpproto.Payload, err error) {
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

func (client *_GooglePlayClient) doAuthedReq(r *http.Request) (res *gpproto.Payload, err error) {
   res, err = client._doAuthedReq(r)
   if err == err_GPTokenExpired {
      err = client._RegenerateGPToken()
      if err != nil {
         return
      }
      res, err = client._doAuthedReq(r)
   }
   return
}

func (client *_GooglePlayClient) _RegenerateGPToken() (err error) {
   client.AuthData._AuthToken, err = client._GenerateGPToken()
   return
}

const _UrlBase               = "https://android.clients.google.com"
const _UrlFdfe               = _UrlBase + "/fdfe"
const _UrlAuth               = _UrlBase + "/auth"
const _UrlCheckIn            = _UrlBase + "/checkin"
const _UrlUploadDeviceConfig = _UrlFdfe + "/uploadDeviceConfig"

var (
   err_GPTokenExpired = errors.New("unauthorized, gp token expired")

   httpClient = &http.Client{}
)

