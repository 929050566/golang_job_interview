// ** 机器人坐标问题
// 有一个机器人，给一串指令，L左转 R右转，F前进一步，B后退一步，问最后机器人的坐标，最开始，机器人位于 0 0，方向为正Y。 可以输入重复指令n ： 比如 R2(LF) 这个等于指令 RLFLF。 问最后机器人的坐标是多少？

package main

import "fmt"

type People struct {
	Name string
}

func (p *People) String() string {
	return fmt.Sprintf("print: %v", p)
}

func main() {
	p := &People{}
	p.String()
}

// 代码中的 String 方法是一个无限递归调用，因为在 String
//  方法中调用了 fmt.Sprintf 方法，而 fmt.Sprintf
//  方法会调用 p 的 String 方法，所以会一直递归下去。
// 这个问题可以通过将 fmt.Sprintf 方法中的 %v 替换为 %s 来解决。
