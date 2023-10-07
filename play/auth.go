package play

import (
   "errors"
   "io"
   "net/http"
   "net/url"
   "strings"
)

// accounts.google.com/embedded/setup/v2/android
// Exchange authorization code (oauth_token) for Refresh token
// the code looks like this:
// 4/0Adeu5B...
// but it should be supplied here with the prefix:
// oauth2_4/0Adeu5B...
func New_Refresh_Token(oauth_token string) ([]byte, error) {
   res, err := http.PostForm(
      "https://android.clients.google.com/auth", url.Values{
         "ACCESS_TOKEN": {"1"},
         "Token": {oauth_token},
         "service": {"ac2dm"},
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

func parse_query(query string) (map[string]string, error) {
   values := make(map[string]string)
   for query != "" {
      var line string
      line, query, _ = strings.Cut(query, "\n")
      key, value, _ := strings.Cut(line, "=")
      var err error
      key, err = url.QueryUnescape(key)
      if err != nil {
         return nil, err
      }
      value, err = url.QueryUnescape(value)
      if err != nil {
         return nil, err
      }
      values[key] = value
   }
   return values, nil
}

type Access_Token map[string]string

func (a Access_Token) auth() string {
   return a["Auth"]
}

// Refresh access token
func (h *Header) Set_Authorization(token []byte) error {
   refresh, err := func() (Refresh_Token, error) {
      return parse_query(string(token))
   }()
   if err != nil {
      return err
   }
   res, err := http.PostForm(
      "https://android.clients.google.com/auth", url.Values{
         "Token": {refresh.token()},
         "app": {"com.android.vending"},
         "client_sig": {"38918a453d07199354f8b19af05ec6562ced5788"},
         "service": {"androidmarket"},
      },
   )
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
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
      return "Authorization", "GoogleLogin auth=" + access.auth()
   }
   return nil
}

type Refresh_Token map[string]string

func (r Refresh_Token) token() string {
   return r["Token"]
}
