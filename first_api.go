package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// to make temp data lets create struct
type Book struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	// since we donot have nay data so lets create data by ourself
	Books := []Book{
		{Id: 1, Name: "a2z by amrit poudel", Price: 999999.99},
		{Id: 2, Name: "laziness by anita poudel", Price: 99.99},
		{Id: 3, Name: "more laziness by anusha poudel ", Price: 9.99},
	}

	// lets create haldler for /books/
	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		// lets extract id from  url /books/2  => so we need to extract 2
		path := r.URL.Path // /books/2
		if len(path) <= len("/books/") {
			http.Error(w, "id is missing", http.StatusBadRequest)
		}
		idStr := path[len("/books/"):] // get string after "/books/"
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid book id", http.StatusBadRequest)
			return
		}
		// searching book in slice
		for _, book := range Books {
			if book.Id == id {
				json.NewEncoder(w).Encode(book)
				return

			}
		}
		http.Error(w, "Book not found", http.StatusNotFound)

		// fmt.Fprintf(w, path)

	})

	// new handler for /books only
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Books)
		return
	})

	// lets create and start server
	address := "127.0.0.1:8000"
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println("error: ", err)
	}
}
