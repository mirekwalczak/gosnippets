package main

import (
	"fmt"
)

func main() {

	digs := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(digs...))
}

func sum(digits ...int) int {
	var sum int
	for d := range digits {
		sum += d
	}
	return sum
}
