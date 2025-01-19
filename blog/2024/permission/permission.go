package main

import (
   "encoding/xml"
   "fmt"
   "os"
   "strings"
)

func include(s string) bool {
   switch {
   case strings.HasPrefix(s, "android.permission."):
      return true
   case strings.HasPrefix(s, "com.android.permission."):
      return true
   }
   return false
}

func main() {
   text, err := os.ReadFile("AndroidManifest.xml")
   if err != nil {
      panic(err)
   }
   var manifest struct {
      UsesPermission []struct {
         Name string `xml:"name,attr"`
      } `xml:"uses-permission"`
   }
   if err := xml.Unmarshal(text, &manifest); err != nil {
      panic(err)
   }
   for _, permission := range manifest.UsesPermission {
      if include(permission.Name) {
         fmt.Printf("<permission name=%q/>\n", permission.Name)
      }
   }
}
