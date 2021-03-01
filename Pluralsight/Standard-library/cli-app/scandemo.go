package main

import "fmt"

func main() {
	var name string
	fmt.Println("What's your name")
	//inputs, _ := fmt.Scanf("%s", &name) // input from scanf separated by space -> prints each value with the new-line character
	inputs, _ := fmt.Scanf("%q", &name) // input has to be within ""

	switch inputs {
	case 0:
		fmt.Printf("You must enter a name!\n")
	case 1:
		fmt.Printf("Hello %s! You input %d input value\n", name, inputs)
	}

}
