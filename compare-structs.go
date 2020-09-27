package main

import (
	"fmt"
	"reflect"
)

type s struct {
	name string
	id   int
}

func main() {
	a := s{"a", 1}
	b := s{"a", 1}
	if reflect.DeepEqual(a, b) {
		fmt.Println("matches")
	} else {
		fmt.Println("not matches")
	}
}
