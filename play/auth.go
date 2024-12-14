package play

import (
   "errors"
   "io"
   "net/http"
   "net/url"
   "strings"
)

type Values map[string]string

func (v Values) Set(query string) error {
   for query != "" {
      var key string
      key, query, _ = strings.Cut(query, "\n")
      key, value, _ := strings.Cut(key, "=")
      v[key] = value
   }
   return nil
}

func (GoogleToken) Marshal(token string) ([]byte, error) {
   resp, err := http.PostForm(
      "https://android.googleapis.com/auth", url.Values{
         "ACCESS_TOKEN": {"1"},
         "Token":        {token},
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

type GoogleToken struct {
   Values Values
}

type GoogleAuth struct {
   Values Values
}

func (g *GoogleToken) Unmarshal(data []byte) error {
   g.Values = Values{}
   g.Values.Set(string(data))
   return nil
}

func (g GoogleToken) token() string {
   return g.Values["Token"]
}

func (g GoogleToken) Auth() (*GoogleAuth, error) {
   resp, err := http.PostForm(
      "https://android.googleapis.com/auth", url.Values{
         "Token":      {g.token()},
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
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   var auth GoogleAuth
   auth.Values = Values{}
   auth.Values.Set(string(data))
   return &auth, nil
}

func (g GoogleAuth) auth() string {
   return g.Values["Auth"]
}
