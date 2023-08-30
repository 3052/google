package acquire

import (
	"154.pages.dev/encoding/protobuf"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

var acquire_2_body = protobuf.Message{
	protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\n*\n$homeworkout.homeworkouts.noequipment\x10\x01\x18\x03\x10\x01\x1a\x008\x01")},
	protobuf.Field{Number: 1, Type: 2, Value: protobuf.Maybe{
		protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\n$homeworkout.homeworkouts.noequipment\x10\x01\x18\x03")},
		protobuf.Field{Number: 1, Type: 2, Value: protobuf.Maybe{
			protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("homeworkout.homeworkouts.noequipment")},
			protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1)},
			protobuf.Field{Number: 3, Type: 0, Value: protobuf.Varint(3)},
		}},
		protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1)},
		protobuf.Field{Number: 3, Type: 2, Value: protobuf.Bytes("")},
		protobuf.Field{Number: 7, Type: 0, Value: protobuf.Varint(1)},
	}},
	protobuf.Field{Number: 4, Type: 2, Value: protobuf.Bytes("\b\n")},
	protobuf.Field{Number: 4, Type: 2, Value: protobuf.Maybe{
		protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(10)},
	}},
	protobuf.Field{Number: 8, Type: 2, Value: protobuf.Bytes("")},
	protobuf.Field{Number: 12, Type: 2, Value: protobuf.Bytes("\bI")},
	protobuf.Field{Number: 12, Type: 2, Value: protobuf.Maybe{
		protobuf.Field{Number: 1, Type: 0, Value: protobuf.Varint(73)},
	}},
	protobuf.Field{Number: 13, Type: 0, Value: protobuf.Varint(1)},
	protobuf.Field{Number: 22, Type: 2, Value: protobuf.Bytes("nonce=zPXXYPxcgwrtB1eIsiqpQrEfdQmj66MyODvudKhhNu_gY3xTVjqdG76ra9Gw3oXjc3xLjg679AwF-wDQQTwK5P71n-0Ibt7EslTYUZNb8MlSMAiPZHy5CUXlP6YsZO65l-2CfgtDINF7ND77Ym6oBL4lpsiKp7KfbGuLXyru0j6roligDJnak6f-zQ7UuNLFa_lQqyezDlLevV2HILD73VUWtDvmNDVR7lqxgsFuZevOMywJ3zVikTXkcYcMeqkthPa5QA9gErMALdpoE1tdnyoHrba-63v2r7lQNY6biCO7TwEsGRR9jZjkfCTEOcqZ7laxXbFa3U1v1w-u03P9vw")},
	protobuf.Field{Number: 25, Type: 0, Value: protobuf.Varint(2)},
	protobuf.Field{Number: 28, Type: 0, Value: protobuf.Varint(1)},
}

func acquire_2() error {
	var req_body = strings.NewReader("\n2\n*\n$homeworkout.homeworkouts.noequipment\x10\x01\x18\x03\x10\x01\x1a\x008\x01\"\x02\b\nB\x00b\x02\bIh\x01\xb2\x01\xdc\x02nonce=zPXXYPxcgwrtB1eIsiqpQrEfdQmj66MyODvudKhhNu_gY3xTVjqdG76ra9Gw3oXjc3xLjg679AwF-wDQQTwK5P71n-0Ibt7EslTYUZNb8MlSMAiPZHy5CUXlP6YsZO65l-2CfgtDINF7ND77Ym6oBL4lpsiKp7KfbGuLXyru0j6roligDJnak6f-zQ7UuNLFa_lQqyezDlLevV2HILD73VUWtDvmNDVR7lqxgsFuZevOMywJ3zVikTXkcYcMeqkthPa5QA9gErMALdpoE1tdnyoHrba-63v2r7lQNY6biCO7TwEsGRR9jZjkfCTEOcqZ7laxXbFa3U1v1w-u03P9vw\xc8\x01\x02\xe0\x01\x01")
	var req http.Request
	req.Header = make(http.Header)
	req.Header["Accept-Encoding"] = []string{"identity"}
	req.Header["Accept-Language"] = []string{"en-US"}
	req.Header["Authorization"] = []string{"Bearer ya29.a0AfB_byACUOQrHlPrWZ4lLFKaWG8DfVnkOKrugh3WH0JD318YLiAkCqQTyJpi_ogWkRPIVaYjNSOWAIJEVLCQ-Q49F0tDtG3660Dw3pwRKNPPb4MIZb0PW_-iYAtj0krgBZb6eXNPkVBWbeMqC4AtejgwolCtglYRW9fSbNIdQ7qk_s2xuHduc4KakntzYx9UHPI8s3EwusVa7j_9Ws6CWxLAXTFIYSRMX2Syty6RIxxnTPl9xvXyt5mz9FNI0Kd4F7aj99Y9_1537ttQnTt8VS9Hzgfj4OQ-7NZd8qsJ8CK2xvRhEeRx9ReqhCjuAMN9fgT7uTklpwaCgYKATISARESFQHsvYlsdOI85Jec-NMpT1S4AmxojQ0337"}
	req.Header["Connection"] = []string{"Keep-Alive"}
	req.Header["Content-Type"] = []string{"application/x-protobuf"}
	req.Header["Host"] = []string{"android.clients.google.com"}
	req.Header["User-Agent"] = []string{"Android-Finsky/20.1.18-all%20%5B0%5D%20%5BPR%5D%20311592326 (api=3,versionCode=82011800,sdk=21,device=generic_x86,hardware=ranchu,product=sdk_google_phone_x86,platformVersionRelease=5.0.2,model=Android%20SDK%20built%20for%20x86,buildId=LSY66K,isWideScreen=0,supportedAbis=x86)"}
	req.Header["X-Dfe-Client-Id"] = []string{"am-google"}
	req.Header["X-Dfe-Content-Filters"] = []string{""}
	req.Header["X-Dfe-Cookie"] = []string{"EAEYACICVVMyUENqa2FOd29UTkRBMk5EZ3dNemt5TVRBeE5qWTVOakV6T0JJZ0NoQXhOamt6TXpVNU9UVTJOekkyTURZNEVnd0kxTDY2cHdZUW9OYWIyZ0k9ShIKAlVTEgwI1JLOqAYQqIbkvgNYAA"}
	req.Header["X-Dfe-Device-Config-Token"] = []string{"CjkaNwoTNDA2NDgwMzkyMTAxNjY5NjEzOBIgChAxNjkzMzU5OTU2NzI2MDY4EgwI1L66pwYQoNab2gI="}
	req.Header["X-Dfe-Device-Id"] = []string{"386915a141d0ad4a"}
	req.Header["X-Dfe-Encoded-Targets"] = []string{"CAESIFelooEG2AMBARShgQHKX6Y1szCeBwGOBcUMrjDtDtkPGnTjBQXQj4EGCcQCBAEBDecBfS+6AVYBIQojDSI3hAEODGMJWL0BjAIUTMoBGeIFBJICPJwCpwekBFm9BpIK4AXgH64c7gbkB5saAgGQBaUHpQLtCvcHqwLjCYMDvALLARakP6JFRbkIrhHpGv4ZkgbHO+A8EQ"}
	req.Header["X-Dfe-Mccmnc"] = []string{"310260"}
	req.Header["X-Dfe-Network-Type"] = []string{"3"}
	req.Header["X-Dfe-Request-Params"] = []string{"timeoutMs=35000"}
	req.Header["X-Public-Android-Id"] = []string{"7cfdaa6592b8e5e4"}
	req.Method = "POST"
	req.ProtoMajor = 1
	req.ProtoMinor = 1
	req.URL = new(url.URL)
	req.URL.Host = "android.clients.google.com"
	req.URL.Path = "/fdfe/acquire"
	req.URL.RawPath = ""
	val := make(url.Values)
	val["theme"] = []string{"2"}
	req.URL.RawQuery = val.Encode()
	req.URL.Scheme = "https"
	req.Body = io.NopCloser(req_body)
	res, err := new(http.Transport).RoundTrip(&req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	res_body, err := httputil.DumpResponse(res, true)
	if err != nil {
		return err
	}
	fmt.Printf("%q\n", res_body)
	return nil
}
