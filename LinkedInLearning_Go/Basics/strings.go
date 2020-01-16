package main

import(
	"fmt"
)

func main()  {
	book := "The colour of magic"
	fmt.Println(book)

	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])
	// unit8 is a byte

	// Strings in go are immutable
	// book[0] = 116		// gives an error
	
}