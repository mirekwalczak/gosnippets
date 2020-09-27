package main

import "fmt"

type printer interface {
	printValue()
}
type intValue int

func (d intValue) printValue() {
	fmt.Println("value for intValue is:", d)
}

type floatValue float64

func (p floatValue) printValue() {
	fmt.Println("value for intValue is:", p)
}

func main() {

	// interface allows to treat diferent datatypes as one (printer)

	values := []printer{intValue(42), floatValue(13.3)}
	for _, v := range values {
		v.printValue()
	}
}
