// package main

import (
	"fmt"
	"sync"
	"time"
)

func sendEmail(workerID int, userID int) {
	fmt.Printf(" Worker %d: starting sending email to user %d\n", workerID, userID)
	time.Sleep(100 * time.Millisecond) // lets assume delay
	fmt.Printf("email send to user %d by %d worker\n", userID, workerID)
}
func worker(workerID int, jobQueue <-chan int, wg *sync.WaitGroup) {
	for userID := range jobQueue {
		sendEmail(workerID, userID)
		wg.Done()
	}

}

func main() {
	totalEmails := 20
	totalWorkers := 4
	jobQueue := make(chan int)
	var wg sync.WaitGroup
	wg.Add(totalEmails) // now it waits for 20 wg.done()

	//1. start the worker
	for id := 1; id <= totalWorkers; id++ {
		go worker(id, jobQueue, &wg)
	}
	//feed jobs to queue  and waitgroup
	for userID := 1; userID <= totalEmails; userID++ {
		jobQueue <- userID
	}
	// Step 3: Close Queue
	// -----------------------
	close(jobQueue) // No more jobs
	// -----------------------
	// Step 4: Wait for All Jobs to Complete
	// -----------------------
	wg.Wait() // Blocks here until all wg.Done() are called

	fmt.Println("✅ All emails sent.")

}
