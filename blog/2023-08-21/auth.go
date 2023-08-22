package main

import (
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

func main() {
   req := new(http.Request)
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "android.googleapis.com"
   req.URL.Path = "/auth"
   req.URL.Scheme = "https"
   req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
   req_body := url.Values{
      "ACCESS_TOKEN":[]string{"1"},
      "Token":[]string{"oauth2_4/0Adeu5BWqSPYQPO4tueE1_leRDb_7Lk9CW8K6gIah31lcxPAns-Kl9-rkWPCJwTbxOgVOZg"},
      "service":[]string{"ac2dm"},
   }.Encode()
   req.Body = io.NopCloser(strings.NewReader(req_body))
   res, err := new(http.Transport).RoundTrip(req)
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
