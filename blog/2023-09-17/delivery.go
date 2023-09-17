package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	var req http.Request
	req.Header = make(http.Header)
	req.Header["Authorization"] = []string{"Bearer ya29.a0AfB_byAXBBcOb0E6BfWrnUzWf1rR7aJGW_jElONLMcws3JhdOrPAX5Zt8wzkD_fcodkBwxeofxtK5htEp4xOX10hf_MQiPN9_AagkzndyraqPp39XY-qDMzqaTi-mR4ZIEGvaYaLRo27cZTKg8MpenWG2TxN12igV3VVOFtOHK-Igf-SX4XHiF2-tyF62Zg6hQoZKBelr_kERd4haLboTznh67syvGB6ElnH8j4QwQZdTyxWNycrC9qaSJz14LPWYW7XwhfFS08JL2i57JaiFdflXd4eIM69UMjABgPWn_DGAnpg-dKVMZmJUAvHBNZ5A8Bc3waCgYKAaISARESFQGOcNnCuty79iAXfjU2kIxMjHf-nw0333"}
	req.Header["Content-Length"] = []string{"0"}
	req.Header["Host"] = []string{"play-fe.googleapis.com"}
	req.Header["User-Agent"] = []string{"Android-Finsky/22.8.44-21%20%5B0%5D%20%5BPR%5D%20342964500 (api=3,versionCode=82284410,sdk=27,device=generic_x86,hardware=ranchu,product=sdk_gphone_x86,platformVersionRelease=8.1.0,model=Android%20SDK%20built%20for%20x86,buildId=OSM1.180201.037,isWideScreen=0,supportedAbis=x86)"}
	req.Header["X-Dfe-Device-Id"] = []string{"344d67278408e17a"}
	req.Method = "GET"
	req.ProtoMajor = 1
	req.ProtoMinor = 1
	req.URL = new(url.URL)
	req.URL.Host = "play-fe.googleapis.com"
	req.URL.Path = "/fdfe/delivery"
	req.URL.RawPath = ""
	val := make(url.Values)
	val["da"] = []string{"3"}
	val["doc"] = []string{"com.duolingo"}
	val["fdcf"] = []string{"1", "2"}
	val["ia"] = []string{"false"}
	val["isid"] = []string{"oLHwp38pRtqEky5tTkuimg"}
	val["ot"] = []string{"1"}
	val["st"] = []string{"EMbEnagGGdo8npHYQdlB"}
	val["vc"] = []string{"1700"}
	req.URL.RawQuery = val.Encode()
	req.URL.Scheme = "https"
	req.Body = nil
	res, err := new(http.Transport).RoundTrip(&req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	res_body, err := httputil.DumpResponse(res, true)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(res_body)
}
