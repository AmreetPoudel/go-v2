// package main

import (
	"fmt"
	"net/http"
)

func main() {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Println("error: ", err)
	}
	defer resp.Body.Close()
	fmt.Println("------")

	fmt.Println(resp.Proto)
	fmt.Println(resp.Header)
	fmt.Println(resp.StatusCode)

	// data, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("error reading response body: ", err)
	// }
	// fmt.Println(string(data))

}
