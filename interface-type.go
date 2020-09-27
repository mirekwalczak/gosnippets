/*
Type swith shows what concreate type was assigned to interface
*/

package main

import (
	"fmt"
)

type Rider interface {
	Ride()
}

func Describe(r Rider) {
	switch r.(type) {
	case Bike:
		fmt.Println("Rider is a Bike")
	case Car:
		fmt.Println("Rider is a Car")
	default:
		fmt.Println("Rider is either a Bike neither a Car")

	}

}

type Bike struct{ name string }

func (b Bike) Ride() {
	fmt.Printf("%s is riding\n", b)
}

type Car struct{ name string }

func (c Car) Ride() {
	fmt.Printf("%s is riding\n", c)
}

type Sth int

func (s Sth) Ride() {}

func main() {
	var r Rider
	var b Bike
	var c Car
	var s Sth
	_ = b
	_ = c
	_ = s
	r = s
	Describe(r)

}
