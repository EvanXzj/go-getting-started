package main

import "strings"
import "fmt"

func main() {
	str := "Hi, I'm Chui, Hi."

	fmt.Printf("The position of \"Chui\" is :")
	fmt.Printf("%d\n", strings.Index(str, "Chui"))

	fmt.Printf("The position of first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Hi"))
	fmt.Printf("The position of last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))

	fmt.Printf("The position of \"Nmae\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Name"))
}
