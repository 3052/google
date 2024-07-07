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

type GoogleAuth struct {
   v url.Values
}

func (g *GoogleAuth) Auth(token GoogleToken) error {
   resp, err := http.PostForm(
      "https://android.googleapis.com/auth", url.Values{
         "Token":      {token.get_token()},
         "app":        {"com.android.vending"},
         "client_sig": {"38918a453d07199354f8b19af05ec6562ced5788"},
         "service":    {"oauth2:https://www.googleapis.com/auth/googleplay"},
      },
   )
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return errors.New(b.String())
   }
   text, err := io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   g.v, err = parse_query(string(text))
   if err != nil {
      return err
   }
   return nil
}

func (g GoogleAuth) get_auth() string {
   return g.v.Get("Auth")
}

type GoogleToken struct {
   Data []byte
   v url.Values
}

func (g *GoogleToken) Auth(oauth_token string) error {
   resp, err := http.PostForm(
      "https://android.googleapis.com/auth", url.Values{
         "ACCESS_TOKEN": {"1"},
         "Token":        {oauth_token},
         "service":      {"ac2dm"},
      },
   )
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return errors.New(b.String())
   }
   g.Data, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   return nil
}

func (g *GoogleToken) Unmarshal() error {
   var err error
   g.v, err = parse_query(string(g.Data))
   if err != nil {
      return err
   }
   return nil
}

func (g GoogleToken) get_token() string {
   return g.v.Get("Token")
}
