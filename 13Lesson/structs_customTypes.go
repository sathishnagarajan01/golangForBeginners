package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	nb := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return nb
}

// receiver functions
// receiver functions with pointer
func (b *bill) format() string {
	fs := "Bill Breakdown: \n"
	var total float64 = 0
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}
	fs += fmt.Sprintf("----------------------------------- \n")
	// add tip
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "Tip:", b.tip)
	// total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "Total:", total+b.tip)
	return fs
}

// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add items
func (b *bill) addItems(name string, price float64) {
	b.items[name] = price
}

// delete items
func (b *bill) deleteItems(key string) {
	delete(b.items, key)
}
