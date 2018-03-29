package main

import "fmt"

// MultiArgs ...
func MultiArgs(a, b int, s ...string) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// MultiArgs2 ...
func MultiArgs2(args ...interface{}) {
	for _, v := range args {
		fmt.Println(v)
	}
}

func main() {

	MultiArgs(1, 2, "s1", "s2", "s3")

	MultiArgs2(1, 2, "s1", "s2", "s3", "s4", "s5")
}
