package main

import (
	"fmt"
	"strconv"
)

func main() {
	s1 := []string{"3", "4", "5"}
	var s2 []int
	for _, v := range s1 {
		temp, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		s2 = append(s2, temp)
	}
	fmt.Printf("%T %v\n", s1, s1)
	fmt.Printf("%T %v\n", s2, s2)
}
