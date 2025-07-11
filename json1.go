// package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Username   string   `json:"username"`
	LoggedIn   bool     `json:"loggedIn"`
	Roles      []string `json:"roles"`
	LoginCount int      `json:"loginCount"`
}

func main() {
	user := User{
		Username:   "amrit",
		LoggedIn:   true,
		Roles:      []string{"admin", "user"},
		LoginCount: 5,
	}

	// Convert struct → JSON (Marshal)
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonBytes))

	// Convert JSON → struct (Unmarshal)
	var decoded User
	json.Unmarshal(jsonBytes, &decoded)
	fmt.Printf("Decoded struct: %+v\n", decoded)
}
