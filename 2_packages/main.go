package main

import (
	"fmt"

	"./client"
)

func main() {

	client1 := client.Client{}
	client1.Firstname = "Firstname"
	client1.Lastname = "Lastname"

	fmt.Println(client1.Firstname, " ", client1.Lastname)

}
