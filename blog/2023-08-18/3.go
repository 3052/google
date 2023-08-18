package play

import (
   "encoding/json"
   "net/http"
   "net/url"
   "strings"
)

func (a account_lookup) signin(password string) (*http.Response, error) {
   body, err := func() (url.Values, error) {
      f_req, err := json.Marshal([]any{
         "", nil, 1, nil, []any{
            1, nil, nil, nil, []any{password},
         },
      })
      if err != nil {
         return nil, err
      }
      v := make(url.Values)
      v.Set("dgresponse", `[null,"/_/signin/challenge",5,0]`)
      v.Set("f.req", string(f_req))
      return v, nil
   }()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://accounts.google.com/_/signin/challenge",
      strings.NewReader(body.Encode()),
   )
   if err != nil {
      return nil, err
   }
   TL, err := a.TL()
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "TL=" + TL
   req.AddCookie(a.host_gaps)
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   req.Header.Set("Google-Accounts-Xsrf", "1")
   return http.DefaultClient.Do(req)
}
