package main

import "fmt"

func main() {
	fmt.Println("Hello")

	// strings
	var nameOne string = "Raphael"
	var nameTwo = "Barreto"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Guilherme"
	nameThree = "Bowser"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Mario"

	fmt.Println(nameFour)

	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	var numOne int8 = 25 
	var numTwo int8 = -128
	var numThree uint = 255

	fmt.Println(numOne, numTwo, numThree)

	// float
	var scoreOne float32 = 44.7
	var scoreTwo float64 = 84897897489489444.7
	scoreThere := 1.5

	fmt.Println(scoreOne, scoreTwo, scoreThere)
}
