package main

import (
	"errors"
	"log"
)

var ErrEmptyUserOrPass = errors.New("Username or password cannot be epty")

func main() {
	if err := validate("a", "b"); err != nil {
		log.Println(err)
	}
	if err := validate("", ""); err != nil {
		log.Println(err)
	}
}

func validate(user, pass string) error {
	if user == "" || pass == "" {
		return ErrEmptyUserOrPass
	}
	return nil
}
