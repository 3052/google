package play

import (
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
)

func Password() {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "accounts.google.com"
   req.URL.Path = "/signin/v2/challenge/password/empty"
   req.URL.Scheme = "https"
   // this comes from the response headers of /_/lookup/accountlookup
   req.Header["Cookie"] = []string{
      "__Host-GAPS=1:pdljXG_kDeyVozXwNv0Jfi1BSVEkRQ:pCfsAWO77eTcMSS5",
   }
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
