package main

import (
	"fmt"
	"math/rand"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano())) // random number generator

func main() {
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1 // book IDs 1-10

		go func(id int) {
			if b, ok := queryCache(id); ok {
				//	if ok is true
				fmt.Println("from cache")
				fmt.Println(b)
			}
		}(id)

		go func(id int) {
			if b, ok := queryDatabase(id); ok {
				fmt.Println("from database")
				fmt.Println(b)
			}
		}(id)

		//fmt.Printf("Book not found with id: '%v'", id)
		time.Sleep(150 * time.Millisecond)		// used to warmup the cache
	}
	time.Sleep(2 * time.Millisecond)		// to make sure that all the go routines are completed
}

func queryCache(id int) (Book, bool) {
	b, ok := cache[id] // ok= true/ false if value is received/ not
	return b, ok
}

func queryDatabase(id int) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			cache[id] = b // update cache with newly found book
			return b, true
		}
	}

	return Book{}, false
}
