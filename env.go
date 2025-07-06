// package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file ")
	}

	// Read values
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")

	// Print them
	fmt.Println("DB_HOST:", host)
	fmt.Println("DB_USER:", user)
	fmt.Println("DB_PASS:", pass)
}
