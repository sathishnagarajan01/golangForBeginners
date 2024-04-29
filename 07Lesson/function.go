package main

import (
	"fmt"
	"math"
)

func sayGreeting(name string) {
	fmt.Printf("Good morning %v \n", name)
}
func sayBye(name string) {
	fmt.Printf("Byebye %v \n", name)
}

func cycleName(name []string, callback func(string)) {
	for _, val := range name {
		callback(val)
	}
}

// return type function
func circleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func main() {
	sayGreeting(`sathish`)
	sayBye(`sathish`)

	cycleName([]string{`sathish`, `kumar`, `nagarajan`}, sayGreeting)
	cycleName([]string{`sathish`, `kumar`, `nagarajan`}, sayBye)

	area1 := circleArea(10.5)
	area2 := circleArea(15)

	fmt.Printf("circle 1 is %0.3f \ncircle 2 is %0.3f", area1, area2)
}
