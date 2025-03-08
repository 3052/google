package main

import "fmt"

func hello(v ...int) int {
   for _, v1 := range v {
      return v1
   }
   return 0
}

func main() {
   fmt.Println(hello(2, 3, 4, 5, 6))
}

