package gplayapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Juby210/gplayapi-go/gpproto"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

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
		if client._SessionFile != "" {
			client._SaveSession(client._SessionFile)
		}
		res, err = client._doAuthedReq(r)
	}
	return
}

func (client *_GooglePlayClient) _RegenerateGPToken() (err error) {
	client._AuthData._AuthToken, err = client._GenerateGPToken()
	return
}

//goland:noinspection GoUnusedConst
const (
	_UrlBase               = "https://android.clients.google.com"
	_UrlFdfe               = _UrlBase + "/fdfe"
	_UrlAuth               = _UrlBase + "/auth"
	_UrlCheckIn            = _UrlBase + "/checkin"
	_UrlDetails            = _UrlFdfe + "/details"
	_UrlDelivery           = _UrlFdfe + "/delivery"
	_UrlPurchase           = _UrlFdfe + "/purchase"
	_UrlToc                = _UrlFdfe + "/toc"
	_UrlTosAccept          = _UrlFdfe + "/acceptTos"
	_UrlUploadDeviceConfig = _UrlFdfe + "/uploadDeviceConfig"
)

type _GooglePlayClient struct {
	_AuthData   *_AuthData
	_DeviceInfo *_DeviceInfo

	// _SessionFile if _SessionFile is set then session will be saved to it after modification
	_SessionFile string
}

var (
	err_GPTokenExpired = errors.New("unauthorized, gp token expired")

	httpClient = &http.Client{}
)

// _NewClient makes new client with Pixel3a config
func _NewClient(email, aasToken string) (*_GooglePlayClient, error) {
	return _NewClientWithDeviceInfo(email, aasToken, _Pixel3a)
}

func _NewClientWithDeviceInfo(email, aasToken string, deviceInfo *_DeviceInfo) (client *_GooglePlayClient, err error) {
	authData := &_AuthData{
		_Email:    email,
		_AASToken: aasToken,
		_Locale:   "en_GB",
	}
	client = &_GooglePlayClient{_AuthData: authData, _DeviceInfo: deviceInfo}

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

	_, err = client.toc()
	return
}

func (client *_GooglePlayClient) _SaveSession(file string) error {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	return json.NewEncoder(f).Encode(client._AuthData)
}

func _LoadSession(file string) (*_GooglePlayClient, error) {
	return _LoadSessionWithDeviceInfo(file, _Pixel3a)
}

func _LoadSessionWithDeviceInfo(file string, deviceInfo *_DeviceInfo) (client *_GooglePlayClient, err error) {
	f, err := os.Open(file)
	if err != nil {
		return
	}
	client = &_GooglePlayClient{_DeviceInfo: deviceInfo}
	err = json.NewDecoder(f).Decode(&client._AuthData)
	return
}
