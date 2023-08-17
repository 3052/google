package play

import (
   "net/http"
   "net/url"
)

func signin() (*http.Response, error) {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "accounts.google.com"
   req.URL.Path = "/_/signin/challenge"
   req.URL.Scheme = "https"
   req.Header["Google-Accounts-Xsrf"] = []string{"1"}
   val := make(url.Values)
   // this comes from the response headers of:
   // /_/lookup/accountlookup
   req.Header["Cookie"] = []string{
      "__Host-GAPS=1:hqi4p3rL9a-3VlQCy8zCdGcaYKJxEw:JJxXj84I5yVpyCIR",
   }
   // this comes from the response body of:
   // /_/lookup/accountlookup
   val["TL"] = []string{"AGEVcSSB93ykkpaW6TycsDaEsjZrKjpjebUf7-59-LaHXk06a5fzHEkmZtlvrma3"}
   req.URL.RawQuery = val.Encode()
   return http.DefaultClient.Do(&req)
}
