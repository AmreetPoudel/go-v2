// package main

import (
	"fmt"
	"sync"
)

type counter struct {
	mu    sync.Mutex
	count int
}

// function to increment count of struct
func (c *counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *counter) getcounter() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
func main() {
	var wg sync.WaitGroup
	totalgoroutines := 1000
	counter := counter{}
	for i := 0; i < totalgoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.increment()

		}()
	}

	wg.Wait()
	fmt.Println("final counter value is ", counter.count)
}
