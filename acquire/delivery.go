package acquire

import (
   "154.pages.dev/google/play"
   "154.pages.dev/protobuf"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strconv"
)

func Delivery(h play.Header, doc string, version uint64) error {
   var req http.Request
   req.Header = make(http.Header)
   req.Header["Accept-Encoding"] = []string{"gzip"}
   req.Header["Accept-Language"] = []string{"en-US"}
   //req.Header["Authorization"] = []string{"Bearer ya29.a0AfB_byBFbblkPw8axrrH36QfGChD59B98T9SvEHUyT9qlXvp3_9bo2eo8o1xSNorKLE2Bow0nDFeMeUTkBNZH2e52N1ptv-2gCb0siAmzhONjHiBe7XvmkrszKepFqLHlE8NJuuSEQcGcpoxlZAh4x0WjhsgjqBwlV1svV3-CqjPm44kwH_yiWGWenAcHfM1KY9wgVXYst-VopZ1VZqYz_wfggktAZrjPSardJZ5t8lUoNr3DFZGNajTo_dVZ_ydlzr3pUMR1HGYfbIRxKtGpXKP29lKBBj-G_0hiSIOP0g8_21UHrjH7kGwYE84iDPNB4sRqAaCgYKAXUSARESFQGOcNnC64WSRP9UZTgXivBtLHHlpg0333"}
   req.Header.Set(h.Authorization())
   req.Header["Connection"] = []string{"Keep-Alive"}
   req.Header["Content-Length"] = []string{"0"}
   req.Header["Host"] = []string{"android.clients.google.com"}
   req.Header["User-Agent"] = []string{"Android-Finsky/22.8.44-21%20%5B0%5D%20%5BPR%5D%20342964500 (api=3,versionCode=82284410,sdk=27,device=generic_x86,hardware=ranchu,product=sdk_gphone_x86,platformVersionRelease=8.1.0,model=Android%20SDK%20built%20for%20x86,buildId=OSM1.180201.037,isWideScreen=0,supportedAbis=x86)"}
   req.Header["X-Dfe-Client-Id"] = []string{"am-unknown"}
   req.Header["X-Dfe-Content-Filters"] = []string{""}
   req.Header["X-Dfe-Cookie"] = []string{"EAEYACICVVMyUENqa2FOd29UTXprME5ETXdOekUzTmpjME9UVTVNemd6TlJJZ0NoQXhOamsxTVRZeU9UZ3hOVGt5T1RNMUVnd0k1Y1NvcUFZUTJQRGRtZ0k9ShIKAlVTEgwI5Zi8qQYQsNLC5AJYAA"}
   req.Header["X-Dfe-Device-Config-Token"] = []string{"CjkaNwoTMzk0NDMwNzE3Njc0OTU5MzgzNRIgChAxNjk1MTYyOTgxNTkyOTM1EgwI5cSoqAYQ2PDdmgI="}
   req.Header["X-Dfe-Device-Id"] = []string{"36bcfe7abb2ef4eb"}
   req.Header["X-Dfe-Encoded-Targets"] = []string{"CAESD1elooEG2AOB4QGmNdE3ARp/4wUF0I+BBgnEAgQBAQ3nAX0vugFWASEKIw0iN4QBDgxjCVi9AYwCFExaARRbGeIFBJICPJwCpwekBFm9BpIK4AXgH64c7gbkB5saAgGQBaUHpQLtCvcHqwLjCYMDvALLARakP58/gwZF5QXUAq4R6RqIDvYL9wKbA8c7wcaQEg"}
   req.Header["X-Dfe-Mccmnc"] = []string{"310260"}
   req.Header["X-Dfe-Network-Type"] = []string{"4"}
   req.Header["X-Dfe-Request-Params"] = []string{"timeoutMs=4000"}
   req.Header["X-Dfe-Userlanguages"] = []string{"en_US"}
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "android.clients.google.com"
   req.URL.Path = "/fdfe/delivery"
   val := make(url.Values)
   val["da"] = []string{"3"}
   val["fdcf"] = []string{"1", "2"}
   val["ia"] = []string{"false"}
   val["isid"] = []string{"sOSbXrwaQ-WUNeCgghpZ3w"}
   val["ot"] = []string{"1"}
   val["st"] = []string{"EPPEqKgGGWgF45yIQtlB"}
   req.Header.Set(h.Authorization())
   val["doc"] = []string{doc}
   val["vc"] = []string{strconv.FormatUint(version, 10)}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   data, err := io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   mes, err := protobuf.Consume(data) // ResponseWrapper
   if err != nil {
      return err
   }
   mes, ok := mes.Message(1)
   if !ok {
      return errors.New("payload")
   }
   mes, ok = mes.Message(21)
   if !ok {
      return errors.New("deliveryResponse")
   }
   status, ok := mes.Varint(1)
   if ok {
      switch status {
      case 3:
         return errors.New("acquire")
      case 5:
         return errors.New("version")
      }
   }
   _, ok = mes.Message(2)
   if !ok {
      return errors.New("appDeliveryData")
   }
   return nil
}
