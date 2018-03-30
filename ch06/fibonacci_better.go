package main

import (
	"fmt"
	"time"
)

// LIMIT ...
const LIMIT = 41

var fibs [LIMIT]uint64

// fibonacci ...
func fibonacci(n int) (ret uint64) {
	if fibs[n] != 0 {
		ret = fibs[n]
		return
	}
	if n <= 1 {
		ret = 1
	} else {
		ret = fibonacci(n-1) + fibonacci(n-2)
	}

	fibs[n] = ret
	return
}

func main() {
	var result uint64

	start := time.Now()

	for i := 0; i < LIMIT; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}

	end := time.Now()

	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}
