package play

import (
   "encoding/json"
   "net/http"
   "net/url"
   "strings"
)

func (es embedded_setup) account_lookup(e eligible) (*http.Response, error) {
   body, err := func() (url.Values, error) {
      f_req, err := json.Marshal([]any{
         "srpen6",
         es.user_hash(),
         []any{},
         nil,
         "US",
         nil,
         nil,
         2,
         false,
         true,
      })
      if err != nil {
         return nil, err
      }
      v := make(url.Values)
      v.Set("bgRequest", `["identifier"]`)
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
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   req.Header.Set("Google-Accounts-Xsrf", "1")
   {
      c, err := host_gaps(e)
      if err != nil {
         return nil, err
      }
      req.AddCookie(c)
   }
   return http.DefaultClient.Do(req)
}
