package main

import "fmt"
/*
func main()  {
  fn := func(){
    fmt.Println("function called")
  }
  fn()

}
*/

func anotherFunction(f func() string) string {
    return f()
}

func main(){
  fn := func() string {
    return "function called"
  }
  fmt.Println(anotherFunction(fn))
}
