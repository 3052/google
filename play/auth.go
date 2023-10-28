package play

import (
   "errors"
   "io"
   "net/http"
   "net/url"
   "strings"
)

func Authorization(r *http.Request, a Access_Token) {
   r.Header.Set("Authorization", "Bearer " + a["Auth"])
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

type Raw_Refresh_Token []byte

func (r *Raw_Refresh_Token) Do(oauth_token string) error {
   res, err := http.PostForm(
      "https://android.clients.google.com/auth", url.Values{
         "ACCESS_TOKEN": {"1"},
         "Token":        {oauth_token},
         "service":      {"ac2dm"},
      },
   )
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   *r, err = io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   return nil
}

func (r Raw_Refresh_Token) Token() (Refresh_Token, error) {
   return parse_query(string(r))
}

type Refresh_Token map[string]string

// Refresh access token
func (r Refresh_Token) Do() (Access_Token, error) {
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
