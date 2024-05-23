package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// this is manually add items
/*func main() {
	myBill := newBill("sathish's bill")
	// fmt.Println(myBill)

	myBill.updateTip(10.0)
	myBill.addItems("pie", 5.99)
	myBill.addItems("cake", 3.99)

	myBill.addItems("sathish", 0.99)
	myBill.deleteItems("sathish")

	fmt.Println(myBill.items)
	fmt.Println(myBill.format())

	// "pie": 5.99, "cake": 3.99
}*/

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("choose Option (a - add item, s - save bill, t - add tip, d - delete item): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Enter item name: ", reader)
		price, _ := getInput("Enter item price: ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Price must be a number")
			promptOptions(b)
		}
		b.addItems(name, p)
		fmt.Println("Item added - ", name, price)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("saved the bill - ", b.name)
	case "t":
		tip, _ := getInput("Enter tip amount: ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("Tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("tip added - ", tip)
		promptOptions(b)
	case "d":
		del, _ := getInput("Enter name of item: ", reader)
		if b.deleteItems(del) {
			fmt.Println("item deleted: ", del)
		} else {
			fmt.Println("You entered wrong item")
		}
		promptOptions(b)
	default:
		fmt.Println("not a valid options")
		promptOptions(b)
	}
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	/*fmt.Print("create a new bill name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)*/
	name, _ := getInput("create a new bill name: ", reader)

	nb := newBill(name)

	return nb
}

// this is getting input from user
func main() {
	myBill := createBill()
	promptOptions(myBill)
	fmt.Println(myBill)
}
