package play

import (
   "io"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

func parse_query(query string) (map[string]string, error) {
   values := make(map[string]string)
   for query != "" {
      var line string
      line, query, _ = strings.Cut(query, "\n")
      key, value, _ := strings.Cut(line, "=")
      var err error
      key, err = url.QueryUnescape(key)
      if err != nil {
         return nil, err
      }
      value, err = url.QueryUnescape(value)
      if err != nil {
         return nil, err
      }
      values[key] = value
   }
   return values, nil
}

type Access_Token map[string]string

func (a Access_Token) auth() string {
   return a["Auth"]
}

type Header struct {
   Device_ID uint64
   Single bool
   Token Access_Token
}

func (h Header) agent() (string, string) {
   var b []byte
   // `sdk` is needed for `/fdfe/delivery`
   b = append(b, "Android-Finsky (sdk="...)
   // valid range 0 - 0x7FFF_FFFF
   b = strconv.AppendInt(b, 9, 10)
   // com.android.vending
   b = append(b, ",versionCode="...)
   if h.Single {
      // valid range 8_03_2_00_00 - 8_09_1_99_99
      b = strconv.AppendInt(b, 8_09_1_99_99, 10)
   } else {
      // valid range 8_09_2_00_00 - math.MaxInt32
      b = strconv.AppendInt(b, 9_99_9_99_99, 10)
   }
   b = append(b, ')')
   return "User-Agent", string(b)
}

func (h Header) authorization() (string, string) {
   return "Authorization", "Bearer " + h.Token.auth()
}

func (h Header) device() (string, string) {
   return "X-DFE-Device-ID", strconv.FormatUint(h.Device_ID, 16)
}

type Raw_Token []byte

// accounts.google.com/embedded/setup/android
// the authorization code (oauth_token) looks like this:
// 4/0Adeu5B...
// but it should be supplied here with the prefix:
// oauth2_4/0Adeu5B...
func New_Raw_Token(code string) (Raw_Token, error) {
   // Android API 21
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
   return io.ReadAll(res.Body)
}

func (r Raw_Token) Refresh() (Refresh_Token, error) {
   return parse_query(string(r))
}

type Refresh_Token map[string]string

func (r Refresh_Token) token() string {
   return r["Token"]
}

func (r Refresh_Token) Access() (Access_Token, error) {
   // Android API 21
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
   return parse_query(string(text))
}
