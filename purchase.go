package googleplay

import "2a.pages.dev/rosso/http"

// Purchase app. Only needs to be done once per Google account.
func (h Header) Purchase(doc string) error {
   req := http.Post()
   req.Body_String("doc=" + doc)
   req.URL.Scheme = "https"
   req.URL.Host = "android.clients.google.com"
   req.URL.Path = "/fdfe/purchase"
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   h.Set_Auth(req.Header)
   h.Set_Device(req.Header)
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return err
   }
   return res.Body.Close()
}
