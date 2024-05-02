package main

import "fmt"

// pass by value
func updateName(x string) {
	x = "kumar"
}

func updateNameReturn(x string) string {
	x = "kumar"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {
	// Group A type
	name := "sathish"

	// this will not change the value of name
	updateName(name)
	name = updateNameReturn(name)
	fmt.Println(name)

	// Group B Type map, slice, functions
	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}
	updateMenu(menu)
	fmt.Println(menu)
}
