package main

import (
	"fmt"
)

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

// *Student 实现了show方法 所以这里不会报错
func live() People {
	var stu *Student
	return stu
}

func main() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}

// 在函数 live() 中：
// ```go
// func live() People {
//     var stu *Student  // stu 为 nil，类型为 *Student
//     return stu        // 返回值转换为 People 接口时，内部保存类型信息 *Student 和值 nil
// }
// ```
// - 虽然 stu 的值为 nil，但接口值包含类型信息，因此接口本身并不为 nil。
// - 在 main() 中，判断 if live() == nil 会返回 false，于是打印 "BBBBBBB"。
