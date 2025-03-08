package play

import (
   "errors"
   "io"
   "iter"
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
   (*t)[0].New(string(data))
   return nil
}

type Auth [1]Values

func (t Token) Token() string {
   return t[0].Get("Token")
}

func (a Auth) Auth() string {
   return a[0].Get("Auth")
}

type Values iter.Seq2[string, string]

func (v *Values) New(data string) {
   *v = func(yield func(string, string) bool) {
      for data != "" {
         var key string
         key, data, _ = strings.Cut(data, "=")
         var value string
         value, data, _ = strings.Cut(data, "\n")
         if !yield(key, value) {
            return 
         }
      }
   }
}

func (v Values) Get(key string) string {
   for key1, value := range v {
      if key1 == key {
         return value
      }
   }
   return ""
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
   value.New(string(data))
   return &Auth{value}, nil
}
