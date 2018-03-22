package main

import (
	"fmt"
	"strings"
)

func main() {
	oStr := "Hi There! "
	var newStr string

	newStr = strings.Repeat(oStr, 3)
	fmt.Printf("The new repeated string is %s\n", newStr)
}
