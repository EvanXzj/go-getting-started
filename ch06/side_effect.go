package main

import "fmt"

// Multiply ...
func Multiply(a, b int, reply *int) {
	*reply = a * b
}

func main() {

	n := 0

	reply := &n
	fmt.Println(reply)

	Multiply(10, 5, reply)

	fmt.Println("Multiply:", *reply) // Multiply: 50
	fmt.Println("Multiply:", n)      // Multiply: 50
}
