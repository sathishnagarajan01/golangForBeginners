package main

import "fmt"

func main() {
	x := 0
	for x < 5 {
		fmt.Println("value is: ", x)
		x++
	}

	fmt.Println("\ntraditional")
	for i := 0; i < 5; i++ {
		fmt.Println("value is: ", i)
	}

	fmt.Println("\nvariable loop")
	names := []string{"sathish", "kumar", "nagarajan", "saraswathi", "dinesh"}
	for i := 0; i < len(names); i++ {
		fmt.Println("name is: ", names[i])
	}

	fmt.Println("\nloop like for in, using range")
	for index, value := range names {
		fmt.Printf("value of name is: %v and index is: %d \n", value, index)
	}

	fmt.Println("\nloop like for in, using range wihout index")
	for _, value := range names {
		fmt.Printf("value of name is: %v \n", value)
	}
}
