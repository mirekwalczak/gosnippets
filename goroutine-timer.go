package main

import (
	"fmt"
	"time"
)

func main() {

	var c chan bool = make(chan bool)

	iterations := 10

	go func() {
		for i := 0; i < iterations; i++ {
			fmt.Printf("%v/%v\n", i, iterations)
			time.Sleep(1 * time.Second)
		}
		c <- true
	}()
	timer := time.NewTimer(time.Second * time.Duration(5))
	select {
	case <-timer.C:
		fmt.Println("Time's out.")
	}
	fmt.Println("After goroutine")
}
