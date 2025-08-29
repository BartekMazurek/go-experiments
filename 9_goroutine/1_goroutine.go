package main

import (
	"fmt"
	"time"
)

func handle(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {

	go handle("operation 1")
	go handle("operation 2")

	time.Sleep(1 * time.Second)
}
