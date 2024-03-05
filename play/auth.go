package play

import (
   "errors"
   "io"
   "net/http"
   "net/url"
   "strings"
)

func parse_query(query string) (url.Values, error) {
   query = strings.ReplaceAll(query, "\n", "&")
   return url.ParseQuery(query)
}

type Access_Token struct {
   Values url.Values
}

type Refresh_Token struct {
   Raw []byte
   Values url.Values
}

func (r *Refresh_Token) Unmarshal() error {
   var err error
   r.Values, err = parse_query(string(r.Raw))
   if err != nil {
      return err
   }
   return nil
}

func Exchange(oauth_token string) (*Refresh_Token, error) {
   res, err := http.PostForm(
      "https://android.googleapis.com/auth", url.Values{
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

func (a *Access_Token) Refresh(r Refresh_Token) error {
   res, err := http.PostForm(
      "https://android.googleapis.com/auth", url.Values{
         "Token":      {r.Values.Get("Token")},
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
   a.Values, err = parse_query(string(body))
   if err != nil {
      return err
   }
   return nil
}
