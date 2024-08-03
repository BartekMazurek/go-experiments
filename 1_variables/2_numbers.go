package main

import "fmt"

func main() {

	// UINT - unsigned integer
	var number uint8 = 10     // RANGE 0 - 255
	var number2 uint16 = 5    // RANGE 0 - 65535
	var number3 uint32 = 10   // RANGE 0 - 4,294,967,295
	var number4 uint64 = 1000 // RANGE 0 - 18,446,744,073,709,551,615

	fmt.Println("Unsigned integers")
	fmt.Println(number, number2, number3, number4)

	// INT
	var number5 int8 = -2         // RANGE -128 - 127
	var number6 int16 = 12000     // RANGE -32,768 - 32,767
	var number7 int32 = -99123123 // RANGE -2,147,483,648 - 2,147,483,647
	var number8 int64 = 1000      // RANGE -9,223,372,036,854,775,808 - 9,223,372,036,854,775,807

	fmt.Println("Signed integers")
	fmt.Println(number5, number6, number7, number8)

	// FLOAT
	var float1 float32 = 100.44    // STANDARD IEEE-754
	var float2 float64 = 5.1231231 // STANDARD IEEE-754

	fmt.Println("Floats")
	fmt.Println(float1, float2)

	// COMPLEX
	var x complex128 = complex(1, 2)

	fmt.Println("Complex")
	fmt.Println(x)

}
