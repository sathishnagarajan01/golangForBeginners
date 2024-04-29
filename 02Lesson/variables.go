package main

import "fmt"

// this will give you error
// syntax error: non-declaration statement outside function body
// variable5 := "sathish"

func main() {

	// string
	var variable1 string = "sathish"
	// eventhough dataType not declare based on initial value go consider the data type
	var variable2 string = "kumar"
	// empty string declaration
	var variable3 string
	// withour var keyword, this method can't able to use outside func as global, see varible5
	variable4 := "sathish"
	fmt.Println(variable1, variable2, variable3, variable4)

	// integer
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40
	fmt.Println(ageOne, ageTwo, ageThree)

	// integer bits memory, 8,16,32,64
	var numOne int8 = 127
	var numTwo uint8 = 255 // unsigned support only possitve number
	var numThree int16 = 25867
	fmt.Println(numOne, numTwo, numThree)

	// float

	var scoreOne float32 = 25.32
	var scoreTwo float64 = 1234567890.1234567890
	fmt.Println(scoreOne, scoreTwo)
}
