package main

import "fmt"

type person struct {
	fname string
	lname string
}

type agent struct {
	person
	gun string
}

func (p person) speak() {
	fmt.Printf("Hello I'm %s %s\n", p.fname, p.lname)
}

func (a agent) speak() {
	fmt.Printf("Hello my name is %s %s and I use %s\n", a.fname, a.lname, a.gun)
}

type speaker interface {
	speak()
}

func saySomething(s speaker) { s.speak() }

func main() {
	p := person{"Rowan", "Atkinson"}
	saySomething(p)

	a := agent{person{"Johny", "English"}, "Beretta 92FS"}
	saySomething(a)
}
