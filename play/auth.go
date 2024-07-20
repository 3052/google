package play

import (
   "errors"
   "io"
   "net/http"
   "net/url"
   "strings"
)

func (g *GoogleToken) Unmarshal(text []byte) error {
   var err error
   g.Values, err = parse_query(string(text))
   if err != nil {
      return err
   }
   return nil
}

func NewGoogleToken(oauth_token string) ([]byte, error) {
   resp, err := http.PostForm(
      "https://android.googleapis.com/auth", url.Values{
         "ACCESS_TOKEN": {"1"},
         "Token":        {oauth_token},
         "service":      {"ac2dm"},
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
   return io.ReadAll(resp.Body)
}

func parse_query(query string) (url.Values, error) {
   query = strings.ReplaceAll(query, "\n", "&")
   return url.ParseQuery(query)
}

type GoogleAuth struct {
   Values url.Values
}

func (g GoogleAuth) get_auth() string {
   return g.Values.Get("Auth")
}

func (g GoogleToken) get_token() string {
   return g.Values.Get("Token")
}

type GoogleToken struct {
   Values url.Values
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
