package main

import (
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Header["Accept-Encoding"] = []string{"gzip"}
   req.Header["Accept-Language"] = []string{"en-US"}
   req.Header["Connection"] = []string{"Keep-Alive"}
   req.Header["Content-Length"] = []string{"11026"}
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Header["Host"] = []string{"play-fe.googleapis.com"}
   req.Header["User-Agent"] = []string{"Android-Finsky/33.2.12-21%20%5B0%5D%20%5BPR%5D%20487986790 (api=3,versionCode=83321210,sdk=24,device=vbox86p,hardware=vbox86,product=vbox86p,platformVersionRelease=7.0,model=Nexus%204,buildId=NBD92Y,isWideScreen=0,supportedAbis=x86)"}
   req.Header["X-Dfe-Build-Fingerprint"] = []string{"google/vbox86p/vbox86p:7.0/NBD92Y/433:userdebug/test-keys"}
   req.Header["X-Dfe-Client-Id"] = []string{"am-unknown"}
   req.Header["X-Dfe-Device-Checkin-Consistency-Token"] = []string{"ABFEt1VnT2UviOAHydPMSmM3NrW_RLHtLIqNrxYrNRPX_5Uhto1QyU3jjlYVOfYBRtjKjh_lJMOST_zB_BmFhykXJwCz7ScpErsPnJ90V6-khkwCS_b58dWE5FGhnqv4CEY4blReAIQi07WOi4hiJRTyg-p2mGsQ1_7jhVOK_5o7HLbalunP81QZZ6FlAWn75pM_VoJ33xSq0mI9GCVJZFhhzkGgtr6klVxVwRYI7A_G8F1OtOywP3DEin0psoPrmoXMv6eAoZ_LUkIyyknDPyDzSZh6QNCOGqUrtjLUv1BJfunZ2oHcQPe0G2zQPLy1pb7RNDsLb9UMs0CSileEk-P1482HCC1_jCJtjJ3yrE4R3y1s7ex66lPJajm5sQ-LgkWQZASHF8DM_vYjlV0odgZYX_UqTGsH870brPBmmBV-Zz1EW1h7cug"}
   req.Header["X-Dfe-Device-Id"] = []string{device}
   req.Header["X-Dfe-Encoded-Targets"] = []string{"CAEaCoqYgQayBZkJ2Qs"}
   req.Header["X-Dfe-Mccmnc"] = []string{"310270"}
   req.Header["X-Dfe-Network-Type"] = []string{"4"}
   req.Header["X-Dfe-Phenotype"] = []string{"H4sIAAAAAAAAABXQO5KbMAAA0EzKlClzh52BCDBbbIEAIVbmzy5YDcNHxgYbm7VARpMiR8_kHeH9-PvdttxsSVzLvHH_CFe59PNg1WzZK9apZEjn0jKsrKmvwaTynUndjoY170TJiDfc0_1nt81syoF9dlwfgPo6ZRp1yZIcH0_kdylurAdw9Kig0Ivs0d88CUa4XPhLKYo42mGRYa7zWSSEHSsDTzD1AyEvYrjkcwvzEN_YmuwlusPCvDhEgndxq0K6BXz8yAyjAtWuVtFat48OroufD0-76xW1MVLRGGGgbpHckuzUZj0yx9accYV1CLzDmWFdvL39_PN_oEic7ayczGtWHqI4zmEIh94o4YemxE8leOWBRr78ilGVnZ-kxabTr1_FC6Gpi0Zald3tOPx2ACdoi5RXKbX7IW1pqileAKocDfb07hKUqDFfGuuzUDbP6HanSf769g8RwBEYgwEAAA"}
   req.Header["X-Dfe-Request-Params"] = []string{"timeoutMs=30000"}
   req.Header["X-Dfe-Userlanguages"] = []string{"en_US"}
   req.Header["X-Ps-Rh"] = []string{"H4sIAAAAAAAAAONqZuRqYPQwKfZ0hIJ4U99KkwpPo5Isx0q_Kv8qXx9fgxIXxwrfqlBj3xDXStconzzXLEcDX3Nfg2IXR6OSkMAK36xQ98r0nIpk0wqzeHdXo1wLZ_ckU6OwoDyv5DS3cgsPxzInxwzf8HhLc2cfzxBnx3In3SAnM7NwsI22tgCrOwMMhgAAAA"}
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "play-fe.googleapis.com"
   req.URL.Path = "/fdfe/autoUpdate"
   req.URL.RawPath = ""
   val := make(url.Values)
   val["nocache_qos"] = []string{"lt"}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(req_body)
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
