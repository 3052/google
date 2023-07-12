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

func (a *Auth) Exchange() error {
   // these values take from Android API 28
   body := url.Values{
      "Token": {a.Get_Token()},
      "app": {"com.android.vending"},
      "client_sig": {"38918a453d07199354f8b19af05ec6562ced5788"},
      "service": {"oauth2:https://www.googleapis.com/auth/googleplay"},
   }.Encode()
   res, err := client.Post(
      "https://android.googleapis.com/auth",
      "application/x-www-form-urlencoded",
      strings.NewReader(body),
   )
   if err != nil {
      return err
   }
   defer res.Body.Close()
   text, err := io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   a.Values, err = parse_query(string(text))
   if err != nil {
      return err
   }
   return nil
}

func (h *Header) Read_Device(name string) error {
   data, err := os.ReadFile(name)
   if err != nil {
      return err
   }
   h.Device.Message, err = protobuf.Unmarshal(data)
   if err != nil {
      return err
   }
   return nil
}

func (h *Header) Read_Auth(name string) error {
   text, err := os.ReadFile(name)
   if err != nil {
      return err
   }
   h.Auth.Values, err = parse_query(string(text))
   if err != nil {
      return err
   }
   return nil
}
type Response struct {
   *http.Response
}

// You can also use host "android.clients.google.com", but it also uses
// TLS fingerprinting.
func New_Auth(email, passwd string) (*Response, error) {
   body := url.Values{
      "Email": {email},
      "Passwd": {passwd},
      "client_sig": {""},
      "droidguard_results": {"-"},
   }.Encode()
   req, err := http.NewRequest(
      "POST", "https://android.googleapis.com/auth", strings.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   client := http.DefaultClient
   client.Transport = &tls.Transport{Spec: tls.Android_API_26}
   res, err := client.Do(req)
   if err != nil {
      return nil, err
   }
   return &Response{res}, nil
}

func (r Response) Write_File(name string) error {
   data, err := io.ReadAll(r.Body)
   if err != nil {
      return err
   }
   return os.WriteFile(name, data, 0666)
}

type Header struct {
   Auth Auth // Authorization
   Device Device // X-DFE-Device-ID
   Single bool
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

func (h Header) Set_Device(head http.Header) error {
   id, err := h.Device.ID()
   if err != nil {
      return err
   }
   head.Set("X-DFE-Device-ID", strconv.FormatUint(id, 16))
   return nil
}

func (h Header) Set_Auth(head http.Header) {
   head.Set("Authorization", "Bearer " + h.Auth.Get_Auth())
}

type Auth struct {
   url.Values
}

func (a Auth) Get_Auth() string {
   return a.Get("Auth")
}

func (a Auth) Get_Token() string {
   return a.Get("Token")
}

// github.com/golang/go/blob/go1.20.4/src/net/url/url.go
func parse_query(query string) (url.Values, error) {
   m := make(url.Values)
   for query != "" {
      var key string
      key, query, _ = strings.Cut(query, "\n")
      key, value, _ := strings.Cut(key, "=")
      key, err := url.QueryUnescape(key)
      if err != nil {
         return nil, err
      }
      value, err = url.QueryUnescape(value)
      if err != nil {
         return nil, err
      }
      m[key] = append(m[key], value)
   }
   return m, nil
}

