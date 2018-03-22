package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "This is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s ?\n", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))

	fmt.Printf("T/F? Does the string \"%s\" have suffix %s ?\n", str, "ing")
	fmt.Printf("%t\n", strings.HasSuffix(str, "ing"))
}
