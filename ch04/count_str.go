package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, how is it going, Hugo?"
	manyG := "gggggggggg"

	fmt.Printf("Number of H's in %s is: ", str)
	fmt.Printf("%d\n", strings.Count(str, "H"))

	fmt.Printf("Number of double g's in %s is: %d\n", str, strings.Count(manyG, "gg"))
}
