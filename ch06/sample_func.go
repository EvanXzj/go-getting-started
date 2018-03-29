package main

import "fmt"

func main() {

	var c *int
	a := 10

	c = &a

	A(c)

	b := &a
	fmt.Println(b)
	B(a)

}

// A ..
func A(a *int) {
	b := a

	fmt.Println(b)
}

// B ...
func B(a int) {
	b := &a

	fmt.Println(b)
}
