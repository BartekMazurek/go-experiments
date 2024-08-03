package main

import "fmt"

func main() {

	const number = 3000

	type TestType int

	const (
		first TestType = iota
		second
		third
	)

	fmt.Println(first, second, third)

}
