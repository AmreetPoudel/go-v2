// package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("App started...") // logs to stderr by default

	user := "amrit"
	log.Printf("User %s logged in", user)
	err := simulateerror()
	if err != nil {
		log.Println(err)
	}

}

func simulateerror() error {
	return fmt.Errorf("User donot have enough privileges to edit that feature")
}
