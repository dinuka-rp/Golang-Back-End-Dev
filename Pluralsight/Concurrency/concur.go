package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)
// Concurrency - composition of independently carrying out tasks
//Parallelism - making 2 go routines to be able to run simultaneously on separate threads/ cors
func main(){
	runtime.GOMAXPROCS(2)	// no. of virtual processors/ threads

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
