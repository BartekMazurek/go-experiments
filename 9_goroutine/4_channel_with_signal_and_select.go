package main

import (
	"fmt"
	"time"
)

func worker(stop chan bool) {
	for {
	    // select - similar to switch
	    // based on channel operations instead values
	    // checks all cases simultaneously
	    // if there are several ready, one is selected randomly
	    // blocks and waits if the "default" instruction is missing and no case is ready
		select {
            case <-stop:
                fmt.Println("Finished.")
                return
            default:
                fmt.Println("Processing...")
                time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {

	stop := make(chan bool)

	go worker(stop)

	time.Sleep(2 * time.Second)

    // signal
    // send signal into the channel (no matter whether true/false or other data type)
    // finish processing goroutine with select in for loop
	stop <- true

	time.Sleep(time.Second)
}
