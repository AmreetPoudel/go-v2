// package main

import (
	"fmt"
	"time"
)

func main() {
	go task1()
	go task2()
	time.Sleep(200 * time.Millisecond)
}

func task1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("task1\n")
		time.Sleep(10 * time.Millisecond)

	}
}
func task2() {
	for i := 0; i < 10; i++ {
		fmt.Printf("task2\n")
		time.Sleep(10 * time.Millisecond)

	}
	fmt.Println()
}
