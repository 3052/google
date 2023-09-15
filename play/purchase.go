package play

import (
   "errors"
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
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   req.Header.Set(h.Authorization())
   req.Header.Set(h.Device())
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   return nil
}

var Native_Platforms = map[string]string{
   // com.google.android.youtube
   "0": "x86",
   // com.miui.weather2
   "1": "armeabi-v7a",
   // com.kakaogames.twodin
   "2": "arm64-v8a",
}
