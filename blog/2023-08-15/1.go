package play

import (
   "154.pages.dev/encoding/xml"
   "encoding/json"
   "io"
   "net/http"
   "strings"
)

func host_gaps(cookies []*http.Cookie) (*http.Cookie, error) {
   for _, cookie := range cookies {
      if cookie.Name == "__Host-GAPS" {
         return cookie, nil
      }
   }
   return nil, http.ErrNoCookie
}

type embedded_setup struct {
   // this is needed for /_/kids/signup/eligible:
   cookies []*http.Cookie
   // this is needed for /_/lookup/accountlookup:
   initial_setup_data []any
}

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
   _, text = xml.Cut(text, []byte("</div>"), []byte(`<div id="view_container"`))
   var view struct {
      Initial_Setup_Data string `xml:"data-initial-setup-data,attr"`
   } 
   if err := xml.Unmarshal(text, &view); err != nil {
      return nil, err
   }
   data := strings.Replace(view.Initial_Setup_Data, "%.@.", "[", 1)
   if err := json.Unmarshal([]byte(data), &e.initial_setup_data); err != nil {
      return nil, err
   }
   return &e, nil
}

func (e embedded_setup) user_hash() string {
   return e.initial_setup_data[13].(string)
}
