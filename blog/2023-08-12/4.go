package play

import (
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

func four() {
   var req_body = strings.NewReader(url.Values{
      "bgRequest":[]string{`["identifier"]`},
      // this comes from the response body of /EmbeddedSetup:
      "f.req":[]string{`
      [
         "srpen6",
         "AEThLlxalyo1zEtFYE_sbRVZ9BzatfkdwF-82iHmORBNCYFL3UAQdoCkxvUI_inBvwk00CFQtKSDLRQgftu3TvM9PgSqofcQaMXZzPvJcAucQc6C2coHoeTPBKCna_66Rrp15WpExvNbkO7311OCeXuOgusZBJEbkecRTXvhuH-npuY396s-sJ4tIi7syAEyDW0hR-YdbW_t",
         [],
         null,
         "US",
         null,
         null,
         2
      ]
      `},
   }.Encode())
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "accounts.google.com"
   req.URL.Path = "/_/lookup/accountlookup"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(req_body)
   req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded;charset=utf-8"}
   req.Header["Google-Accounts-Xsrf"] = []string{"1"}
   // this comes from the response headers of /_/kids/signup/eligible
   req.Header["Cookie"] = []string{
      "__Host-GAPS=1:OE8e0FiTalfoTnpW78lcyLbF1Z0K-g:LM80V62JRq4tnZbz",
   }
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res_body, err := httputil.DumpResponse(res, false)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(res_body)
}

