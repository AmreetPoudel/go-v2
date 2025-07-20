// package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string) // unbuffered channel for synchronization

	// Start 5 goroutines to send data to the channel
	for i := 1; i <= 5; i++ {
		go func(id int) {
			// simulate work using sleep
			time.Sleep(time.Duration(id) * 100 * time.Millisecond)

			// send message to channel — this will block until receiver is ready
			ch <- fmt.Sprintf("Message from goroutine %d", id)
		}(i)
	}

	// Receiver goroutine using for...range
	// Receives until the channel is closed
	go func() {
		for msg := range ch { // blocks if no data is in the channel
			fmt.Println("Received:", msg)
		}
		fmt.Println("Receiver done")
	}()

	// Wait enough time for all goroutines to send data
	time.Sleep(1 * time.Second)

	// Close the channel — this is necessary to exit the range loop
	// If you don’t close, `range` will wait forever and never exit
	close(ch)

	// Wait a bit to let receiver finish printing
	time.Sleep(100 * time.Millisecond)
}
