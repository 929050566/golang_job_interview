package main

import "fmt"

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type People1 struct{}

func (p *People1) ShowA() {
	fmt.Println("showA1")
	p.ShowB()
}
func (p *People1) ShowB() {
	fmt.Println("showB1")
}

type Teacher struct {
	People
	People1
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
	t.People1.ShowA()
}

func main() {
	t := Teacher{}
	t.ShowB()
}

// go中没有继承的概念 但是可以通过匿名组合来实现类似继承的效果
// 在这个例子中，Teacher 匿名组合了 People，所以 Teacher 就拥有了 People 的所有方法。
// 当调用 t.ShowA() 时，由于 Teacher 没有实现 ShowA 方法，所以会去调用 People 的 ShowA 再就到了People的shouB方法。
