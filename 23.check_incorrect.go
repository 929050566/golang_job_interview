package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

// defer 在定义的时候会计算好参数的值，但是并不会立刻执行，而是在函数 return 的时候执行。所以这里的输出是：
// 10 1 2 3
// 20 0 2 2
// 2 0 2 2
// 1 1 3 4
