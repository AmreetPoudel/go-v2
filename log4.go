// package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("error opening file app.log")
	}
	warnlogger := log.New(file, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	inforlogger := log.New(file, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorlogger := log.New(file, "error: ", log.Ldate|log.Ltime|log.Lshortfile)

	//lets simulate
	//warning logs to file so that can retrive later on
	warnlogger.Println("this is warning log")
	inforlogger.Println("this is infor log")
	errorlogger.Println("this is error log")

}
