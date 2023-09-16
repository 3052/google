package main

import (
   "154.pages.dev/google/aurora"
   "154.pages.dev/google/gplayapi"
   "encoding/json"
   "os"
)

func main() {
   var auth aurora.Aurora_OSS
   {
      b, err := os.ReadFile("aurora.json")
      if err != nil {
         panic(err)
      }
      if err := auth.Unmarshal(b); err != nil {
         panic(err)
      }
   }
   c, err := gplayapi.NewClient("srpen6@gmail.com", auth.Auth_Token)
   if err != nil {
      panic(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   enc.Encode(c.AuthData)
}
