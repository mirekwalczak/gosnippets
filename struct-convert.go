package main

import "fmt"

type bill struct {
	name string
	age  int
}
type alice struct {
	name string
	age  int
}

func main() {
	t := struct {
		name string
		age  int
	}{
		name: "Tom",
		age:  21,
	}
	var b bill
	var a alice
	b = bill(a) // Converting alice to bill must be explicite
	b = t       // Converting t to bill can be implicite because t is literal (unnamed struct)
	// All conversions all possible because structures has the same fields.
	fmt.Println(b)

}
