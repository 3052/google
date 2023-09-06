package gplayapi

import (
   "fmt"
   "net/url"
   "os"
   "strings"
   "testing"
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

func TestPlay(t *testing.T) {
   token, err := new_token()
   if err != nil {
      t.Fatal(err)
   }
   c, err := _NewClient("srpen6@gmail.com", token)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", c)
}
