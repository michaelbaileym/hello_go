package main

import (
  "fmt"
)

var s = "hello world"

func main()  {
  fmt.Printf("Print s from the outer block %v\n", s)
  b := true
  if b {
    fmt.Printf("Printing 'b' from the outer block %v\n", b)
  }
}
