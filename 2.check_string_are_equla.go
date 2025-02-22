// ** 判断字符串中字符是否全都不同 即无重复字符串 不使用额外空间
package main

func main() {
	s := "abcdefghijklnmopqrstuvwxyza"
	println(isUniqueString(s))
	println(isUniqueString2(s))
}

func isUniqueString(s string) bool {
	// 利用map
	m := make(map[byte]struct{})
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			return false
		}
		m[s[i]] = struct{}{}
	}
	return true
	// 利用count
	// 利用index
}

func isUniqueString2(s string) bool {
	var make1, make2, make3, make4 uint64
	for _, c := range s {
		c -= 'A'
		if c < 64 {
			if make1&(1<<c) != 0 {
				return false
			}
			make1 |= 1 << c
		} else if c < 128 {
			c -= 64
			if make2&(1<<c) != 0 {
				return false
			}
			make2 |= 1 << c
		} else if c < 192 {
			c -= 128
			if make3&(1<<c) != 0 {
				return false
			}
			make3 |= 1 << c
		} else {
			c -= 192
			if make4&(1<<c) != 0 {
				return false
			}
			make4 |= 1 << c
		}
	}
	return true
}
