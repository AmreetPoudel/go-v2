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
	Profile    Profile  `json:"profile"` // 👈 nested struct
}

type Profile struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Location string `json:"location"`
}

func main() {
	user := User{
		Username:   "amrit",
		LoggedIn:   true,
		Roles:      []string{"admin", "user"},
		LoginCount: 17,
		Profile: Profile{
			FullName: "Amrit Poudel",
			Email:    "amrit@example.com",
			Location: "Lalitpur, Nepal",
		},
	}

	jsonBytes, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonBytes))
}
