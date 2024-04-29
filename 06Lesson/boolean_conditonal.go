// Boolean & Conditionals

package main

import "fmt"

func main() {
	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age <= 45)
	fmt.Println(age != 50)

	if age > 30 {
		fmt.Println("age is less than 30")
	} else if age > 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less than 45")
	}

	fmt.Println("\n loop range")
	names := []string{"sathish", "kumar", "nagarajan", "saraswathi", "dinesh"}
	for index, value := range names {
		if index == 1 {
			fmt.Println(`continue at position`, index)
			continue
		}
		if index > 2 {
			fmt.Println(`break at position`, index)
			break
		}
		fmt.Printf("value of name is: %v and index is: %d \n", value, index)
	}
}
