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
   req.Header["Authorization"] = []string{"Bearer ya29.a0AfB_byBYA973oW8A1xpq7OkcqufeMw9ObvTfs-Qh2oir2SfE0RdNj9h7BnE5COg_gmLhl4jDT1WLP1HRSXSAKgafFah24S_9tfxlInPoE2jelVDkHu8mCi3FhccwK4dmLnbdp2c0bCT9d6cdupLdketkfU08qGEKB4mNTty5TgtJobsn_L5q6JCPFlo89maKLbGWa-be0_0ocSx6RaHW4lMEqalPEk1HO0AD04m7VLDdKWPMfx_0n5ky75nA50Zn9Bt7scW_g7V8j6rKKlqJU3fyGgyNdLU6J7TSx1r-DoJV-20JFQSF6_6K2SE1PfVEwP2RswaCgYKAUUSARESFQGOcNnCwMjJJkdsJJfm1QI9lxipYA0333"}
   req.Header["Content-Length"] = []string{"422"}
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Header["Host"] = []string{"play-fe.googleapis.com"}
   req.Header["X-Dfe-Device-Id"] = []string{"36bcfe7abb2ef4eb"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "play-fe.googleapis.com"
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

var req_body = strings.NewReader("\n\x1f\n\x17\n\x11com.mcdonalds.app\x10\x01\x18\x03\x10\x01\x1a\x008\x01B\x03\xa0\x01\x00b\x05\b\xdf\t\x18\x00h\x01\xb2\x01\xdc\x02nonce=EBWK4qj_fNtEs0tox4hbh-G1u4JopnoReuV2oKghIivOHFEeiTi6Sp5RYynfywoaku9lU9HemuJ8qRVxKCCF6jPL1lrWj6i2OGFqYowiAgKzjqPjAgQMFGKYCRWvnxZeQqWjhzLE1yulSwmeFuZ9V380vfBvevWkGK82JemK8cOwWOiYUyYWnKO05ODUrpowvHs8hqFe8HaRM_D3_c9VZYgkMkL-RKsBQ3nn5jvkMDcbNeOt71LZ0INcu28k8lLOaDDJSNb7Ip4aSBLN427tDCnmNFhfKvOJJHwvrSiJCrHTh4GJFOYkfrUI3b1EhcEvA6KVGliLsZMJXJXm8g8mug\xc8\x01\x02\xf2\x01\x02\b\x02\xfa\x01\f\b\xfe\x9a\xae\xa8\x06\x10\xc0ؑ\xac\x01")
