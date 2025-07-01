// package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	data := []byte("my name is amrit poudel and this will come to terminal\n")
	//we are dealing with bytes here but we can simply use writestring and that accept data as string
	n, err := writer.Write(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("wrote %d bytes of data\n", n)

	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
	}

	data1 := "this is string data\n"
	_, err = writer.WriteString(data1)

	if err != nil {
		fmt.Println(err)
	}
	err = writer.Flush()
}
