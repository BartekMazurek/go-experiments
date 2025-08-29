package main

import "fmt"

func getOperations(ch chan string) {

    operations := []string{"operation1", "operation2", "operation3"}

    for _, o := range operations {
        fmt.Println("Sending",o,"to the channel")
        ch <- o // blocks until something takes that "o" value
    }
    close(ch)
    fmt.Println("Channel closed")
}

func main() {

	ch := make(chan string)

	go getOperations(ch)

    // each "send" waits for "receive"
	// works until channel is not closed
	for operation := range ch { // blocks until something is sent into the channel
		fmt.Println("Received", operation,"from the channel")
	}

    fmt.Println("Finished")
}
