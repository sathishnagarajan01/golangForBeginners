package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}
	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// int as key
	phoneBook := map[int]string{
		123: "mario",
		456: "luigi",
		789: "peach",
	}
	fmt.Println(phoneBook)
	fmt.Println(phoneBook[123])
	phoneBook[456] = "sathish"
	fmt.Println(phoneBook)
}
