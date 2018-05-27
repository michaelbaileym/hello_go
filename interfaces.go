package main

import (
	"errors"
	"fmt"
)

type Robot interface {
	PowerOn() error
}

type T850 struct {
	Name string
}

func (a *T850) PowerOn() error {
	return nil
}

type R2D2 struct {
	broken bool
}

func (r *R2D2) PowerOn() error {
	if r.broken {
		return errors.New("R2D2 is broken!")
	} else {
		return nil
	}
}

func Boot(r Robot) error {
	return r.PowerOn()
}

func main() {
	t := T850{
		Name: "Terminator",
	}

	r := R2D2{
		broken: true,
	}

	err := Boot(&r)

	if err != nil {
		fmt.Println("Robot is on!")
	}

	err = Boot(&t)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is powered on!")
	}

}
