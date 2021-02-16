package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano())) // random number generator

func main() {
	wg := &sync.WaitGroup{} // & => pointer to the waitgroup
	m := &sync.RWMutex{}

	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1 // book IDs 1-10
		wg.Add(2)              // how many concurrent tasks that the wait group will wait for
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := queryCache(id, m); ok {
				//	if ok is true
				fmt.Println("from cache")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m)

		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := queryDatabase(id, m); ok {
				fmt.Println("from database")
				fmt.Println(b)
			}
			wg.Done()

		}(id, wg, m)

		//fmt.Printf("Book not found with id: '%v'", id)
		time.Sleep(150 * time.Millisecond) // used to warmup the cache
	}
	wg.Wait()
	time.Sleep(2 * time.Millisecond) // to make sure that all the go routines are completed
}

func queryCache(id int, m *sync.RWMutex) (Book, bool) {
	m.RLock()		// multiple readers
	b, ok := cache[id] // ok= true/ false if value is received/ not
	m.RUnlock()
	return b, ok
}

func queryDatabase(id int, m *sync.RWMutex) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			m.Lock()		// no readers - only one writer
			cache[id] = b // update cache with newly found book
			m.Unlock()
			return b, true
		}
	}

	return Book{}, false
}
