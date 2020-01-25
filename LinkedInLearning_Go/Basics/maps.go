// Exercise on maps

package main

import(
	"fmt"
)

func main()  {

	stocks := map[string]float64{
		"AMZN": 1699.8,
		"GOOG": 1129.19,
		"MSFT": 98.61,		// Must have trailing comma in multi line
	}

	// Number of items
	fmt.Println(len(stocks))

	// Get a value
	fmt.Println(stocks["MSFT"])

	// Get zero value if not found
	fmt.Println(stocks["TSLA"])
	
	// Use two values to see if found
	value, ok := stocks["TSLA"]
	// ok takes true or false
	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(value)
	}

	// Set value
	stocks["TSLA"] = 322.12
	fmt.Println(stocks)

	// Set value
	stocks["TSLA"] = 50
	fmt.Println(stocks)

	// Delete value
	delete(stocks, "AMZN")
	fmt.Println(stocks)

	// Single value "for" is on keys
	for key := range stocks{
		fmt.Println(key)
	}

	// Double value "for" is key, value
	for key, value := range stocks{
		fmt.Printf("%s -> %.2f\n", key , value)
	}

}