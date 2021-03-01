package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// keyboard input
	//scanner := bufio.NewScanner(os.Stdin)
	//for scanner.Scan() { // infinite loop
	//	if scanner.Text() == "/quit" { // end program if /quit is input
	//		fmt.Println("Quitting")
	//		os.Exit(3) // exit program with status 3
	//	} else {
	//		fmt.Println("You typed " + scanner.Text())
	//	}
	//}

	// file input
	file, err := os.Open("./test-files/test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) // this scanner will close when there're no more lines in the file
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err();
		err != nil {
		fmt.Println(err)
	}

}
