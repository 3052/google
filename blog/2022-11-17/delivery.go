package main

import (
   "bytes"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "os"
   "time"
)

type version struct {
   major int
   minor int
   patch int
}

func (v version) String() string {
   b := []byte{'8'}
   b = fmt.Appendf(b, "%02v", v.major)
   b = fmt.Append(b, v.minor)
   b = fmt.Appendf(b, "%02v", v.patch)
   b = append(b, "00"...)
   if len(b) > 8 {
      panic(b)
   }
   return string(b)
}

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Host = "play-fe.googleapis.com"
   req.URL.Path = "/fdfe/delivery"
   val := make(url.Values)
   val["doc"] = []string{"com.android.vending"}
   req.URL.Scheme = "https"
   req.Header["User-Agent"] = []string{"Android-Finsky (sdk=9,versionCode=99999999)"}
   req.Header["X-Dfe-Device-Id"] = []string{device}
   // done
   // v := version{major: 4}
   // v := version{major: 5}
   // v := version{major: 6}
   // v := version{major: 7}
   // v := version{major: 8}
   
   // to do
   v := version{major: 9}
   // v := version{major: 10}
   // v := version{major: 11}
   // v := version{major: 12}
   // v := version{major: 13}
   // v := version{major: 14}
   // v := version{major: 15}
   // v := version{major: 16}
   // v := version{major: 17}
   // v := version{major: 18}
   // v := version{major: 19}
   // v := version{major: 20}
   // v := version{major: 21}
   // v := version{major: 22}
   // v := version{major: 23}
   // v := version{major: 24}
   // v := version{major: 25}
   // v := version{major: 26}
   // v := version{major: 27}
   // v := version{major: 28}
   // v := version{major: 29}
   // v := version{major: 30}
   for v.minor = 0; v.minor <= 9; v.minor++ {
      for v.patch = 0; v.patch <= 99; v.patch++ {
         val["vc"] = []string{v.String()}
         req.URL.RawQuery = val.Encode()
         res, err := new(http.Transport).RoundTrip(&req)
         if err != nil {
            panic(err)
         }
         body, err := io.ReadAll(res.Body)
         if err != nil {
            panic(err)
         }
         if bytes.Contains(body, []byte("/by-token/")) {
            fmt.Println(res.Status, "pass", v.String())
            file, err := os.Create(v.String())
            if err != nil {
               panic(err)
            }
            if err := file.Close(); err != nil {
               panic(err)
            }
         } else {
            fmt.Println(res.Status, "fail", v.String())
         }
         if err := res.Body.Close(); err != nil {
            panic(err)
         }
         time.Sleep(199 * time.Millisecond)
      }
   }
}
