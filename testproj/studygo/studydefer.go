package studygo

import "fmt"

// 1. defer是栈方式调用，后写先调用
// 2. 参数如果是函数，则会先计算
// 3. defer的函数参数，会固定下来

// 结果
// 1 2 3
// 0 2 2
// 0 2 2
// 1 3 4

func DeferCalc() {
	a := 1
	b := 2
	defer calc(a, calc(a, b))
	a = 0
	defer calc(a, calc(a, b))
}

func calc(x, y int) int {
	fmt.Println(x, y, x+y)
	return x + y
}
