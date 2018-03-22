package main

import (
	"fmt"
	"strconv"
)

func main() {
	var oring = "666"
	var an int
	var newS string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	an, _ = strconv.Atoi(oring)
	fmt.Println(an)

	an += 5

	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}
