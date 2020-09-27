package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func f(url string, ch chan bool) {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(len(body), ":", url)
	ch <- true
}

func main() {
	t0 := time.Now()
	urls := []string{
		"https://www.google.com/",
		"https://www.microsoft.com/",
		"https://www.apple.com/",
		"https://www.amazon.com/",
	}
	task := make(chan bool)
	for _, url := range urls {
		go f(url, task)
	}
	for range urls { // for _, _ = range urls
		<-task
	}
	elapsed := time.Since(t0)
	fmt.Printf("Took %s\n", elapsed)
}
