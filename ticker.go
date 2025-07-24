// package main

import (
	"fmt"
	"time"
)

func main() {
	// Step 1: Create ticker — it ticks every 1 second
	ticker := time.NewTicker(1 * time.Second)

	// Step 2: Always defer stopping the ticker to avoid memory leaks
	defer ticker.Stop()

	// Step 3: Create a signal to stop the program after 5 seconds
	done := make(chan bool)

	// Step 4: Start a goroutine to stop after 5 seconds
	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()

	// Step 5: Main loop — listen for ticker or done signal
	for {
		select {
		case t := <-ticker.C:
			// Every tick, print the time
			fmt.Println("Tick at:", t)
		case <-done:
			// Exit the loop
			fmt.Println("Done.")
			return
		}
	}
}
