package play

import (
   "io"
   "net/http"
   "strings"
)

// the response for this is good for at least 4 years
func static_JS() ([]byte, error) {
   req, err := http.NewRequest("GET", "https://ssl.gstatic.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/accounts/static/_/js/" + strings.Join([]string{
      // this comes from response body of /EmbeddedSetup:
      "k=gaia.gaiafe_glif.en.DL5rcs6IWg4.O",
      // this comes from response body of /EmbeddedSetup:
      "am=AgAABvAJEsEf_H8NyEEAAAAAAAAgAAAQAjXiXhoqSAE",
      // this comes from response body of /EmbeddedSetup:
      "rs=ABkqax1GxYdmsoXxFit_qF1awcl_spmjcg",
   }, "/")
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return io.ReadAll(res.Body)
}
