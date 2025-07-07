// package main

import (
	"log"
	"os"
)

func main() {
	// Open (or create) the log file ,append,create,writeonly
	f, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer f.Close()

	// Send log output to the file
	log.SetOutput(f)

	log.Println("App started")
	log.Println("Doing some work...")
	log.Println("Finished.")
}
