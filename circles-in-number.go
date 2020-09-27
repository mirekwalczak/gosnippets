package main

import "fmt"

// how many circles are in number: 2 -> 0; 10 -> 1; 88 -> 4
func main() {
	var number string
	number = "889302"
	dict := map[string]int{"0": 1, "1": 0, "2": 0, "3": 0, "4": 0, "5": 0, "6": 1, "7": 0, "8": 2, "9": 1}

	var result int
	for _, digit := range number {
		result = result + dict[string(digit)]
	}
	fmt.Printf("Number %v has %v circles.", number, result)
}
