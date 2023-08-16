package play

import "net/http"

func (a account_lookup) password() (*http.Response, error) {
   req, err := http.NewRequest(
      "POST",
      "https://accounts.google.com/signin/v2/challenge/password/empty",
      nil,
   )
   if err != nil {
      return nil, err
   }
   req.AddCookie(a.host_gaps)
   return http.DefaultClient.Do(req)
}
