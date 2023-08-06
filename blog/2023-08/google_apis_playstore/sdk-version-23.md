# SDK version 23

1. Android Studio
2. `google_apis_playstore` image

starting with this:

<http://dl.google.com/android/repository/sys-img/google_apis_playstore/x86-24_r12.zip>

this Go program fails:

~~~go
package main

import (
   "fmt"
   "net/http"
   "time"
)

func main() {
   req, err := http.NewRequest("HEAD", "http://dl.google.com", nil)
   if err != nil {
      panic(err)
   }
   for r := 0; r <= 64; r++ {
      req.URL.Path = "/android/repository/sys-img/google_apis_playstore"
      req.URL.Path += fmt.Sprintf("/x86-23_r%02v.zip", r)
      res, err := http.DefaultClient.Do(req)
      if err != nil {
         panic(err)
      }
      fmt.Println(res.Status, req.URL)
      if res.StatusCode == http.StatusOK {
         break
      }
      time.Sleep(time.Second)
   }
}
~~~
