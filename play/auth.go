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

type AccessToken struct {
   Values url.Values
}

func (r *RefreshToken) Unmarshal() error {
   var err error
   r.Values, err = parse_query(string(r.Data))
   if err != nil {
      return err
   }
   return nil
}

func (a *AccessToken) Refresh(r RefreshToken) error {
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

type RefreshToken struct {
   Data []byte
   Values url.Values
}

func (r *RefreshToken) New(oauth_token string) error {
   res, err := http.PostForm(
      "https://android.googleapis.com/auth", url.Values{
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
   r.Data, err = io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   return nil
}
