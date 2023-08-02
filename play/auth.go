package play

import (
   "154.pages.dev/encoding/protobuf"
   "154.pages.dev/tls"
   "io"
   "net/http"
   "net/url"
   "os"
   "strconv"
   "strings"
)

type Auth struct {
   v url.Values
}

type auth map[string]string

// godocs.io/flag#Value
func (a auth) String() string {
   var buf strings.Builder
   for k, v := range a {
      if buf.Len() >= 1 {
         buf.WriteByte('\n')
      }
      buf.WriteString(url.QueryEscape(k))
      buf.WriteByte('=')
      buf.WriteString(url.QueryEscape(v))
   }
   return buf.String()
}

// godocs.io/flag#Value
func (a auth) Set(query string) (err error) {
   for query != "" {
      var key string
      key, query, _ = strings.Cut(query, "\n")
      if key == "" {
         continue
      }
      key, value, _ := strings.Cut(key, "=")
      key, err1 := url.QueryUnescape(key)
      if err1 != nil {
         if err == nil {
            err = err1
         }
         continue
      }
      value, err1 = url.QueryUnescape(value)
      if err1 != nil {
         if err == nil {
            err = err1
         }
         continue
      }
      a[key] = value
   }
   return err
}

// cs.opensource.google/go/go/+/refs/tags/go1.20.7:src/net/url/url.go
func parse_query(m url.Values, query string) (err error) {
   for query != "" {
      var key string
      key, query, _ = strings.Cut(query, "\n")
      if key == "" {
         continue
      }
      key, value, _ := strings.Cut(key, "=")
      key, err1 := url.QueryUnescape(key)
      if err1 != nil {
         if err == nil {
            err = err1
         }
         continue
      }
      value, err1 = url.QueryUnescape(value)
      if err1 != nil {
         if err == nil {
            err = err1
         }
         continue
      }
      m[key] = append(m[key], value)
   }
   return err
}

// You can also use host "android.clients.google.com", but it also uses
// TLS fingerprinting.
func New_Auth(email, passwd string) (*Response, error) {
   client := *http.DefaultClient
   client.Transport = &tls.Transport{Spec: tls.Android_API_26}
   res, err := client.PostForm(
      "https://android.googleapis.com/auth",
      url.Values{
         "Email": {email},
         "Passwd": {passwd},
         "client_sig": {""},
         "droidguard_results": {"-"},
      },
   )
   if err != nil {
      return nil, err
   }
   return &Response{res}, nil
}

func (h *Header) Read_Auth(name string) error {
   text, err := os.ReadFile(name)
   if err != nil {
      return err
   }
   return parse_query(h.Auth.v, string(text))
}

func (a *Auth) Exchange() error {
   // these values take from Android API 28
   res, err := http.PostForm(
      "https://android.googleapis.com/auth",
      url.Values{
         "Token": {a.Get_Token()},
         "app": {"com.android.vending"},
         "client_sig": {"38918a453d07199354f8b19af05ec6562ced5788"},
         "service": {"oauth2:https://www.googleapis.com/auth/googleplay"},
      },
   )
   if err != nil {
      return err
   }
   defer res.Body.Close()
   text, err := io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   return parse_query(a.v, string(text))
}

func (a Auth) Get_Auth() string {
   return a.v.Get("Auth")
}

func (a Auth) Get_Token() string {
   return a.v.Get("Token")
}

type Header struct {
   Auth Auth // Authorization
   Device Device // X-DFE-Device-ID
   Single bool
}

func (h *Header) Read_Device(name string) error {
   data, err := os.ReadFile(name)
   if err != nil {
      return err
   }
   h.Device.m, err = protobuf.Unmarshal(data)
   if err != nil {
      return err
   }
   return nil
}

func (h Header) Set_Agent(head http.Header) {
   var b []byte
   // `sdk` is needed for `/fdfe/delivery`
   b = append(b, "Android-Finsky (sdk="...)
   // valid range 0 - 0x7FFF_FFFF
   b = strconv.AppendInt(b, 9, 10)
   // com.android.vending
   b = append(b, ",versionCode="...)
   if h.Single {
      // valid range 8032_0000 - 8091_9999
      b = strconv.AppendInt(b, 8091_9999, 10)
   } else {
      // valid range 8092_0000 - 0x7FFF_FFFF
      b = strconv.AppendInt(b, 9999_9999, 10)
   }
   b = append(b, ')')
   head.Set("User-Agent", string(b))
}

func (h Header) Set_Auth(head http.Header) {
   head.Set("Authorization", "Bearer " + h.Auth.Get_Auth())
}

func (h Header) Set_Device(head http.Header) error {
   id, err  := h.Device.ID()
   if err != nil {
      return err
   }
   head.Set("X-DFE-Device-ID", strconv.FormatUint(id, 16))
   return nil
}
