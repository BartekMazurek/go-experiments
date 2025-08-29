package main

import "fmt"

func getOperations(ch chan string) {

    operations := []string{"operation1", "operation2", "operation3", "operation4", "operation5", "operation6"}

    for _, o := range operations {
        fmt.Println("Sending",o,"to the channel")
        ch <- o // sending blocked when there are already 3 items in the channel
    }
    close(ch)
    fmt.Println("Channel closed")
}

func main() {

	ch := make(chan string, 3) // created buffered channel with size of 3 elements with string type

	go getOperations(ch)

    // blocks only when len(ch)==0 and the channel is not closed yet
    // processes all remaining elements even if the channel has been closed
    // for range will stop when the channel is empty and closed
	for operation := range ch {
		fmt.Println("Received", operation,"from the channel")
	}

    fmt.Println("Finished")
}
