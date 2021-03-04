package main

import "fmt"

func main() {

	fmt.Printf("|%4.2f|%4.2f|%4.2f|\n", 23.3774, 577.45, 1234.56)
	fmt.Printf("|%4.2f|%4.2f|%4.2f|\n", 98.999, 12.3456, 12.01)

	//	spreadsheet view - right aligned
	fmt.Printf("|%7.2f|%7.2f|%7.2f|\n", 23.3774, 577.45, 1234.56)
	fmt.Printf("|%7.2f|%7.2f|%7.2f|\n", 98.999, 12.3456, 12.01)

	//	spreadsheet view - left aligned
	fmt.Printf("|%-7.2f|%-7.2f|%-7.2f|\n", 23.3774, 577.45, 1234.56)
	fmt.Printf("|%-7.2f|%-7.2f|%-7.2f|\n", 98.999, 12.3456, 12.01)

	//	spreadsheet view - left aligned - strings & quoted text
	fmt.Printf("|%-7.2s|%-7.2s|%-7.2q|\n", "foo", "bar", "go")
}
