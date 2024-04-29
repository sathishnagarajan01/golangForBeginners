package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	randomStr := "Hello world program!"

	fmt.Println(strings.Contains(randomStr, "Hello"))
	fmt.Println(strings.ReplaceAll(randomStr, "Hello", "Hai"))
	fmt.Println(strings.ToUpper(randomStr))
	fmt.Println(strings.Index(randomStr, "llo"))
	fmt.Println(strings.Split(randomStr, " "))

	// the original value is unchanged
	fmt.Println("original str:= ", randomStr)

	ages := []int{45, 40, 20, 25, 66, 7, 3, 21, 20, 88}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 20)
	fmt.Println(index)

	names := []string{"sathish", "kumar", "nagarajan", "saraswathi", "dinesh"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "sathish"))

}
