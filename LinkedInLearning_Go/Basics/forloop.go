package main

import (
	"fmt"
)

func main()  {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("----------")
	for i := 0; i < 3; i++ {
		if i > 1{
			break
		}
		fmt.Println(i)
	}

	fmt.Println("----------")
	for i := 0; i < 3; i++ {
		if i < 1{
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("----------")
	a := 0
	for a < 3{			//this is like a while loop in other languages
		fmt.Println(a)
		a++
	}

	fmt.Println("----------")
	b := 0
	for {			//this is like a while true loop in other languages
		if b > 2{
			break
		}
		fmt.Println(b)
		b++
	}
}