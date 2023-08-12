package play

import (
   "net/http"
   "strings"
)

func (e embedded_setup) eligible() (*http.Response, error) {
   req, err := http.NewRequest(
      "POST", "https://accounts.google.com/_/kids/signup/eligible",
      strings.NewReader("f.req=[]"),
   )
   if err != nil {
      return nil, err
   }
   req.AddCookie(e.host_gaps())
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   return http.DefaultClient.Do(req)
}
