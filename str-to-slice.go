package main

import (
	"fmt"
)

func main() {
	text := "Hello, playground"
	fmt.Println(strToSlice(text))
}
func strToSlice(s string) []string {
	var slice []string
	for _, char := range s {
		slice = append(slice, string(char))
	}
	return slice
}
