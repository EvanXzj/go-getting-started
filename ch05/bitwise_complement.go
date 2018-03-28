package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("the complement of %b is: %b\n", i, ^i)
	}

	for i := 0; i <= 10; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}

	i := 0
	for {
		if i >= 3 {
			break
		}

		fmt.Println(i)
		i++
	}

	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
