// ** 机器人坐标问题
// 有一个机器人，给一串指令，L左转 R右转，F前进一步，B后退一步，问最后机器人的坐标，最开始，机器人位于 0 0，方向为正Y。 可以输入重复指令n ： 比如 R2(LF) 这个等于指令 RLFLF。 问最后机器人的坐标是多少？

package main

import (
	"unicode"
)

func main() {
	s := "R2(LF)"
	x, y := robotLocation(s)
	println(x, y)
}

func robotLocation(s string) (int, int) {
	x, y := 0, 0
	repeatCount := 0
	lastLeftParenthesis := -1
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'L':
			x = x - 1
		case 'R':
			x = x + 1
		case 'F':
			y = y + 1
		case 'B':
			y = y - 1
		default:
			if unicode.IsDigit(rune(s[i])) {
				repeatCount = int(s[i]-'0') - 1
			}
			if s[i] == '(' {
				lastLeftParenthesis = i
			}
			// 遇到右括号，回退到左括号位置 回退次数由repeatCount决定
			if s[i] == ')' && repeatCount > 0 {
				i = lastLeftParenthesis
				repeatCount = repeatCount - 1
			}
		}
	}
	return int(x), int(y)
}
