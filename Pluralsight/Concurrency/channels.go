package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	// unbuffered channel
	//ch := make(chan int) // make channel with type of message the channel is going to be working with
	//
	//wg.Add(2)
	//
	//go func(ch chan int, wg *sync.WaitGroup) {
	//	fmt.Println(<-ch) // receive a message from a channel
	//	wg.Done()
	//}(ch, wg)
	//
	//go func(ch chan int, wg *sync.WaitGroup) {
	//	ch <- 42 //	pass a message into the channel
	//	wg.Done()
	//}(ch, wg)

	// buffered channel
	ch := make(chan int, 1) // make channel with type of message the channel is going to be working with
	// , internal capacity of the channel

	wg.Add(2)

	go func(ch <-chan int, wg *sync.WaitGroup) { // receive only channel
		// -- -channels in if statements
		//if msg, ok := <-ch; ok {                 // ok is used to check if channel is open
		//	fmt.Println(msg, ok) // receive a message from a channel
		//}

		// -- -channels in loops
		for msg := range ch { // no explicit range, range of the channel is used
			fmt.Println(msg)
		}

		wg.Done()
	}(ch, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) { // send only channel
		//ch <- 42 //	pass a message into the channel
		//ch <- 27

		// -- -channels in if statements
		//ch <- 0

		// -- -channels in loops
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)	// this is needed to let the for loop know that the channel is over - to avoid deadlock

		wg.Done()
	}(ch, wg)

	wg.Wait()
}
