package play

import (
   "154.pages.dev/encoding/protobuf"
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

type Refresh_Token map[string]string

func (r Refresh_Token) token() string {
   return r["Token"]
}

func New_Header(token, checkin []byte, single bool) (*Header, error) {
   refresh, err := func() (Refresh_Token, error) {
      return parse_query(string(token))
   }()
   if err != nil {
      return nil, err
   }
   // Android API 21
   res, err := http.PostForm(
      "https://android.googleapis.com/auth", url.Values{
         "Token": {refresh.token()},
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
   var head Header
   head.token, err = parse_query(string(text))
   if err != nil {
      return nil, err
   }
   head.checkin, err = func() (*Android_Checkin, error) {
      m, err := protobuf.Consume(checkin)
      if err != nil {
         return nil, err
      }
      return &Android_Checkin{m}, nil
   }()
   if err != nil {
      return nil, err
   }
   head.single = single
   return &head, nil
}

type Header struct {
   token Access_Token
   checkin *Android_Checkin
   single bool
}

func (h Header) authorization(r *http.Request) {
   r.Header.Set("Authorization", "Bearer " + h.token.auth())
}

func (h Header) device(r *http.Request) error {
   id, err := h.checkin.ID()
   if err != nil {
      return err
   }
   r.Header.Set("X-DFE-Device-ID", strconv.FormatUint(id, 16))
   return nil
}

func (h Header) agent(r *http.Request) {
   var b []byte
   // `sdk` is needed for `/fdfe/delivery`
   b = append(b, "Android-Finsky (sdk="...)
   // valid range 0 - 0x7FFF_FFFF
   b = strconv.AppendInt(b, 9, 10)
   // com.android.vending
   b = append(b, ",versionCode="...)
   if h.single {
      // valid range 8_03_2_00_00 - 8_09_1_99_99
      b = strconv.AppendInt(b, 8_09_1_99_99, 10)
   } else {
      // valid range 8_09_2_00_00 - math.MaxInt32
      b = strconv.AppendInt(b, 9_99_9_99_99, 10)
   }
   b = append(b, ')')
   r.Header.Set("User-Agent", string(b))
}

// accounts.google.com/embedded/setup/android
// the authorization code (oauth_token) looks like this:
// 4/0Adeu5B...
// but it should be supplied here with the prefix:
// oauth2_4/0Adeu5B...
func New_Refresh_Token(code string) ([]byte, error) {
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
