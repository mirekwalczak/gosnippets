package main

import (
	"fmt"
)

func main() {
	ages := make(map[string]int)
	ages["John"] = 1
	age, ok := ages["Bob"]
	fmt.Println(age, ok)
	age, ok = ages["John"]
	fmt.Println(age, ok)
}
