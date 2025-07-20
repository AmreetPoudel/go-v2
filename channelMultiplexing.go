// package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Create two channels that will simulate different data sources
	channel1 := make(chan string)
	channel2 := make(chan string)

	// Start a goroutine that sends data to channel1 after a random delay
	go func() {
		//simulating getting things work done
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		// things got done now send data to channel
		channel1 <- "Data from channel 1"
	}()

	// Start another goroutine that sends data to channel2 after a random delay
	go func() {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		channel2 <- "Data from channel 2"
	}()

	// Use select to wait for data from either channel
	// This simulates multiplexing: whichever channel sends data first will be received
	select {
	case msg1 := <-channel1:
		fmt.Println("Received:", msg1)
	case msg2 := <-channel2:
		fmt.Println("Received:", msg2)
	case <-time.After(2 * time.Second):
		// This case avoids indefinite blocking
		// If no channel sends data in 2 seconds, we timeout
		fmt.Println("Timeout: no data received in 2 seconds")
	}
}
