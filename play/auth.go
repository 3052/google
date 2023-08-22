package play

import (
   "io"
   "net/http"
   "net/url"
)

// accounts.google.com/embedded/setup/android
// the authorization code (oauth_token) looks like this:
// 4/0Adeu5B...
// but it should be supplied here with the prefix:
// oauth2_4/0Adeu5B...
func New_Auth(code string) (Auth, error) {
   res, err := http.PostForm(
      "https://android.googleapis.com/auth",
      url.Values{
         "ACCESS_TOKEN": {"1"},
         "Token": {code},
         "service": {"ac2dm"},
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
