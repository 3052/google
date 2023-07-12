package play

import (
   "net/http"
   "strings"
)

// Purchase app. Only needs to be done once per Google account.
func (h Header) Purchase(doc string) error {
   req, err := http.NewRequest(
      "POST", "https://android.clients.google.com/fdfe/purchase",
      strings.NewReader("doc=" + doc),
   )
   if err != nil {
      return err
   }
   h.Set_Auth(req.Header)
   h.Set_Device(req.Header)
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   res, err := client.Do(req)
   if err != nil {
      return err
   }
   return res.Body.Close()
}

var client = http.Client{
   CheckRedirect: func(*http.Request, []*http.Request) error {
      return http.ErrUseLastResponse
   },
}
