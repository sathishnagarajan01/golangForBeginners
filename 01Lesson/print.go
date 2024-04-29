package main

import "fmt"

func main() {
	// print
	fmt.Print("sathish")
	fmt.Print("Kumar \n")
	fmt.Print("New line")

	// Println
	fmt.Println("sathish")
	fmt.Println("kumar")

	// printf (formatted string)
	age := 20
	name := "sathish"
	fmt.Printf("My Age is %v and my name is %v \n", age, name)
	fmt.Printf("My Age is %q and my name is %q \n", age, name)
	fmt.Printf("age is type of %T and name is type of %T \n", age, name)
	fmt.Printf("your score is %f \n", 225.551)
	fmt.Printf("your score is %0.2f \n", 225.551)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("My Age is %v and my name is %v \n", age, name)
	fmt.Println("The saved string is", str)
}
