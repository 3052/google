package main

import (
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Header["Accept"] = []string{"*/*"}
   req.Header["Accept-Encoding"] = []string{"identity"}
   req.Header["Accept-Language"] = []string{"en-US,en;q=0.5"}
   req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded;charset=utf-8"}
   req.Header["Cookie"] = []string{"__Host-GAPS=1:6Eczdv0q7WhaN-iVmKy6MEkWh4T2BA:0bpcYyu1FVB9QZ0i", "GEM=CgptaW51dGVtYWlkEOP-4JKeMQ=="}
   req.Header["Google-Accounts-Xsrf"] = []string{"1"}
   req.Header["Origin"] = []string{"https://accounts.google.com"}
   req.Header["Referer"] = []string{"https://accounts.google.com/embedded/setup/android/signinconsent?flowName=EmbeddedSetupAndroid&cid=1&TL=AGEVcSTKYldYax7-8nVxR0FXjrohuMqXHOQuqY5M6at9gqC0csGGB5ymdYd4409r"}
   req.Header["Sec-Fetch-Dest"] = []string{"empty"}
   req.Header["Sec-Fetch-Mode"] = []string{"cors"}
   req.Header["Sec-Fetch-Site"] = []string{"same-origin"}
   req.Header["Te"] = []string{"trailers"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:101.0) Gecko/20100101 Firefox/101.0"}
   req.Header["X-Same-Domain"] = []string{"1"}
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "accounts.google.com"
   req.URL.Path = "/_/signin/speedbump/embeddedsigninconsent"
   val := make(url.Values)
   val["TL"] = []string{"AGEVcSTKYldYax7-8nVxR0FXjrohuMqXHOQuqY5M6at9gqC0csGGB5ymdYd4409r"}
   val["_reqid"] = []string{"375986"}
   val["ardt"] = []string{"AFWLbD3TbQrQZ2iJbrzfLtwqMR0MRyK1a5FQ-pk8taekLv_MWw9R3ybL_U9fN1YXeRLCm8MzxIhY9DN2DoGYRHLj17XybH46eNruwAtR6NoDekxwWUhIidlU4f0V"}
   val["hl"] = []string{"en"}
   val["rt"] = []string{"j"}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(req_body)
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res_body, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(res_body)
}

var req_body = strings.NewReader(`continue=https%3A%2F%2Faccounts.google.com%2Fo%2Fandroid%2Fauth%3Flang%3Den%26cc%26langCountry%3Den_%26xoauth_display_name%3DAndroid%2BDevice%26tmpl%3Dnew_account%26source%3Dandroid%26return_user_id%3Dtrue&f.req=%5B%22gf.siesic%22%2C1%2C%22US%22%2C%22en%22%2C%22default%22%2C%5B%22XFFfAjryY6A%3D%22%2C%22UpmUvzm8N0Y%3D%22%2C%22R0MVauKaZ8Y%3D%22%2C%22F7TABF8fS3E%3D%22%5D%5D&azt=AFoagUWPCOhtE79kYHsj_yd9ocyeTySQXw%3A1691719581544&dgresponse=%5Bnull%2C%22%2F_%2Fsignin%2Fspeedbump%2Fembeddedsigninconsent%22%2C3%2C0%5D&cookiesDisabled=false&deviceinfo=%5Bnull%2Cnull%2Cnull%2C%5B%5D%2Cnull%2C%22US%22%2Cnull%2Cnull%2Cnull%2C%22EmbeddedSetupAndroid%22%2Cnull%2C%5B0%2Cnull%2C%5B%5D%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2Cnull%2C7%2Cnull%2Cnull%2C%5B%5D%2Cnull%2Cnull%2Cnull%2Cnull%2C%5B%5D%5D%2Cnull%2Cnull%2Cnull%2Cnull%2C0%2Cnull%2C0%2C1%2C%22%22%2Cnull%2Cnull%2C2%2C0%5D&gmscoreversion=undefined&flowName=EmbeddedSetupAndroid&`)
