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
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "accounts.google.com"
   req.URL.Path = "/_/signin/challenge"
   req.URL.Scheme = "https"
   req.Header["Google-Accounts-Xsrf"] = []string{"1"}
   req.Header["Cookie"] = []string{
      "__Host-GAPS=1:hqi4p3rL9a-3VlQCy8zCdGcaYKJxEw:JJxXj84I5yVpyCIR",
   }
   val := make(url.Values)
   val["TL"] = []string{"AGEVcSSB93ykkpaW6TycsDaEsjZrKjpjebUf7-59-LaHXk06a5fzHEkmZtlvrma3"}
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
