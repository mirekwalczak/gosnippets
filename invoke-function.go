package main

import (
	"fmt"
)

func main() {
	v := func(param int) int {
		return param
	}(3)
	fmt.Println(v)

}
