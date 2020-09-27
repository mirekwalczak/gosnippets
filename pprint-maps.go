package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	myMap := map[string]string{
		"jfahdsfa":                    "klsadj kdjfkas dj",
		"8743rkdjfsdlk  djfha hjkas ": "jhfa bhjee",
		"khja kdhj ka kdj":            "984qhhjdal dsfas",
		"jj3 kljq":                    "kjdfkja",
		"98kfadaj 3":                  "kdf ak",
	}
	pprintMap(myMap)
}

// lmax returns maximum length element in slice
func lmax(l []string) int {
	max := 0
	for _, v := range l {
		if elen := utf8.RuneCountInString(v); elen > max {
			max = elen
		}
	}
	return max
}

// pprintMap prints map object in readable form
func pprintMap(m map[string]string) {
	var keys []string
	for key, _ := range m {
		keys = append(keys, key)
	}
	lmax := lmax(keys) + 1
	for i, key := range keys {
		klen := len(key)
		ilen := len(strconv.Itoa(i))
		indentLength := int(lmax - klen - ilen)
		var spaces string
		for i := 0; i < indentLength; i++ {
			spaces += " "
		}
		fmt.Println(i+1, key, spaces, m[key])
	}
}
