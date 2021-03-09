package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	// create a file if doesn't exist, append if exists, write only, file permissions ,
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println("This is a log message")
}
