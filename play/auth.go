package play

import (
   "io"
   "net/http"
   "net/url"
   "strings"
)

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

type Raw_Token []byte

// accounts.google.com/embedded/setup/android
// the authorization code (oauth_token) looks like this:
// 4/0Adeu5B...
// but it should be supplied here with the prefix:
// oauth2_4/0Adeu5B...
func New_Raw_Token(code string) (Raw_Token, error) {
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

type Refresh_Token map[string]string

func New_Refresh_Token(r Raw_Token) (Refresh_Token, error) {
   return parse_query(string(r))
}

func (r Refresh_Token) Refresh() (Access_Token, error) {
   // these values take from Android API 28
   res, err := http.PostForm(
      "https://android.googleapis.com/auth", url.Values{
         "Token": {r.token()},
         "app": {"com.android.vending"},
         "client_sig": {"38918a453d07199354f8b19af05ec6562ced5788"},
         "service": {"oauth2:https://www.googleapis.com/auth/googleplay"},
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
   return parse_query(string(text))
}

func (r Refresh_Token) token() string {
   return r["Token"]
}
