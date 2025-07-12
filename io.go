// package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	src, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("error: ", err)
	}
	defer src.Close()

	dest, err := os.Create("final.txt")
	if err != nil {
		fmt.Println("error: ", err)
	}
	defer dest.Close()

	bytescopied, err := io.Copy(dest, src)
	if err != nil {
		fmt.Println("error: ", err)

	}
	fmt.Printf("bytes %d copied. \n", bytescopied)

}
