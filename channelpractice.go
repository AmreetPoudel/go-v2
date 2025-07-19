// package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {
		// data receive via channnel define first because sender need a receiver to receive that should be ready

		data := <-ch
		fmt.Println(data)
	}()
	//data sending to channel
	ch <- "data from channel"
	time.Sleep(100 * time.Millisecond)

}
