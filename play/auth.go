package play

import (
   "154.pages.dev/tls"
   "io"
   "net/http"
   "net/url"
)

// You can also use host "android.clients.google.com", but it also uses
// TLS fingerprinting.
func New_Auth(email, passwd string) (Auth, error) {
   client := *http.DefaultClient
   client.Transport = &tls.Transport{Spec: tls.Android_API_26}
   res, err := client.PostForm(
      "https://android.googleapis.com/auth",
      url.Values{
         "Email": {email},
         "Passwd": {passwd},
         "client_sig": {""},
         "droidguard_results": {"-"},
      },
   )
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   text, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   a := make(Auth)
   if err := a.UnmarshalText(text); err != nil {
      return nil, err
   }
   return a, nil
}

func (a Auth) Auth() string {
   return a["Auth"]
}

func (a Auth) Exchange() error {
   // these values take from Android API 28
   res, err := http.PostForm(
      "https://android.googleapis.com/auth",
      url.Values{
         "Token": {a.Token()},
         "app": {"com.android.vending"},
         "client_sig": {"38918a453d07199354f8b19af05ec6562ced5788"},
         "service": {"oauth2:https://www.googleapis.com/auth/googleplay"},
      },
   )
   if err != nil {
      return err
   }
   defer res.Body.Close()
   text, err := io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   return a.UnmarshalText(text)
}

func (a Auth) Token() string {
   return a["Token"]
}
