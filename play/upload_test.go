package play

import (
   "fmt"
   "net/url"
   "os"
   "strings"
   "testing"
)

func Test_Upload(t *testing.T) {
   token, err := new_token()
   if err != nil {
      panic(err)
   }
   c, err := gplayapi.NewClientWithDeviceInfo("srpen6@gmail.com", token)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%+v\n", c.AuthData.GsfID)
}

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
