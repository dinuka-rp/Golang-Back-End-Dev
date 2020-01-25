// Example of slices | similar to arrays | built on top of arrays

package main

import(
	"fmt"
)

func main(){
	// Slices have to be of the same type
	loons:= []string{"bugs", "daffy", "taz"}
	fmt.Printf("loons = %v (type %T)\n", loons, loons)

	//length
	fmt.Println(len(loons))

	fmt.Println("----")
	// index starts from 0
	fmt.Println(loons[1])

	fmt.Println("----")
	// slices
	fmt.Println(loons[1:])

	fmt.Println("----")
	for i := 0; i < len(loons); i++{
		fmt.Println(loons[i])
	}

	fmt.Println("----")
	for i := range loons{
		fmt.Println(i)
	}

	fmt.Println("----")
	for i, name := range loons{
		fmt.Printf("%s at %d\n", name, i)
	}

	fmt.Println("----")
	for _, name := range loons{			// _ is an unused variable
		fmt.Println(name)
	}

	fmt.Println("----")
	// extending the slice
	loons = append(loons, "elmer")
	fmt.Println(loons)


}