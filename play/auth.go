package play

import (
   "errors"
   "io"
   "net/http"
   "net/url"
   "strings"
)

type GoogleAuth struct {
   Values url.Values
}

func (g GoogleAuth) get_auth() string {
   return g.Values.Get("Auth")
}

func parse_query(query string) (url.Values, error) {
   query = strings.ReplaceAll(query, "\n", "&")
   return url.ParseQuery(query)
}

func (g GoogleToken) get_token() string {
   return g.values.Get("Token")
}

func (g *GoogleToken) New(oauth_token string) error {
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
   g.values, err = parse_query(string(g.Data))
   if err != nil {
      return err
   }
   return nil
}

type GoogleToken struct {
   Data []byte
   values url.Values
}

func (g GoogleToken) Auth() (*GoogleAuth, error) {
   resp, err := http.PostForm(
      "https://android.googleapis.com/auth", url.Values{
         "Token":      {g.get_token()},
         "app":        {"com.android.vending"},
         "client_sig": {"38918a453d07199354f8b19af05ec6562ced5788"},
         "service":    {"oauth2:https://www.googleapis.com/auth/googleplay"},
      },
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   text, err := io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   query, err := parse_query(string(text))
   if err != nil {
      return nil, err
   }
   return &GoogleAuth{query}, nil
}
