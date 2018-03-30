package main

import (
	"fmt"
	"time"
)

// fibonacci ...
func fibonacci(n int) (ret int) {
	if n <= 1 {
		ret = 1
	} else {
		ret = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

func main() {
	result := 0

	start := time.Now()

	for i := 0; i <= 40; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}

	end := time.Now()

	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}
