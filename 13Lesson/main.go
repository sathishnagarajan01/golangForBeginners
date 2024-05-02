package main

import "fmt"

func main() {
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
}
