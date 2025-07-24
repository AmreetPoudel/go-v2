// package main

import (
	"fmt"
	"time"
)

// func main() {

// 	fmt.Println("start: ", time.Now())
// 	// create timer that fires after 2 seconds
// 	timer := time.NewTimer(2 * time.Second)
// 	//wait for it to fire
// 	<-timer.C
// 	fmt.Println("timer fired: ", time.Now())

// }

func main() {
	// Create a timer that will fire after 3 seconds
	timer := time.NewTimer(3 * time.Second)

	// Simulate some work that may take variable time
	done := make(chan string)

	go func() {
		// Simulate work taking 5 seconds (slower than the timer)
		time.Sleep(1 * time.Second)
		done <- " Work completed"
	}()

	// Use select to wait for whichever comes first: work done or timeout
	select {
	case result := <-done:
		// If work finishes first
		fmt.Println(result)
	case <-timer.C:
		// If timer expires first
		fmt.Println("Timeout: work took too long")
	}
}
