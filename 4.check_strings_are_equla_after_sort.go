// ** check string are equal after sort
// https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/
package main

import "strings"

func main() {
	s := "abcdefghijklnmopqrstuvwxyza"
	println(checkEqual(s, s))
}

func checkEqual(s1 string, s2 string) bool {
	str1 := len([]rune(s1))
	str2 := len([]rune(s2))
	if str1 != str2 || str1 > 5000 || str2 > 5000 {
		return false
	}
	for _, v := range s1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}
	return true
}
