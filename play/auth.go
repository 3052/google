package play

import (
   "bytes"
   "io"
   "net/http"
   "net/url"
   "strings"
)

func (r Refresh_Token) Refresh() (*Access_Token, error) {
   // these values take from Android API 28
   res, err := http.PostForm(
      "https://android.googleapis.com/auth", url.Values{
         "Token": {r.Token()},
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
   token.v = make(values)
   if err := token.v.UnmarshalText(text); err != nil {
      return nil, err
   }
   return &token, nil
}

type Access_Token struct {
   v values
}

func (a Access_Token) Auth() string {
   return a.v["Auth"]
}

// cs.opensource.google/go/go/+/refs/tags/go1.20.7:src/net/url/url.go
func (vs values) MarshalText() ([]byte, error) {
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

type values map[string]string

// github.com/golang/go/wiki/CodeReviewComments#receiver-type
func (vs values) UnmarshalText(text []byte) error {
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
   token.v = make(values)
   if err := token.v.UnmarshalText(text); err != nil {
      return nil, err
   }
   return &token, nil
}

func (r Refresh_Token) Token() string {
   return r.v["Token"]
}

type Refresh_Token struct {
   v values
}
