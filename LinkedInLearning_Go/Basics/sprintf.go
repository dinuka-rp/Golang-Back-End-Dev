package main

import(
	"fmt"
)

func main()  {
	n := 42
	s := fmt.Sprintf("%d", n)	// Converts number to a String

	fmt.Printf("s = %s (type %T)\n", s, s)	// doesn't display the format
	fmt.Printf("s = %q (type %T)\n", s, s)	// displays the format
}