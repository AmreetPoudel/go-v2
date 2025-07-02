// package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "sample.txt"
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("error opening a file .", err)
	}
	data := "this data will be written to file\n"
	_, err = file.Write([]byte(data))
	if err != nil {
		fmt.Println("error writing data to file. ", err)
	}
	defer file.Close()
	//now we will use writestring
	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println(err)
	}

}
