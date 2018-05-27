package main

import "fmt"

func Half(numberToHalf int) (int, error) {
	if numberToHalf%2 != 0 {
		return -1, fmt.Errorf("Cannot halve number %v", numberToHalf)
	}
	return numberToHalf / 2, nil
}

func main() {
	n, err := Half(20)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
}

/*
import (
	"fmt"
	"io/ioutil"
)

func main() {
	file, err := ioutil.ReadFile("foo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("%s", file)
} */
