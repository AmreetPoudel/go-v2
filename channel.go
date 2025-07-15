// package main

import (
	"fmt"
	"time"
)

func main() {
	// Create an unbuffered channel of type string.
	// Channels are used for communication between goroutines.
	channel := make(chan string)

	// Launch a goroutine (concurrent function execution).
	// This anonymous function sends a message into the channel.
	go func() {
		messageforchannel := "this message will flow to other goroutines via channel"
		channel <- messageforchannel // Send message to channel. Blocks until someone receives.
	}()

	// Launch another goroutine to receive the message from the channel.
	go func() {
		incomingmessage := <-channel // Receive message from channel. Blocks until a message is available.
		fmt.Println(incomingmessage) // Print the received message.
	}()

	// Without this sleep, main() would finish before the goroutines run,
	// and all goroutines would be killed immediately (Go doesn’t wait for them).
	// So we sleep to give time for both send and receive operations to complete.
	time.Sleep(100 * time.Millisecond)
}
