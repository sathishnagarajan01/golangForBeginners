package main

import (
	"fmt"
	"reflect"
)

func main() {
	// array
	var ages [3]int = [3]int{1, 2, 3}
	var agesTwo = [3]int{1, 2, 3}
	fmt.Println(ages, agesTwo)

	name := [4]string{"sathish", "kumar", "nagarajan", "testing"}
	fmt.Println(name, len(name))

	// slice
	// array without static length
	var scores = []int{20, 30, 40}
	//update to specific index
	scores[2] = 25
	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

	// slice range
	rangeOne := name[1:3]
	rangeTwo := name[2:]
	rangeThree := name[:3]
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "dummy")
	fmt.Println(rangeOne, "\n")
	fmt.Println(reflect.ValueOf(name).Kind())

	x := [5]int{1: 10, 3: 30}
	fmt.Println(x)

	y := []interface{}{1, 2, "sathish", 1.112}
	fmt.Println(reflect.ValueOf(y).Kind())

	i := y[0]
	fmt.Printf("i: %d, i type: %T\n", i, i)

	s := y[2]
	fmt.Printf("b: %s, i type: %T\n", s, s)
}
