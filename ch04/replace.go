package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "My Name is Chui Chui"

	fmt.Println(str)
	fmt.Println("After Replace:")

	//n 为替换个数， 当n= -1 时，替换所有的匹配的字符
	str = strings.Replace(str, "Chui", "Evan", 0)
	fmt.Println(str)

	str = strings.Replace(str, "Chui", "Evan", 1)
	fmt.Println(str)

	str = strings.Replace(str, "Chui", "Evan", 2)
	fmt.Println(str)

	str = strings.Replace(str, "Chui", "Evan", 3)
	fmt.Println(str)

	str = strings.Replace(str, "Chui", "Evan", -1)
	fmt.Println(str)
}
