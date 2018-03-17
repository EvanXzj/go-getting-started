package trans

import "math"

var Pi float64

func init() {
	// 变量除了可以在全局中初始化，也可以在init函数中初始化。这是一个非常特殊的函数，它不能被人调用
	// 有且只有一个init函数，优先级比mian函数高，单线程中自动执行
	Pi = 4 * math.Atan(1) // init() function computes Pi
}
