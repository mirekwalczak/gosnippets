package main

import (
	"errors"
	"fmt"
)

func main() {
	someSlice := []string{"A", "B", "C", "D", "E", "F"}
	fmt.Println(someSlice)
	someSlice, err := delete(someSlice, "C")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(someSlice)
}

func delete(slc []string, element string) ([]string, error) {
	var index int
	var found bool
	for i, s := range slc {
		if s == element {
			index = i
			found = true
		}
	}
	if !found {
		return slc, errors.New("element not found")
	}
	slc = append(slc[:index], slc[index+1:]...)
	return slc, nil
}
