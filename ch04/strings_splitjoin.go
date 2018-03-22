package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "the quick brown fo jumps over the laze dog"
	s1 := strings.Fields(str)
	fmt.Println(s1)

	for _, val := range s1 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()

	str2 := "go|asdflk ajsd|asjlkf"
	s2 := strings.Split(str2, "|")
	fmt.Println(s2)

	for _, val := range s2 {
		fmt.Printf("%s - ", val)
	}

	fmt.Println()

	str3 := strings.Join(s2, "-")
	fmt.Println(str3)
}
