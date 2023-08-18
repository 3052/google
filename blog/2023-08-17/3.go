package play

import "net/http"

func (a account_lookup) signin() (*http.Response, error) {
   req, err := http.NewRequest(
      "POST", "https://accounts.google.com/_/signin/challenge", nil,
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
   req.Header.Set("Google-Accounts-Xsrf", "1")
   return http.DefaultClient.Do(req)
}
