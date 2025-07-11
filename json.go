// package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name       string   `json:"name"`
	LoggedIn   bool     `json:"loggedin"`
	Roles      []string `json:"roles"`
	LoginCount int      `json:"logincount"`
}

func main() {

	user := User{
		Name:       "amrit poudel",
		LoggedIn:   true,
		Roles:      []string{"superadmin", "infraadmin"},
		LoginCount: 2,
	}
	datab, err := json.Marshal(user)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(string(datab))

	//lets decode and see original
	var ori User
	err = json.Unmarshal(datab, &ori)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ori: ", ori)

}
