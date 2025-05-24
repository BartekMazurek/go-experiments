package main

import "fmt"

func main() {

    // A declared length array cannot be extended
    // Declare array of integers with 4 indexes
    var array1[4]int

    fmt.Println(array1)

    array1[0] = 11
    array1[1] = 22
    array1[2] = 33
    array1[3] = 44

    fmt.Println(array1)

    // Define array of integers with 4 indexes & assign 3 of them
    var array2 = [4]int{11, 22, 33}
    fmt.Println(array2)

    // Define array of strings with 3 indexes
    var array3 = [3]string{
        "Hey",
        "Hi",
        "Hello",
    }
    fmt.Println(array3)

    // Define array of strings with undefined amount of elements
    var array4 = [...]string{
        "Another",
        "Hey",
        "Hi",
        "Hello",
    }
    fmt.Println(array4)

    // Define a two-element multidimensional array with three elements in each array
    var array5 = [2][3]int {
        [3]int{1,3,4},
        [3]int{5,6,7},
    }
    fmt.Println(array5)

    // Define multidimensional array with two elements in each array
    var array6 = [...][2]int{{1,2},{3,4},{5,6},{7,8}}
    fmt.Println(array6)
}