package main

import "fmt"

func doSomething() {
    fmt.Println("Doing something ...")
}

func main() {

    fmt.Println("Starting ...")

    defer doSomething()

    fmt.Println("Ending ...")
}
