// ** 机器人坐标问题
// 有一个机器人，给一串指令，L左转 R右转，F前进一步，B后退一步，问最后机器人的坐标，最开始，机器人位于 0 0，方向为正Y。 可以输入重复指令n ： 比如 R2(LF) 这个等于指令 RLFLF。 问最后机器人的坐标是多少？

package main

type Param map[string]interface{}

type Show struct {
	Param
}

func main1() {
	s := new(Show)
	s.Param["RMB"] = 10000
}

// 共发现两个问题：
// main 函数不能加数字。
// new 关键字无法初始化 Show 结构体中的 Param 属性，所以直接对 s.Param 操作会出错。

type student struct {
	Name string
}

func zhoujielun(v interface{}) {
	// switch msg := v.(type) {
	// case *student, student:
	// 	// msg.Name
	// }
}
