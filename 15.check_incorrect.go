// 找问题
package main

type Student struct {
	name string
}

func main() {
	m := map[string]*Student{"people": {"zhoujielun"}}
	m["people"].name = "wuyanzu"
}

// map的value取出来的都是副本。 不能直接修改副本的属性。
// 但是可以修改副本的属性，然后再赋值给原来的map。
// 但是这里的map的value是指针，所以可以直接修改属性。
