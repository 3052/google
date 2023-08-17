package play

import (
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

func Consent() {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "accounts.google.com"
   req.URL.Path = "/_/signin/speedbump/embeddedsigninconsent"
   req.URL.Scheme = "https"
   val := make(url.Values)
   req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded;charset=utf-8"}
   // this comes from the response headers of:
   // /_/signin/challenge
   // /signin/v2/challenge/password/empty
   // this breaks the login:
   // /~u _.signin.challenge/200
   // we can do this:
   // /~u signin.v2.challenge.password.empty/204
   req.Header["Cookie"] = []string{"__Host-GAPS=1:JMj67XI_OL4priAm3iOEO1AtIzBLmw:xUE_yg3eof0TLuyp"}
   // this comes from the response body of:
   // /_/signin/challenge
   // /_/lookup/accountlookup
   // these break the login:
   // /~u _.signin.challenge/200
   // /~u _.lookup.accountlookup/200
   val["TL"] = []string{"AGEVcSQw-DNHaJ1fG5-Y6N8PuMaKdeMpVe36fSfAvnd15RZKlvDYUAg8wqYmrvvb"}
   // this comes from the response body of /EmbeddedSetup:
   val["ardt"] = []string{"AFWLbD0xC8RxdHCGv8sPzRm4Jx_4WrKSYGYp28h8tsVQUsF19UgUyRIpAAqSfpZUS7UaKudWqoJQy_cDLmgk8m5xQzIeUWGxKMg6dFQzPtSNq_BRwN1PQ8pv68mT"}
   // this comes from response body of:
   // ~u accounts.static._.js.+signinconsent
   // lasts for at least 4 years. this breaks the login:
   // /~u accounts.static._.js.+signinconsent/200
   req_body := url.Values{
      "f.req": {`
      [
         "gf.siesic", 1, "US", "en", "default", ["XFFfAjryY6A="]
      ]
      `},
   }.Encode()
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
