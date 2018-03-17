// Go 程序的一般结构
// 1. 定义包名，导入包，定义常量、变量、类型或声明
// 2. 如果存在init 函数，就先定义改函数。 init 是一个特殊函数，含有该函数的包都会先执行这个函数
// 3. 如果当前包是main， 则定义main 函数
// 4. 然后定义其他函数，首先是类型的方法，然后按照mian函数中调用的顺序来定义相关函数

package main

import (
	"fmt"
)

const c = "C"

var v = 5

// T ...
type T struct{}

func main() {
	a := 10
	Func1()
	fmt.Println()
}

// Method1 ...
func (t T) Method1() {
	// code here ...
}

// Func1 ...
func Func1() {
	// code here ...
}
