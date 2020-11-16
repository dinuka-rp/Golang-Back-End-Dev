package main

import "fmt"

func main(){
	fmt.Println("Hello, playground")
	port := 3000
	isStarted := startWebServer(port, 2)
	fmt.Println(isStarted)
}

func startWebServer(port, numberOfRetries int) bool {		// data type followed after multiple commas will be applied to all vars before it
	fmt.Println("Start server...")

	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)

	return true
}
