package main

import "fmt"

func main() {

    // Maps - store key-value pairs where each key is unique

    // Declare map with keys of type integer and values of type string
    map1 := make(map[int]string)

    map1[1] = "Customer1"
    map1[2] = "Customer2"
    map1[3] = "Customer3"

    fmt.Println(map1)

    // Define map with keys of type string and values of type integer
    map2 := map[string]int{
        "Product1" : 1,
        "Product2" : 2,
        "Product3" : 3,
    }
    fmt.Println(map2)

    // Remove element form map1 with id = 2
    delete(map1, 2)
    fmt.Println(map1)
}
