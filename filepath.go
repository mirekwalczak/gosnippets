package main

import (
	"fmt"
	"path/filepath"
)

var path = []string{"/", "nmtlog", "palo", "ispmpan01o", "traffic"}

func main() {
	logpath := filepath.Join(path...)
	fmt.Println(logpath)
}
