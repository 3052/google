package play

import (
   "154.pages.dev/protobuf"
   "io"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

func (h *Header) Set_Authorization(token []byte) error {
   refresh, err := func() (Refresh_Token, error) {
      return parse_query(string(token))
   }()
   if err != nil {
      return err
   }
   // Google Services Framework 5
   res, err := http.PostForm(
      "https://android.googleapis.com/auth", url.Values{
         "Token": {refresh.token()},
         "app": {"com.android.vending"},
         "client_sig": {"38918a453d07199354f8b19af05ec6562ced5788"},
         "service": {"oauth2:https://www.googleapis.com/auth/googleplay"},
      },
   )
   if err != nil {
      return err
   }
   defer res.Body.Close()
   access, err := func() (Access_Token, error) {
      b, err := io.ReadAll(res.Body)
      if err != nil {
         return nil, err
      }
      return parse_query(string(b))
   }()
   if err != nil {
      return err
   }
   h.Authorization = func() (string, string) {
      return "Authorization", "Bearer " + access.auth()
   }
   return nil
}

// accounts.google.com/embedded/setup/android
// the authorization code (oauth_token) looks like this:
// 4/0Adeu5B...
// but it should be supplied here with the prefix:
// oauth2_4/0Adeu5B...
func New_Refresh_Token(code string) ([]byte, error) {
   // Google Services Framework 5
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
   Agent func() (string, string)
   Authorization func() (string, string)
   Device func() (string, string)
}

func (h *Header) Set_Agent(single bool) {
   var b []byte
   // `sdk` is needed for `/fdfe/delivery`
   b = append(b, "Android-Finsky (sdk="...)
   // valid range 0 - 0x7FFF_FFFF
   b = strconv.AppendInt(b, 9, 10)
   // com.android.vending
   b = append(b, ",versionCode="...)
   if single {
      // valid range 8_03_2_00_00 - 8_09_1_99_99
      b = strconv.AppendInt(b, 8_09_1_99_99, 10)
   } else {
      // valid range 8_09_2_00_00 - math.MaxInt32
      b = strconv.AppendInt(b, 9_99_9_99_99, 10)
   }
   b = append(b, ')')
   h.Agent = func() (string, string) {
      return "User-Agent", string(b)
   }
}

func (h *Header) Set_Device(device []byte) error {
   var (
      dev Device
      err error
   )
   dev.m, err = protobuf.Consume(device)
   if err != nil {
      return err
   }
   id, err := dev.android_ID()
   if err != nil {
      return err
   }
   h.Device = func() (string, string) {
      return "X-DFE-Device-ID", strconv.FormatUint(id, 16)
   }
   return nil
}

type Refresh_Token map[string]string

func (r Refresh_Token) token() string {
   return r["Token"]
}
