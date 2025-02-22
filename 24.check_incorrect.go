package main

import "fmt"

func main() {
	s := make([]int, 5, 6)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

// append 会改变 slice 的长度，所以这里会输出 [0 0 0 0 0 1 2 3]。
