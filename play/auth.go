package play

import (
   "bytes"
   "io"
   "net/http"
   "net/url"
   "strings"
)

type Access_Token struct {
   v Values
}

func (a Access_Token) auth() string {
   return a.v["Auth"]
}

type Refresh_Token struct {
   Values Values
}

// accounts.google.com/embedded/setup/android
// the authorization code (oauth_token) looks like this:
// 4/0Adeu5B...
// but it should be supplied here with the prefix:
// oauth2_4/0Adeu5B...
func New_Token(code string) (*Refresh_Token, error) {
   res, err := http.PostForm(
      "https://android.googleapis.com/auth", url.Values{
         "ACCESS_TOKEN": {"1"},
         "Token": {code},
         "service": {"ac2dm"},
      },
   )
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   text, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   var token Refresh_Token
   token.Values = make(Values)
   if err := token.Values.UnmarshalText(text); err != nil {
      return nil, err
   }
   return &token, nil
}

func (r Refresh_Token) Refresh() (*Access_Token, error) {
   // these values take from Android API 28
   res, err := http.PostForm(
      "https://android.googleapis.com/auth", url.Values{
         "Token": {r.token()},
         "app": {"com.android.vending"},
         "client_sig": {"38918a453d07199354f8b19af05ec6562ced5788"},
         "service": {"oauth2:https://www.googleapis.com/auth/googleplay"},
      },
   )
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   text, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   var token Access_Token
   token.v = make(Values)
   if err := token.v.UnmarshalText(text); err != nil {
      return nil, err
   }
   return &token, nil
}

func (r Refresh_Token) token() string {
   return r.Values["Token"]
}

type Values map[string]string

// cs.opensource.google/go/go/+/refs/tags/go1.20.7:src/net/url/url.go
func (vs Values) MarshalText() ([]byte, error) {
   var b bytes.Buffer
   for k, v := range vs {
      if b.Len() >= 1 {
         b.WriteByte('\n')
      }
      b.WriteString(url.QueryEscape(k))
      b.WriteByte('=')
      b.WriteString(url.QueryEscape(v))
   }
   return b.Bytes(), nil
}

// github.com/golang/go/wiki/CodeReviewComments#receiver-type
func (vs Values) UnmarshalText(text []byte) error {
   query := string(text)
   for query != "" {
      var line string
      line, query, _ = strings.Cut(query, "\n")
      key, value, _ := strings.Cut(line, "=")
      var err error
      key, err = url.QueryUnescape(key)
      if err != nil {
         return err
      }
      value, err = url.QueryUnescape(value)
      if err != nil {
         return err
      }
      vs[key] = value
   }
   return nil
}
