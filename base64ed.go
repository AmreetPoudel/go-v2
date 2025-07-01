// package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	//encrypting data
	data1 := "super secrettttttt is being transfered"
	encrypted := base64.StdEncoding.EncodeToString([]byte(data1))
	fmt.Println(encrypted)

	//decrypting data
	decrypted, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		fmt.Println("error decrypting: ", err)
	}
	fmt.Println(string(decrypted))

	//urlsafe encoding thought differnt but it is different i was like url incode special character only
	//not everything but it is what i donot know
	data2 := "https://amritpoudel.com/getdata?user=amrit&id=1"
	urlsafe := base64.URLEncoding.EncodeToString([]byte(data2))
	fmt.Println(urlsafe)

}
