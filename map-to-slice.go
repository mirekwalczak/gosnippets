package main

import (
	"fmt"
)

type person struct {
	name    string
	age     int
	male    bool
	counter int
}

// key of people map is person
type people map[person]int

// ConvToSlice converts people map to slice of person
func (pp *people) ConvToSlice() []person {
	var p []person
	for k, v := range *pp {
		k.counter = v
		p = append(p, k)
	}
	return p
}

func (pp *people) AddPerson(p person) {
	(*pp)[p]++
}

func main() {
	pp := people{
		person{name: "Adam", age: 25, male: true}:      1, // counter of Adam = 1 (default)
		person{name: "Veronika", age: 30, male: false}: 1,
		person{name: "Tom", age: 13, male: true}:       1,
	}
	fmt.Println(pp.ConvToSlice())
	pp.AddPerson(person{name: "kinga", age: 43, male: false})
	pp.AddPerson(person{name: "Adam", age: 25, male: true}) // couter of Adam = 2
	pp.AddPerson(person{name: "Adam", age: 25, male: true}) // couter of Adam = 3
	fmt.Println(pp.ConvToSlice())

}
