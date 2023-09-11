package main

import (
   "bytes"
   "os"
   "path/filepath"
)

func main() {
   matches, err := filepath.Glob("*.java")
   if err != nil {
      panic(err)
   }
   for _, match := range matches {
      os.Stdout.WriteString(match + "\n")
      text, err := os.ReadFile(match)
      if err != nil {
         panic(err)
      }
      text = bytes.ReplaceAll(text, []byte{'\r'}, nil)
      os.WriteFile(match, text, 0666)
   }
}
