package main

import "fmt"

type Interface_ interface {
	Method_()
}

type Class_ struct {
	String_ string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t Class_) Method_() {
	fmt.Println(t.String_)
}

func main() {
	// with interface
	var i Interface_ = Class_{"hello"}
	i.Method_()

	// without interface
	var i2 Class_
	i2.String_ = "hello2"
	i2.Method_()
}
