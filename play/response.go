package play

import (
   "154.pages.dev/tls"
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
)

func (r Response) Write_File(name string) error {
   data, err := io.ReadAll(r.Body)
   if err != nil {
      return err
   }
   return os.WriteFile(name, data, 0666)
}

type Response struct {
   *http.Response
}

// You can also use host "android.clients.google.com", but it also uses
// TLS fingerprinting.
func New_Auth(email, passwd string) (*Response, error) {
   body := url.Values{
      "Email": {email},
      "Passwd": {passwd},
      "client_sig": {""},
      "droidguard_results": {"-"},
   }.Encode()
   req, err := http.NewRequest(
      "POST", "https://android.googleapis.com/auth", strings.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   client := http.DefaultClient
   client.Transport = &tls.Transport{Spec: tls.Android_API_26}
   res, err := client.Do(req)
   if err != nil {
      return nil, err
   }
   return &Response{res}, nil
}
