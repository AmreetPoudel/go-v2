// package main

import (
	"fmt"
	"strconv"
	// "strconv"
)

func main() {

	numstr := "25"
	// fmt.Println(numstr + 1)

	// num, _ := strconv.Atoi(numstr)
	// fmt.Println(num + 1)
	num, err := strconv.ParseInt(numstr, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num + 1)
	binstr := "10101010111"
	bin, err := strconv.ParseInt(binstr, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bin)
}
