package play

import "net/http"

func (e embedded_setup) host_gaps() *http.Cookie {
   for _, cookie := range e {
      if cookie.Name == "__Host-GAPS" {
         return cookie
      }
   }
   return nil
}

type embedded_setup []*http.Cookie

func new_embedded_setup() (embedded_setup, error) {
   res, err := http.Get("https://accounts.google.com/EmbeddedSetup")
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return res.Cookies(), nil
}
