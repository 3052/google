package play

import (
   "errors"
   "io"
   "net/http"
   "net/url"
   "strings"
)

type Values func() ([2]string, bool)

func (v *Values) Set(data string) error {
   *v = func() ([2]string, bool) {
      var (
         v1 [2]string
         ok bool
      )
      v1[0], data, ok = strings.Cut(data, "=")
      if ok {
         v1[1], data, _ = strings.Cut(data, "\n")
         return v1, true
      }
      return v1, false
   }
   return nil
}

func (v Values) Get(key string) string {
   for {
      v1, ok := v()
      if !ok {
         return ""
      }
      if v1[0] == key {
         return v1[1]
      }
   }
}
