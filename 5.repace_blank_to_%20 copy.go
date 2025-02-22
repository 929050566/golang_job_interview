// ** 替换字符串 空额为%20
package main

import (
	"strings"
	"unicode"
)

func main() {
	s := "ab cdefghij klnmopq rstuvwxyza"
	ss, ok := repalceBlanks(s)
	if ok {
		print(ss)
	} else {
		print("string has letter and blank")
	}
}

func repalceBlanks(s string) (string, bool) {
	for _, v := range s {
		if v == ' ' && unicode.IsLetter(v) {
			return s, false
		}
	}
	return strings.Replace(s, " ", "%20", -1), true
}
