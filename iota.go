package main

import (
	"fmt"
)

func main() {
	const (
		Mon int = iota
		Tue
		Wed
		Thu
		Fri
		Sat
		Sun
	)
	fmt.Println(Mon, Tue, Wed, Thu, Fri, Sat, Sun)
}
