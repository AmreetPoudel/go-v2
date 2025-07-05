// package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "sample.txt"
	file, err := os.Open(filename)
	defer func() {
		fmt.Println("closing the file")
		file.Close()
	}()
	if err != nil {
		fmt.Println("error opening file: ", err)
	}
	buff := make([]byte, 1024)
	_, err = file.Read(buff)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(buff))
}
