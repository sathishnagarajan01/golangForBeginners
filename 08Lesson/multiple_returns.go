package main

import (
	"fmt"
	"strings"
)

func getStrings(names string) (string, string) {
	s := strings.Split(strings.ToUpper(names), " ")
	if len(s) > 1 {
		return s[0], s[1]
	}
	return s[0], "_"
}

func getInitials(names string) (string, string) {
	s := strings.Split(strings.ToUpper(names), " ")
	var initials []string
	for _, val := range s {
		initials = append(initials, val[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}

func main() {
	res1, res2 := getStrings("sathish kumar")
	res11, res21 := getInitials("kumar nagarajan")
	res12, res22 := getInitials("sathish")
	fmt.Println(res1, res2)
	fmt.Println(res11, res21)
	fmt.Println(res12, res22)
}
