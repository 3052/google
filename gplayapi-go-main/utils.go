package gplayapi

import (
	"io"
	"net/http"
	"strings"

	"github.com/Juby210/gplayapi-go/gpproto"
	"google.golang.org/protobuf/proto"
)

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
		return nil, GPTokenExpired
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
	if err == GPTokenExpired {
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
