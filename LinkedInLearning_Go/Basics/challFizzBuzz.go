// Challenge: FizzBuzz

package main

import(
	"fmt"
)

func main()  {
	for x := 1; x <=20; x++{
		if x%3==0 && x%5==0 {
			fmt.Println("fizz buzz")
		} else if x%3==0 {
			fmt.Println("fizz")
		}else if x%5==0 {
			fmt.Println("buzz")
		} else{
			fmt.Println(x)
		}
	}
}