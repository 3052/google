package play

import (
   "errors"
   "net/http"
   "strings"
)

type eligible []*http.Cookie

func (e embedded_setup) eligible() (eligible, error) {
   req, err := http.NewRequest(
      "POST", "https://accounts.google.com/_/kids/signup/eligible",
      strings.NewReader("f.req=[]"),
   )
   if err != nil {
      return nil, err
   }
   req.AddCookie(host_gaps(e.cookies))
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   return res.Cookies(), nil
}
