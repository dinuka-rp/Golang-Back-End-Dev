package main

import(
	"fmt"
)

func main() {
	count := 0

	for i := 0000; i <= 9999; i++ {
		for n := i; n <= 9999; n++ { // don't count twice
			q := i * n
			num := fmt.Sprintf("%d", q)

			if num[0] == num[len(num)-1] {
				count++
			}
		}
	}

	fmt.Println(count)
}
