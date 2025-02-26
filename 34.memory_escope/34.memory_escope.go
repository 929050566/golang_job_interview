package main

import "fmt"

// 知道golang的内存逃逸吗？什么情况下会发生内存逃逸？

type A struct {
	s string
}

// 这是上面提到的 "在方法内把局部变量指针返回" 的情况
func foo(s string) *A {
	a := new(A)
	a.s = s
	return a //返回局部变量a,在C语言中妥妥野指针，但在go则ok，但a会逃逸到堆
}

type Student struct {
	Name string
}

func main() {
	a := foo("hello")
	b := a.s + " world"
	c := b + "!"
	fmt.Println(c)

	str1 := []string{"a", "b", "c"}
	str2 := str1[1:]
	str2[1] = "new"
	fmt.Println(str1)
	fmt.Println(str2)
	str2 = append(str2, "z", "x", "y")
	fmt.Println(str1)
	fmt.Println(str2)

	fmt.Println(&Student{Name: "menglu"} == &Student{Name: "menglu"}) // false
	fmt.Println(Student{Name: "menglu"} == Student{Name: "menglu"})   // true

	fmt.Println([...]string{"1"} == [...]string{"1"})
	// fmt.Println([]string{"1"} == []string{"1"})

	kv := map[string]Student1{"menglu": {Age: 21}}
	if name, ok := kv["menglu"]; ok {
		name.Age = 22
	}
	s := []Student1{{Age: 21}}
	s[0].Age = 22
	fmt.Println(kv, s)
}

type Student1 struct {
	Age int
}
