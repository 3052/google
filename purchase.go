package googleplay

import (
   "net/http"
   "net/url"
   "strings"
)

// Purchase app. Only needs to be done once per Google account.
func (h Header) Purchase(app string) error {
   body := make(url.Values)
   body.Set("doc", app)
   req, err := http.NewRequest(
      "POST", "https://android.clients.google.com/fdfe/purchase",
      strings.NewReader(body.Encode()),
   )
   if err != nil {
      return err
   }
   h.Set_Auth(req.Header)
   h.Set_Device(req.Header)
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   res, err := Client.Do(req)
   if err != nil {
      return err
   }
   return res.Body.Close()
}
