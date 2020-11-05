package main

import (
	"fmt"
	"sync"
	"time"
)
func main(){

	var waitGrp sync.WaitGroup
	waitGrp.Add(2)		// number of Go Routines expected to end before exiting

	go func(){	// go routine (concurrent function)
		defer waitGrp.Done()

		time.Sleep(5*time.Second)
		fmt.Println("Hello")
	}()		// self executing function

	go func(){	// go routine (concurrent function)
		defer waitGrp.Done()

		fmt.Println("Pluralsight")
	}()

	waitGrp.Wait()
}
