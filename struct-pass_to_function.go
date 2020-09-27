package main

import (
	"fmt"
)

type AuthMethod struct {
	Credentials []string
	APIKey      string
}

func main() {
	Login("192.168.1.1", &AuthMethod{Credentials: []string{"admin", "p4ss"}})
	Login("192.168.1.1", &AuthMethod{APIKey: "12345"})
}

func Login(host string, authmethod *AuthMethod) {
	if len(authmethod.Credentials) > 0 {
		fmt.Println("Authenticated with username and password")
	}
	if len(authmethod.APIKey) > 0 {
		fmt.Println("Authenticated with API key")
	}
}
