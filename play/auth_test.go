package play

import (
   "fmt"
   "os"
   "testing"
)

func TestAuth(t *testing.T) {
   data, err := os.ReadFile("c:/users/steven/google/play/Token")
   if err != nil {
      t.Fatal(err)
   }
   var v Values
   v.New(string(data))
   fmt.Printf("%q\n", v.Get("Token"))
   fmt.Printf("%q\n", v.Get("Token"))
}
