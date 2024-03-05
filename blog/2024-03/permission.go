package main

import (
   "encoding/xml"
   "fmt"
   "os"
   "strings"
)

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
      if strings.HasPrefix(permission.Name, "android.permission.") {
         fmt.Printf("<permission name=%q/>\n", permission.Name)
      }
   }
}
