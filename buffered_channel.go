// package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a buffered channel with capacity = 2
	ch := make(chan string, 2)

	// CODE FLOW STARTS HERE

	// -------------------------------
	// 1. Send two items — will NOT block
	// -------------------------------

	fmt.Println("Sending 1st item to channel")
	ch <- "message 1" // ✅ buffer has space
	fmt.Println("Sent 1st")

	fmt.Println("Sending 2nd item to channel")
	ch <- "message 2" // ✅ still space (now full)
	fmt.Println("Sent 2nd")

	// -------------------------------
	// 2. Attempting 3rd send — will BLOCK until space is freed
	// -------------------------------

	go func() {
		fmt.Println("Sending 3rd item to channel — this will block because buffer is full")
		ch <- "message 3"       // ❌ blocks because buffer is full
		fmt.Println("Sent 3rd") // This line prints **after** receiver takes one
	}()

	// Small delay to show that 3rd send is blocked
	time.Sleep(2 * time.Second)

	// -------------------------------
	// 3. Receive one item — this unblocks the sender
	// -------------------------------
	fmt.Println("Receiving 1st item from channel")
	val := <-ch // Frees up space in buffer
	fmt.Println("Received:", val)

	// -------------------------------
	// 4. Receive the rest
	// -------------------------------
	val2 := <-ch
	fmt.Println("Received:", val2)

	val3 := <-ch
	fmt.Println("Received:", val3)

	// Now buffer is empty again
}
