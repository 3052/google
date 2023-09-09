package main

import (
   "154.pages.dev/google/acquire"
   "fmt"
   "net/url"
   "os"
   "strings"
)

func new_token() (string, error) {
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

func main() {
   token, err := new_token()
   if err != nil {
      panic(err)
   }
   c, err := acquire.NewClientWithDeviceInfo("srpen6@gmail.com", token)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%+v\n", c.AuthData.GsfID)
}
