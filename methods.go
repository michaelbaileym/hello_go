package main

import (
	"fmt"
	"strconv"
)

type Movies struct {
	Name   string
	Rating float64
}

func (m *Movies) summary() string {
	r := strconv.FormatFloat(m.Rating, 'f', 1, 64)

	return m.Name + ", " + r
}

func main() {
	m := Movies{
		Name:   "The Matrix",
		Rating: 3.2,
	}
	fmt.Println(m.summary())
}
