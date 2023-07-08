package google_play

import (
   "net/http"
   "net/url"
)

// Purchase app. Only needs to be done once per Google account.
func (h Header) Purchase(doc string) error {
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "android.clients.google.com",
      Path: "/fdfe/purchase",
   })
   req.Body_String("doc=" + doc)
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   h.Set_Auth(req.Header)
   h.Set_Device(req.Header)
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return err
   }
   return res.Body.Close()
}
