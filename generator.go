package main

import (
	"fmt"
)

func CounterFactory(j int) func() int {
	i := j
	return func() int {
		i++
		return i
	}
}

func main() {
	r := CounterFactory(1)
	fmt.Println(r())
	fmt.Println(r())
	fmt.Println(r())
}
