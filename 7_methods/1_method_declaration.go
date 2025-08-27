package main

import "fmt"

type Product struct {
    Id int
    Name string
    Price float32
    ProductType
}

type ProductType struct {
    Type string
}

// * pointer to use original object instead of copy
func (p *Product) updatePrice(newPrice float32) {
    p.Price = newPrice
}

func (p *Product) setType(typeName string) {
    p.Type = typeName
}

func main() {

    product := &Product{
        Id: 1,
        Name: "Product 1",
        Price: 10.99,
    }

    fmt.Println("Initial price: ", product.Price)
    product.updatePrice(9.99)
    fmt.Println("Final price: ", product.Price)

    product.setType("some type")
    fmt.Println("Product type: ", product.Type)
}
