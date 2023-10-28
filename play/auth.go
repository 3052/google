package play

import (
   "errors"
   "io"
   "net/http"
   "net/url"
)

type Access_Token map[string]string

type Raw_Refresh_Token []byte

func Auth(oauth_token string) (Raw_Refresh_Token, error) {
   res, err := http.PostForm(
      "https://android.clients.google.com/auth", url.Values{
         "ACCESS_TOKEN": {"1"},
         "Token":        {oauth_token},
         "service":      {"ac2dm"},
      },
   )
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   return io.ReadAll(res.Body)
}

func (r Raw_Refresh_Token) Refresh_Token() (Refresh_Token, error) {
   return parse_query(string(r))
}

type Refresh_Token map[string]string

// Refresh access token
func (r Refresh_Token) Auth() (Access_Token, error) {
   res, err := http.PostForm(
      "https://android.clients.google.com/auth", url.Values{
         "Token":      {r["Token"]},
         "app":        {"com.android.vending"},
         "client_sig": {"38918a453d07199354f8b19af05ec6562ced5788"},
         "service":    {"oauth2:https://www.googleapis.com/auth/googleplay"},
      },
   )
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   body, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return parse_query(string(body))
}
