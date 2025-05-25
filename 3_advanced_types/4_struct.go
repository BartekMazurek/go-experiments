package main

import "fmt"

type Product struct{
    Id int
    Name string
    Price float32
}

type Order struct {
    Id int
    Products []Product
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

    fmt.Println(products)

    order := &Order{
        Id: 1,
        Products: products,
    }

    fmt.Println(order)
}
