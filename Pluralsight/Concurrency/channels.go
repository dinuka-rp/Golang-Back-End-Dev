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

	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch) // receive a message from a channel
		wg.Done()
	}(ch, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 42 //	pass a message into the channel
		ch <- 27
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
