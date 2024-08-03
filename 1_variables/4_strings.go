package main

import "fmt"

func main() {

	var text string = "Some random text"

	fmt.Println("Text:")
	fmt.Println(text)

	fmt.Println("Length:")
	fmt.Println(len(text))

	fmt.Println("Part of text in array:")
	fmt.Println(text[0:4])

	fmt.Println("Another part of text:")
	fmt.Println(text[4:])

	// COPY STRING VALUE (NOT OBJECT REFERENCE)
	first := "First value"
	second := first
	first += ", Second value"

	fmt.Println(first, second)

}
