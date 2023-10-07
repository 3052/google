package main

import (
   "fmt"
   "net/http"
   "time"
)

const (
   target    = "google_apis"
   abi       = "x86"
   api_level = 26
   revision  = 16
)

func main() {
   req, err := http.NewRequest("HEAD", "http://dl.google.com", nil)
   if err != nil {
      panic(err)
   }
   for r := 0; r < revision; r++ {
      req.URL.Path = "/android/repository/sys-img/"
      req.URL.Path += fmt.Sprintf("%v/%v-%v", target, abi, api_level)
      req.URL.Path += fmt.Sprintf("_r%02v.zip", r)
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
