// package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "application is running on 8000")
	})
	http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "stranger"
		}
		fmt.Fprintf(w, "welcome, %s", name)
	})

	address := "127.0.0.1:8000"
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println("server started at : ", address)
	}

}
