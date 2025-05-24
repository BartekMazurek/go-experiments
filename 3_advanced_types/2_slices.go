package main

import "fmt"

func main() {

    // Slices - represents fragment of array
    // Contains a pointer to the elements, length and capacity
    // Can be easily extended, cut, transferred
    array1 := [...]string{"test1", "test2", "test3", "test4", "test5"}

    // Get array1 slice with elements from 0 to 2
    slice1 := array1[:3]
    fmt.Println(slice1)

    // Get array1 slice with elements from 3 to 4
    slice2 := array1[3:5]
    fmt.Println(slice2)

    // Create slice with new element
    slice3 := append(slice2, "test6")
    fmt.Println(slice3)
}
