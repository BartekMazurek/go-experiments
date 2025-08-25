package main

import "fmt"

func main() {

    fmt.Println("Set first parameter:")
    var parameter1 string
    fmt.Scanln(&parameter1)

    fmt.Println("Set second parameter:")
    var parameter2 string
    fmt.Scanln(&parameter2)

    fmt.Println("Parameter 1:", parameter1, "Parameter 2:", parameter2)
}
