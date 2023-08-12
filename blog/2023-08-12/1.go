package play

import (
   "154.pages.dev/encoding/xml"
   "io"
   "net/http"
)

func new_embedded_setup() (*embedded_setup, error) {
   res, err := http.Get("https://accounts.google.com/EmbeddedSetup")
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var e embedded_setup
   e.cookies = res.Cookies()
   text, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   _, after := xml.Cut(text, []byte("</div>"), []byte(`<div id="view_container"`))
   if err := xml.Unmarshal(after, &e.view_container); err != nil {
      return nil, err
   }
   return &e, nil
}

type embedded_setup struct {
   cookies []*http.Cookie
   view_container struct {
      Initial_Setup_Data string `xml:"data-initial-setup-data,attr"`
   } 
}

func (e embedded_setup) host_gaps() *http.Cookie {
   for _, cookie := range e.cookies {
      if cookie.Name == "__Host-GAPS" {
         return cookie
      }
   }
   return nil
}
