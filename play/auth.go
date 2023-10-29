package play

import (
   "errors"
   "io"
   "net/http"
   "net/url"
   "strings"
)

func Exchange(oauth_token string) (*Refresh_Token, error) {
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
   var token Refresh_Token
   token.Raw, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return &token, nil
}

type Access_Token struct {
   v url.Values
}

func (r *Refresh_Token) Unmarshal() error {
   var err error
   r.v, err = parse_query(string(r.Raw))
   if err != nil {
      return err
   }
   return nil
}

type Refresh_Token struct {
   Raw []byte
   v url.Values
}

func (a *Access_Token) Refresh(r Refresh_Token) error {
   res, err := http.PostForm(
      "https://android.clients.google.com/auth", url.Values{
         "Token":      {r.v.Get("Token")},
         "app":        {"com.android.vending"},
         "client_sig": {"38918a453d07199354f8b19af05ec6562ced5788"},
         "service":    {"oauth2:https://www.googleapis.com/auth/googleplay"},
      },
   )
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   body, err := io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   a.v, err = parse_query(string(body))
   if err != nil {
      return err
   }
   return nil
}

func parse_query(query string) (url.Values, error) {
   query = strings.ReplaceAll(query, "\n", "&")
   return url.ParseQuery(query)
}
