package main

import "fmt"

func main() {

    // 1
    condition := true

    if condition {
        fmt.Println("Condition ")
    }

    // 2
    number := 6

    if number > 5 {
        fmt.Println("Number > 5")
    } else if number > 2  {
        fmt.Println("Number > 2")
    } else {
        fmt.Println("Unknown condition")
    }

    // 3
    switch number {
        case 5:
            fmt.Println("5")
        default:
            fmt.Println("Unknown")
    }

    // 4
    character1 := "b"
    switch character1 {

        case "a", "b", "c", "d":
            fmt.Println("Character: " + character1)
        default:
            fmt.Println("Unknown")
    }

    // 5
    switch character2 := "f"; character2 {
        case "e", "f", "g", "h":
            fmt.Println("Character: " + character2)
        default:
            fmt.Println("Unknown")
    }

    // 6
    fallthrough_case := 10
    switch fallthrough_case {
        case 10:
            fmt.Println(fallthrough_case)
            fallthrough
        case 9:
            fmt.Println(fallthrough_case)
        case 8:
        case 7:
        case 6:
        case 5:
        case 4:
        case 3:
        case 2:
        case 1:
            fmt.Println(fallthrough_case)
        default:
            fmt.Println("Wrong value")
    }

    // LOOPS

    // 7
    // INFINITE LOOP
//     for i := 1; i < 6; {
//         fmt.Printf("Loop %d", i)
//     }

    // 8
    j := 1
    for j < 5 {
        fmt.Printf("Loop %d ", j)
        j++
    }
    fmt.Println("Current value: ", j)

    // 9
    k := 1
    for {
        if k == 5 {
            break
        }
        fmt.Println("Value: ", k)
        k++
    }
}
