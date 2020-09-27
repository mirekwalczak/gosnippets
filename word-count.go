package main

import (
	"fmt"
	"strings"
)

func WordCountByMirek(str string) map[string]int {
	words := strings.Split(str, " ")
	counts := make(map[string]int)
	for _, word := range words {
		counts[word] += 1
	}
	return counts
}

func WordCountByCaleb(str string) map[string]int {
	counts := map[string]int{}
	for _, word := range strings.Fields(str) {
		counts[word]++
	}
	return counts
}

func main() {
	str := "abc def ghi abc"
	result := WordCountByMirek(str)
	fmt.Println(result["abc"])
	result = WordCountByCaleb(str)
	fmt.Println(result["abc"])
}
