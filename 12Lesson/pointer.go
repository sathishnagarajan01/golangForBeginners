package main

import "fmt"

func updateName(x *string) {
	*x = "kumar"
}

func main() {
	name := "sathish"
	fmt.Println(name)

	fmt.Println("address of variable name is: ", &name)

	updateName(&name)
	fmt.Println(name)
}
