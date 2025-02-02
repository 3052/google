package play

import (
   "errors"
   "io"
   "net/http"
   "net/url"
   "strings"
)

type Values map[string]string

func (v Values) Set(data string) error {
   for data != "" {
      var key string
      key, data, _ = strings.Cut(data, "\n")
      key, value, _ := strings.Cut(key, "=")
      v[key] = value
   }
   return nil
}

type Token struct {
   Values Values
}

func (Token) Marshal(oauth_token string) ([]byte, error) {
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
      var data strings.Builder
      resp.Write(&data)
      return nil, errors.New(data.String())
   }
   return io.ReadAll(resp.Body)
}

func (t *Token) Unmarshal(data []byte) error {
   t.Values = Values{}
   t.Values.Set(string(data))
   return nil
}

func (a Auth) auth() string {
   return a.Values["Auth"]
}

func (t Token) token() string {
   return t.Values["Token"]
}

type Auth struct {
   Values Values
}

func (t Token) Auth() (*Auth, error) {
   resp, err := http.PostForm(
      "https://android.googleapis.com/auth", url.Values{
         "Token":      {t.token()},
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
      var data strings.Builder
      resp.Write(&data)
      return nil, errors.New(data.String())
   }
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   value := Values{}
   value.Set(string(data))
   return &Auth{value}, nil
}
