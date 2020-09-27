package main

import (
	"fmt"
)

func main() {
	var x interface{}

	x = 1
	v, ok := x.(int)
	fmt.Println(v, ok)

	x = "aaa"
	v2, ok := x.(string)
	fmt.Println(v2, ok)
}
