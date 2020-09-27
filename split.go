package main

import (
	"fmt"
	"strings"
)

func main() {
	email := "foo@bar.com"
	components := strings.Split(email, "@")
	fmt.Printf("Username: %s, Domain: %s\n", components[0], components[1])
}
