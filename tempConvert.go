package main

import "fmt"

func tempCovert(cel float32) float32 {
	convTemp := (cel - 32) * .556
	return convTemp
}

func main() {

	fmt.Println(tempCovert(5))
}
