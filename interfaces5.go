package main

import "fmt"

type Sayer interface {
	Say()
}

func saySomething(s Sayer) {
	s.Say()
}

type Duck struct{}

func (d Duck) Say() {
	fmt.Println("Duck says quack quack")
}

type Bird struct{}

func (b Bird) Say() {
	fmt.Println("Bird says tweet tweet")
}

type Man struct{}

func (m Man) Say() {
	fmt.Println("Man says hello")
}

func main() {
	duck := new(Duck)
	saySomething(duck)

	bird := new(Bird)
	saySomething(bird)

	man := new(Man)
	saySomething(man)
}
