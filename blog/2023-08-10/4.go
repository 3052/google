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
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "accounts.google.com"
   req.URL.Path = "/_/signin/speedbump/embeddedsigninconsent"
   req.URL.Scheme = "https"
   req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded;charset=utf-8"}
   val := make(url.Values)
   //////////////////////////////////////////////////////////////////////////////
   req.Header["Cookie"] = []string{"__Host-GAPS=1:JMj67XI_OL4priAm3iOEO1AtIzBLmw:xUE_yg3eof0TLuyp"}
   val["TL"] = []string{"AGEVcSQw-DNHaJ1fG5-Y6N8PuMaKdeMpVe36fSfAvnd15RZKlvDYUAg8wqYmrvvb"}
   val["ardt"] = []string{"AFWLbD0xC8RxdHCGv8sPzRm4Jx_4WrKSYGYp28h8tsVQUsF19UgUyRIpAAqSfpZUS7UaKudWqoJQy_cDLmgk8m5xQzIeUWGxKMg6dFQzPtSNq_BRwN1PQ8pv68mT"}
   req_body := url.Values{
      "f.req": {`
      [
         "gf.siesic", 1, "US", "en", "default", ["XFFfAjryY6A="]
      ]
      `},
   }.Encode()
   //////////////////////////////////////////////////////////////////////////////
   req.Body = io.NopCloser(strings.NewReader(req_body))
   req.URL.RawQuery = val.Encode()
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
