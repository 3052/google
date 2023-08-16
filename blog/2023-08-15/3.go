package play

import (
   "encoding/json"
   "net/http"
   "net/http/httputil"
   "net/url"
   "strings"
)

type account_lookup struct {
   host_gaps *http.Cookie
}

func (e embedded_setup) account_lookup() (*account_lookup, error) {
   body, err := func() (url.Values, error) {
      f_req, err := json.MarshalIndent([]any{
         "srpen6",
         e.user_hash(),
         []any{},
         nil,
         "US",
         nil,
         nil,
         2,
         1,
         1,
         []any{},
         "srpen6",
      }, "", " ")
      if err != nil {
         return nil, err
      }
      println(string(f_req))
      v := make(url.Values)
      v.Set("bgRequest", `["identifier",""]`)
      v.Set("f.req", string(f_req))
      return v, nil
   }()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://accounts.google.com/_/lookup/accountlookup",
      strings.NewReader(body.Encode()),
   )
   if err != nil {
      return nil, err
   }
   req.AddCookie(e.host_gaps)
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   req.Header.Set("Google-Accounts-Xsrf", "1")
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   {
      b, err := httputil.DumpResponse(res, true)
      if err != nil {
         return nil, err
      }
      println(string(b))
   }
   var lookup account_lookup
   lookup.host_gaps, err = host_gaps(res)
   if err != nil {
      return nil, err
   }
   return &lookup, nil
}
