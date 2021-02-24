package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//fmt.Printf("We are using Go %v running in %v\n " , runtime.Version(), runtime.GOOS)

	reader := bufio.NewReader(os.Stdin)		// reading from OS standard input = terminal

	fmt.Println("What's your name?")
	text, _ := reader.ReadString('\n')	// take input onEnter (new line)
	fmt.Printf("Hello %v", text)		// the return character is kept at the end of text

}
