package main

import "fmt"

func main() {
	//s := "a"
	/*
		switch s {
		case "a":
			fmt.Println("This letter is a")
		case "b":
			fmt.Println("This letter is b")
		default:
			fmt.Println("Not either of those buddy")
		}

		for i := 1; i <= 10; i++ {
			fmt.Println("Count baby", i)
		}

	*/

	numbers := []int{1, 2, 3, 4}

	for i, n := range numbers {
		fmt.Println("The index is ", i)
		fmt.Println("The value of the array is ", n)
	}

}
