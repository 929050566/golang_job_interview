package main

import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Student struct{}

// func (stu *Student) Speak(think string) (talk string) {
func (stu Student) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return ""
}

func main() {
	var peo People = Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}

// *Student 改为 Student就行，因为压根不是两种类型
