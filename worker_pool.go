// package main

import (
	"fmt"
	"time"
)

// ==========================
// Simulate Final Task: Send Email
// ==========================

// sendEmail simulates sending an email to a user.
// This is your real "end job" that a worker(goroutine) does.
func sendEmail(workerID int, userID int) {
	fmt.Printf("Worker %d: sending email to user %d\n", workerID, userID)
	time.Sleep(500 * time.Millisecond) // simulate email latency
}

// ==========================
// Worker Function
// ==========================

// worker is a goroutine that pulls from the jobQueue channel
// and performs sendEmail for each job (userID).
func worker(workerID int, jobQueue <-chan int) {
	for userID := range jobQueue {
		sendEmail(workerID, userID)
	}
	// When jobQueue is closed and emptied, this worker exits.
}

// ==========================
// Main: Coordinator of Job Dispatch
// ==========================
func main() {
	totalEmails := 20 // Total emails to send
	totalWorkers := 4 // Number of concurrent email senders

	jobQueue := make(chan int) // Job queue to hold user IDs

	// -----------------------
	// Step 1: Start Workers
	// -----------------------
	for id := 1; id <= totalWorkers; id++ {
		go worker(id, jobQueue) // Each worker pulls jobs from the queue
	}

	// -----------------------
	// Step 2: Feed Jobs to Queue
	// -----------------------
	for userID := 1; userID <= totalEmails; userID++ {
		jobQueue <- userID // Push job (userID) into the job queue
	}

	// -----------------------
	// Step 3: Close the Queue
	// -----------------------
	close(jobQueue) // No more jobs will be sent

	// -----------------------
	// Step 4: Give Time to Workers to Finish (TEMP solution)
	// -----------------------
	time.Sleep(5 * time.Second) // crude wait; use WaitGroup in prod

	fmt.Println("✅ All emails sent.")
}
