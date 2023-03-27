package googleplay

import "2a.pages.dev/rosso/http"

// Purchase app. Only needs to be done once per Google account.
func (h Header) Purchase(c http.Client, doc string) error {
   req := http.Post()
   req.URL.Scheme = "https"
   req.URL.Host = "android.clients.google.com"
   req.URL.Path = "/fdfe/purchase"
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   h.Set_Auth(req.Header)
   h.Set_Device(req.Header)
   req.Body_String("doc=" + doc)
   res, err := c.Do(req)
   if err != nil {
      return err
   }
   return res.Body.Close()
}
