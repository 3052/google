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
   req.Header["Authorization"] = []string{"Bearer ya29.a0AfB_byAoIBf1zus62HhiNeCeeYzfjMoCQrePnesbbnHIE41u1IRfZji2NpjwiUKKa7mUvH473sS5ykvBSAGItOK2lqPTenL5CCimJrE7Se-ZublD7J0ij_O7je5ENbAMj6RWonqKeKv9gchIMCKOsDoRYd-htrKEjJdwadqEKzbhWaZuD3eoehbts6ltNxs4pbI7YT4KQDIRMa-2teUL2jEd8CxNkXi3jMd4YUdUsnCreuwg2nHyuig6Rdgi6eKpDHpvSK9N3HpUP8yEDNFZeiRC58raFlLwvj8rM0CmyUltRHxdKQkyZ39kBjvajOFH70J7-gaCgYKARISARESFQGOcNnC9DFuvWMX4eHk3QpVJ7AQfQ0333"}
   req.Header["User-Agent"] = []string{"Android-Finsky/14.9.76-all%20%5B0%5D%20%5BPR%5D%20248813671 (api=3,versionCode=81497600,sdk=26,device=generic_x86,hardware=ranchu,product=sdk_gphone_x86,platformVersionRelease=8.0.0,model=Android%20SDK%20built%20for%20x86,buildId=MASTER,isWideScreen=0,supportedAbis=x86)"}
   req.Header["X-Dfe-Device-Id"] = []string{"39f20681ab4baadb"}
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "play-fe.googleapis.com"
   req.URL.Path = "/fdfe/delivery"
   val := make(url.Values)
   val["da"] = []string{"3"}
   val["doc"] = []string{"com.duolingo"}
   val["fdcf"] = []string{"1", "2"}
   val["ia"] = []string{"false"}
   val["isid"] = []string{"dummysessionid"}
   val["ot"] = []string{"1"}
   val["st"] = []string{"EJSTnagGGWcqNmXSQdlB"}
   val["vc"] = []string{"1700"}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
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
