// ** 翻转字符串，不使用额外空间
package main

func main() {
	s := "abcdefghijklnmopqrstuvwxyza"
	println(reverStr(s))
}

func reverStr(s string) string {
	str := []rune(s)
	for i := 0; i < len(str)/2; i++ {
		str[i], str[len(str)-1-i] = str[len(str)-1-i], str[i]
	}
	return string(str)
}
