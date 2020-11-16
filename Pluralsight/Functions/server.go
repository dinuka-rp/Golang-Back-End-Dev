package main

import (
	"errors"
	"fmt"
)

func main(){
	fmt.Println("Hello, playground")
	port := 3000
	err := startWebServer(port, 2)
	fmt.Println(err)
}

func startWebServer(port, numberOfRetries int) error {		// data type followed after multiple commas will be applied to all vars before it
	fmt.Println("Start server...")

	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)

	//return nil
	return errors.New("something went wrong")
}
