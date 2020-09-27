package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	// concatenate two slices:
	// s2... unpack values of a slice,
	// then append these values to first slice
	s3 := append(s1, s2...)
	fmt.Println(s3)
}
