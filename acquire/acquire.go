package acquire

import (
   "154.pages.dev/google/play"
   "154.pages.dev/protobuf"
   "bytes"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strings"
)

func Acquire(h play.Header, doc string, version uint64) error {
   body := protobuf.Message{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Bytes("com.duolingo")},
            protobuf.Field{Number: 2, Type: 0,  Value: protobuf.Varint(1)},
            protobuf.Field{Number: 3, Type: 0,  Value: protobuf.Varint(3)},
         }},
         protobuf.Field{Number: 2, Type: 0,  Value: protobuf.Varint(1)},
         protobuf.Field{Number: 3, Type: 2,  Value: protobuf.Bytes("")},
      }},
      protobuf.Field{Number: 4, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 0,  Value: protobuf.Varint(10)},
      }},
      protobuf.Field{Number: 8, Type: 2,  Value: protobuf.Bytes("")},
      protobuf.Field{Number: 12, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 0,  Value: protobuf.Varint(1700)},
         protobuf.Field{Number: 3, Type: 0,  Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 13, Type: 0,  Value: protobuf.Varint(1)},
      protobuf.Field{Number: 15, Type: 0,  Value: protobuf.Varint(0)},
      protobuf.Field{Number: 22, Type: 2,  Value: protobuf.Bytes("nonce=LHEvHQSIv40zqfHYASj1rKZuAlwtyacUe2pe1tNd9mNv0suSHagkp-88-WaZC48Ev8JTtj9AsH_w6waS7IMFUlrOtxQ70udW-oUjACf7qMlgGnoMLTS2h5pOb_n327LSWmX2lnOoSIMjw6Z1bi5DkPGPrxzPDCnJQfxjke_XvsScFbFqoAdDegtrxKViI0r6dOC3KtcaavBPnNBOZa8DH6RY8Dg83QxzAgIvm7zyUSgrAJIBPErFoR3uDnnb_6nKPV5w-PFat1BgOupbCWAoaO0HTmHaAuE51PcafSGP0DC-D_dlDiJrwbHq8J2F8-3MeHWUwxHmJ1NVhTXYM_njvQ")},
      protobuf.Field{Number: 25, Type: 0,  Value: protobuf.Varint(2)},
   }
   var req http.Request
   req.Header = make(http.Header)
   req.Header["Accept-Encoding"] = []string{"gzip"}
   req.Header["Accept-Language"] = []string{"en-US"}
   //req.Header["Authorization"] = []string{"Bearer ya29.a0AfB_byBFbblkPw8axrrH36QfGChD59B98T9SvEHUyT9qlXvp3_9bo2eo8o1xSNorKLE2Bow0nDFeMeUTkBNZH2e52N1ptv-2gCb0siAmzhONjHiBe7XvmkrszKepFqLHlE8NJuuSEQcGcpoxlZAh4x0WjhsgjqBwlV1svV3-CqjPm44kwH_yiWGWenAcHfM1KY9wgVXYst-VopZ1VZqYz_wfggktAZrjPSardJZ5t8lUoNr3DFZGNajTo_dVZ_ydlzr3pUMR1HGYfbIRxKtGpXKP29lKBBj-G_0hiSIOP0g8_21UHrjH7kGwYE84iDPNB4sRqAaCgYKAXUSARESFQGOcNnC64WSRP9UZTgXivBtLHHlpg0333"}
   req.Header["Connection"] = []string{"Keep-Alive"}
   req.Header["Content-Length"] = []string{"398"}
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
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
   req.Header["X-Dfe-Request-Params"] = []string{"timeoutMs=35000"}
   req.Header["X-Public-Android-Id"] = []string{"c347e6c1b1c92e4f"}
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
   req.Body = io.NopCloser(bytes.NewReader(body.Append(nil)))
   req.Header.Set(h.Authorization())
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   {
      b, err := io.ReadAll(res.Body)
      if err != nil {
         return err
      }
      s := string(b)
      for _, err := range acquire_errors {
         if strings.Contains(s, err) {
            return new_error(err)
         }
      }
   }
   return nil
}

var acquire_errors = []string{
   "Please open my apps to establish a connection with the server.",
   "Error",
}

func new_error(s string) error {
   s = strings.TrimSuffix(s, ".")
   return errors.New(strings.ToLower(s))
}
