package main

import "os"

func main() {

    panic("Some error")

    _, err := os.Create("/file.txt")

    if err !=nil {
        panic(err)
    }
}
