package main

import (
	"fmt"
)

type Stack struct {
	data []interface{}
}

func NewStack() *Stack {
	s := Stack{data: make([]interface{}, 0)}
	return &s
}

func (s *Stack) Push(data interface{}) {
	s.data = append(s.data, data)
}

func (s *Stack) Pop() interface{} {
	if len(s.data) == 0 {
		return nil
	}
	data := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return data
}

func main() {
	s := NewStack()
	s.Push(10)
	fmt.Println("stack=", *s)
	s.Push("Hello")
	fmt.Println("stack=", *s)
	obj1 := s.Pop()
	fmt.Println("remove element", obj1)
	obj2 := s.Pop()
	fmt.Println("remove element", obj2)
	// To retrieve int value, use type assertion.
	elem1 := obj1.(string)
	fmt.Printf("Type of first elelemnt is %T\n", elem1)
	elem2 := obj2.(int)
	fmt.Printf("Type of second elelemnt is %T\n", elem2)
}
