// package main

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {

	password := "super secrettttt"

	// data := "data to be encrypted super secrettt"
	// enc := sha512.Sum512([]byte(data))
	// fmt.Printf("here is encrypted text,%x", enc)
	salt, error := GenerateSalt()
	if error != nil {
		fmt.Printf("error generating salt: %s", error)
	}
	SignUPpassHash := HashPassword(password, salt)
	// now store salt and hashed password  in db simulate print as db store
	// following printed values will be stored in db donot encrypt salt because you need actual salt to append
	//to user input value then  hash that and check with db cred so it not that risk , putting salt normally
	// no because it is generated randomly and it get appended to password then hashed so attacker has no way
	// to get password if user table is ocmpromised too
	fmt.Printf("passwordHash: %s\n", SignUPpassHash)
	fmt.Printf("saltHash: %x\n", salt)

	// now lets simulate login user put this password in login
	userpass := "super secretttt"
	//taking salt from db
	usrsalt := salt
	loginHash := HashPassword(userpass, usrsalt)
	if SignUPpassHash == loginHash {
		fmt.Println("loggedin")
	} else {
		fmt.Println("invalid cred")
	}

}

func GenerateSalt() ([]byte, error) {
	//lets create empty salt
	salt := make([]byte, 16)
	// generate salt value
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil

}

func HashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha512.Sum512(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}
