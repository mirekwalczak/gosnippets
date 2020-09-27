package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := make(map[string]int)
	ages["John"] = 22
	ages["Martha"] = 21
	ages["Ted"] = 3
	ages["Mark"] = 30
	fmt.Println("unsorted:\n", ages)

	var names []string
	for name := range ages {
		names = append(names, name)
	}
	fmt.Println("\nsorted:")
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}
