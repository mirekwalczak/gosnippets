package main

import (
	"fmt"
)

func safeDiv(num1, num2 int) int {
	defer func() {
		fmt.Println("call recover:", recover())
	}()
	solution := num1 / num2
	return solution
}

func main() {
	fmt.Println("call safeDiv first time:", safeDiv(3, 0))
	fmt.Println("call safeDiv second time:", safeDiv(3, 2))
}
