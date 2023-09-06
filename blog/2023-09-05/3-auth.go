package play

import (
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
)

func token() (string, error) {
   b, err := os.ReadFile(`C:\Users\Steven\google\play\token.txt`)
   if err != nil {
      return "", err
   }
   v, err := url.ParseQuery(strings.ReplaceAll(string(b), "\n", "&"))
   if err != nil {
      return "", err
   }
   return v.Get("Token"), nil
}

func (c checkin) auth() (url.Values, error) {
   req := new(http.Request)
   req.Header = make(http.Header)
   req.Header["App"] = []string{"com.google.android.gms"}
   req.Header["Host"] = []string{"android.clients.google.com"}
   req.Header["User-Agent"] = []string{"GoogleAuth/1.4 sargo PQ3B.190705.003"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "android.clients.google.com"
   req.URL.Path = "/auth"
   val := make(url.Values)
   val["app"] = []string{"com.google.android.gms"}
   val["callerPkg"] = []string{"com.google.android.gms"}
   val["callerSig"] = []string{"38918a453d07199354f8b19af05ec6562ced5788"}
   val["check_email"] = []string{"1"}
   val["client_sig"] = []string{"38918a453d07199354f8b19af05ec6562ced5788"}
   val["device_country"] = []string{"us"}
   val["email"] = []string{"srpen6@gmail.com"}
   val["google_play_services_version"] = []string{"203615028"}
   val["lang"] = []string{"en-gb"}
   val["oauth2_foreground"] = []string{"1"}
   val["sdk_version"] = []string{"28"}
   val["service"] = []string{"oauth2:https://www.googleapis.com/auth/googleplay"}
   val["system_partition"] = []string{"1"}
   val["token_request_options"] = []string{"CAA4AVAB"}
   {
      t, err := token()
      if err != nil {
         return nil, err
      }
      val["Token"] = []string{t}
   }
   {
      id, err := c.id()
      if err != nil {
         return nil, err
      }
      req.Header["Device"] = []string{id}
      val["androidId"] = []string{id}
   }
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   text, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return url.ParseQuery(string(text))
}
