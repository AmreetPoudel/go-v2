// package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	// Non-blocking send - this will work because buffer has space
	go func() {
		select {
		case ch <- "first message":
			fmt.Println("Sent: first message")
		default:
			fmt.Println("Send blocked")
		}
	}()
	// Non-blocking send - this will fail because buffer is full now
	go func() {

		select {
		case ch <- "second message":
			fmt.Println("Sent: second message")
		default:
			fmt.Println("Send blocked: buffer is full")
		}
	}()
	time.Sleep(2 * time.Second)

	// Non-blocking receive - will work, since there's 1 item in buffer
	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("Receive blocked: buffer is empty")
	}

	// Non-blocking receive - will fail because buffer is now empty
	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("Receive blocked: nothing to receive")
	}
}
