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
   req.Header["Authorization"] = []string{"Bearer ya29.a0AfB_byBFbblkPw8axrrH36QfGChD59B98T9SvEHUyT9qlXvp3_9bo2eo8o1xSNorKLE2Bow0nDFeMeUTkBNZH2e52N1ptv-2gCb0siAmzhONjHiBe7XvmkrszKepFqLHlE8NJuuSEQcGcpoxlZAh4x0WjhsgjqBwlV1svV3-CqjPm44kwH_yiWGWenAcHfM1KY9wgVXYst-VopZ1VZqYz_wfggktAZrjPSardJZ5t8lUoNr3DFZGNajTo_dVZ_ydlzr3pUMR1HGYfbIRxKtGpXKP29lKBBj-G_0hiSIOP0g8_21UHrjH7kGwYE84iDPNB4sRqAaCgYKAXUSARESFQGOcNnC64WSRP9UZTgXivBtLHHlpg0333"}
   req.Header["Content-Length"] = []string{"420"}
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Header["Host"] = []string{"android.clients.google.com"}
   req.Header["X-Dfe-Device-Id"] = []string{"36bcfe7abb2ef4eb"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "android.clients.google.com"
   req.URL.Path = "/fdfe/acquire"
   req.URL.RawPath = ""
   val := make(url.Values)
   val["theme"] = []string{"2"}
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

var req_body = strings.NewReader("\n.\n(\n\"com.noctuasoftware.stellarium_free\x10\x01\x18\x03\x10\x01\x1a\x00\"\x02\b\nB\x00b\x05\b\xfaV\x18\x00h\x01x\x00\xb2\x01\xdc\x02nonce=KNlD7exRsKR1kL2izKR5wfnfoROpOP1ANRQ1ZZR_WyC4OAn3IcK2YNreAlrHjcYFP8wgG4yRdW9oWUTKfMHknSXZaA74Te49ADtZCI93meV1fxnHgJNjX_d9V6YhKuRVmLfqIhZDhZtNnvVSogFilll2w3mE_njdhcfpYwCr4tQkStB7cm8lcoPBfqz7Tf1QngunAQjem_PzKlpPSdHe49BrrXamSugaQL__rnJme5hQ6bWLbVNV2e6N3uQrM2Wk3V1dM2QGUN5SEha6d82PNy_Ox--_xRI5G596uVVJdrSXECBxYMPcsHI3NaLuVkZMuif6pC8lf5eshxwq8N-BQw\xc8\x01\x02")
