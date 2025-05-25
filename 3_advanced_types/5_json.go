package main

import "fmt"
import "encoding/json"

type Product struct{
    Id      int         `json:"Id"`
    Name    string      `json:"Name"`
    Price   float32     `json:"Price"`
}

type Order struct {
    Id          int         `json:"Id"`
    Products    []Product   `json:"Products"`
}

func main() {

    product1 := Product{
        Id: 1,
        Name: "product1",
        Price: 9.99,
    }

    product2 := Product{
        Id: 2,
        Name: "product2",
        Price: 11.99,
    }

    product3 := Product{
        Id: 3,
        Name: "product3",
        Price: 5.99,
    }

    products := []Product{product1, product2, product3}

    order := &Order{
        Id: 1,
        Products: products,
    }

    // Serialize data
    jsonOrder, err := json.Marshal(order)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(string(jsonOrder))

    // API call ...

    // Deserialize data
    var deserializedOrder Order
    err2 := json.Unmarshal(jsonOrder, &deserializedOrder)

    if err2 != nil {
        fmt.Println(err2)
    }

    fmt.Println(deserializedOrder)
}
