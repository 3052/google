package play

import (
   "errors"
   "io"
   "net/http"
   "net/url"
   "strings"
)

func (r *Refresh_Token) Unmarshal(t Raw_Refresh_Token) error {
   var err error
   r.v, err = parse_query(string(t))
   if err != nil {
      return err
   }
   return nil
}

func (r Refresh_Token) Refresh(a *Access_Token) error {
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
   text, err := io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   a.v, err = parse_query(string(text))
   if err != nil {
      return err
   }
   return nil
}

func parse_query(query string) (url.Values, error) {
   query = strings.ReplaceAll(query, "\n", "&")
   return url.ParseQuery(query)
}

type Raw_Refresh_Token []byte

func Exchange(oauth_token string) (Raw_Refresh_Token, error) {
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

type Refresh_Token struct {
   v url.Values
}

type Access_Token struct {
   v url.Values
}
