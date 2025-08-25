package main

import "fmt"

func doSomething() {
    fmt.Println("Doing something ...")
}

func doSomething2(operation string) {
    fmt.Println("Doing", operation)
}

func calculate(a int, b int) int64 {
    return int64(a * b)
}

func calculateAgain(a,b int)(int, int) {
    return a + b, a * b
}

func namedArguments(a,b int)(add int, multiplication int) {
    add = a + b
    multiplication = a * b
    return
}

func main() {
    //doSomething()

    //doSomething2("operation")

    //result := calculate(66, 99)
    //fmt.Println("Result:", result)

    //result1, result2 := calculateAgain(66, 99)
    //fmt.Println("Result 1:", result1, "Result 2:", result2)

    add, multiplication := namedArguments(66, 99)
    fmt.Println("Add:", add, "Multiplication:", multiplication)

}
