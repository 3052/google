package play

import (
   "errors"
   "io"
   "net/http"
   "net/url"
   "strings"
)

func (g *GoogleToken) Parse() error {
   var err error
   g.v, err = parse_query(string(g.Data))
   if err != nil {
      return err
   }
   return nil
}

func parse_query(query string) (url.Values, error) {
   query = strings.ReplaceAll(query, "\n", "&")
   return url.ParseQuery(query)
}

type GoogleAuth struct {
   v url.Values
}

func (a *GoogleAuth) Auth(t GoogleToken) error {
   res, err := http.PostForm(
      "https://android.googleapis.com/auth", url.Values{
         "Token":      {t.GetToken()},
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

func (g GoogleAuth) GetAuth() string {
   return g.v.Get("Auth")
}

type GoogleToken struct {
   Data []byte
   v url.Values
}

func (g *GoogleToken) Auth(oauth_token string) error {
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
      var b strings.Builder
      res.Write(&b)
      return errors.New(b.String())
   }
   g.Data, err = io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   return nil
}

func (g GoogleToken) GetToken() string {
   return g.v.Get("Token")
}

