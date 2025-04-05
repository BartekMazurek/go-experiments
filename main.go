package main

import (
	"fmt"
	"go-experiments/2_packages"
)

func main() {

    // 2_packages
	client1 := client.Client{}
	client1.Firstname = "Firstname"
	client1.Lastname = "Lastname"

	fmt.Println(client1.Firstname, " ", client1.Lastname)

}
