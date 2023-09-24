package play

import (
   "io"
   "net/http"
   "net/url"
)

func (h *Header) Set_Authorization(token []byte) error {
   refresh, err := func() (Refresh_Token, error) {
      return parse_query(string(token))
   }()
   if err != nil {
      return err
   }
   // Google Services Framework 5
   res, err := http.PostForm(
      "https://android.googleapis.com/auth", url.Values{
         "Token": {refresh.token()},
         "app": {"com.android.vending"},
         "client_sig": {"38918a453d07199354f8b19af05ec6562ced5788"},
         "service": {"oauth2:https://www.googleapis.com/auth/googleplay"},
      },
   )
   if err != nil {
      return err
   }
   defer res.Body.Close()
   access, err := func() (Access_Token, error) {
      b, err := io.ReadAll(res.Body)
      if err != nil {
         return nil, err
      }
      return parse_query(string(b))
   }()
   if err != nil {
      return err
   }
   h.Authorization = func() (string, string) {
      return "Authorization", "Bearer " + access.auth()
   }
   return nil
}

// accounts.google.com/embedded/setup/android
// the authorization code (oauth_token) looks like this:
// 4/0Adeu5B...
// but it should be supplied here with the prefix:
// oauth2_4/0Adeu5B...
func New_Refresh_Token(code string) ([]byte, error) {
   // Google Services Framework 5
   res, err := http.PostForm(
      "https://android.googleapis.com/auth", url.Values{
         "ACCESS_TOKEN": {"1"},
         "Token": {code},
         "service": {"ac2dm"},
      },
   )
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return io.ReadAll(res.Body)
}
