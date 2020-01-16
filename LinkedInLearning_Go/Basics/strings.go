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
	
	// Slice (start, end)
	fmt.Println(book[4:11])

	// Slice (no end)
	fmt.Println(book[4:])

	// Slice (no start)
	fmt.Println(book[:4])

	// String concatenation
	fmt.Println("t"+book[1:])

	// Unicode support
	fmt.Println("☆")
	
	//Multi line strings using back tick ``
	poem := `
	The road goes ever on
	Down from the door where it began
	...
	`
	fmt.Println(poem)
}