package main

import "fmt"

// If variable/function name begins with capital, then it is public
var someBoolean = false

// Declared variables are required to be used in order to save memory.

func main() {

	// Various ways of defining variables
	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree, someBoolean)

	nameOne = "peach"
	nameThree = "bowser"

	fmt.Print(nameOne, nameTwo, nameThree)

	// the following is allowed inside functions only
	nameFour := "yoshi"
	fmt.Print(nameFour)

	// int variables
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Print(ageOne, ageTwo, ageThree)

	// bits & memory
	// var numOne int8 = 25
	// var numTwo int8 = 128 // too large a number for 8-bit
	// var numTwo uint = -25 unsigned ints cannot be negative

	var scoreOne float32 = 25.98
	var scoreTwo float64 = 1965385877.5
	var scoreThree = 1.5 // inferred as float64

	fmt.Print(scoreOne, scoreTwo, scoreThree)

	// for more info see https://golang.org/ref/spec#Numeric_types

}
