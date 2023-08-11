package play

import (
   "154.pages.dev/encoding/json"
   "io"
   "net/http"
)

func get_embedded_info() (string, string, error) {
   res, err := http.Get("https://accounts.google.com/embedded/setup/android")
   if err != nil {
      return "", "", err
   }
   defer res.Body.Close()
   var data struct {
      OewCAd string
   }
   {
      s, err := io.ReadAll(res.Body)
      if err != nil {
         return "", "", err
      }
      sep := json.Split(".WIZ_global_data =")
      if _, err := sep.After(s, &data); err != nil {
         return "", "", err
      }
   }
   var xsrf string
   {
      s := []byte(data.OewCAd)
      _, err := json.Split(`"xsrf",null,[""],`).After(s, &xsrf)
      if err != nil {
         return "", "", err
      }
   }
   return xsrf, "", nil
}
