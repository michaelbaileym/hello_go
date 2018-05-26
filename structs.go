package main

import "fmt"

type Movie struct {
	Name   string
	Rating float32
}
type SuperHero struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	Number int
	Street string
	City   string
}

type Alarm struct {
	Time  string
	Sound string
}

func NewAlarm(time string) Alarm {
	a := Alarm{
		Time:  time,
		Sound: "Klaxon",
	}
	return a
}

func main() {
	e := SuperHero{
		Name: "Batman",
		Age:  32,
		Address: Address{
			Number: 1503,
			Street: "High Point Dr",
			City:   "Gotham",
		},
	}
	fmt.Println(e.Address.Street)
	fmt.Printf("%+v\n", NewAlarm("9:00"))
}

/*
	m := new(Movie)
	m.Name = "Full Metal Jacket"
	m.Rating = .99999

	fmt.Printf("%+v\n", m)



One way to make a struct
func main() {
	var m Movie
	fmt.Printf("%+v\n", m)

	m.Name = "The Matrix"
	m.Rating = .981
	fmt.Printf("%+v\n", m)
} */
