package play

import (
   "errors"
   "io"
   "net/http"
   "net/url"
   "strings"
)

type Byte[T any] []byte

func NewToken(oauth_token string) (Byte[Token], error) {
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

type Token [1]Values

func (t *Token) Unmarshal(data Byte[Token]) error {
   (*t)[0].Set(string(data))
   return nil
}

type Auth [1]Values

func (v *Values) Set(data string) error {
   *v = func() ([2]string, bool) {
      var (
         v1 [2]string
         ok bool
      )
      v1[0], data, ok = strings.Cut(data, "=")
      if ok {
         v1[1], data, _ = strings.Cut(data, "\n")
         return v1, true
      }
      return v1, false
   }
   return nil
}

func (v Values) Get(key string) string {
   for {
      v1, ok := v()
      if !ok {
         return ""
      }
      if v1[0] == key {
         return v1[1]
      }
   }
}

type Values func() ([2]string, bool)

func (t Token) Token() string {
   return t[0].Get("Token")
}

func (a Auth) Auth() string {
   return a[0].Get("Auth")
}

func (t Token) Auth() (*Auth, error) {
   resp, err := http.PostForm(
      "https://android.googleapis.com/auth", url.Values{
         "Token":      {t.Token()},
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
   var value Values
   value.Set(string(data))
   return &Auth{value}, nil
}
