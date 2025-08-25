package main

import "fmt"

var multiply = func(a,b int) int {
    return a * b
}

func main() {

    // 1
    fmt.Println("Result:", multiply(66, 99))

    // 2
    sum := func(a int, b int) int {
        return a + b
    }(66 ,99)

    fmt.Println("Result:", sum)
}
