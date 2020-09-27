package main

import "fmt"

func main() {

	var c chan bool = make(chan bool)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
		c <- true
	}()

	<-c

	fmt.Println("after goroutine")
}
