// package main

import (
	"fmt"
	"os"
)

func main() {
	// path := "/home/amrit/practice"
	// file := "sample.txt"
	// fmt.Println(filepath.Base(path))
	// fmt.Println(filepath.Join(path, file))
	//now directories creating listing and more
	fmt.Println(os.ReadDir("/Users/amritpoudel/go-V2"))
}
