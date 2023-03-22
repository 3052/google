package googleplay

import (
   "2a.pages.dev/rosso/http"
   "strings"
)

// Purchase app. Only needs to be done once per Google account.
func (h Header) Purchase(app string) error {
   req := http.Post()
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   req.Set_Body(strings.NewReader("doc=" + app))
   req.URL.Host = "android.clients.google.com"
   req.URL.Path = "/fdfe/purchase"
   req.URL.Scheme = "https"
   h.Set_Auth(req.Header)
   h.Set_Device(req.Header)
   res, err := Client.Do(req)
   if err != nil {
      return err
   }
   return res.Body.Close()
}
