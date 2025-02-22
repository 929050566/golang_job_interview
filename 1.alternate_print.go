// ** 交替打印数字和字母
package main

import "sync"

func main() {
	letter, number := make(chan bool, 1), make(chan bool, 1)
	var group = sync.WaitGroup{}

	group.Add(2)
	go func() {
		defer func() {
			group.Done()
			print("number print is done")
		}()
		for i := 1; i <= 26; i++ {
			<-number
			print(i)
			letter <- true
		}
	}()
	go func() {
		defer func() {
			group.Done()
			print("letter print is done")
		}()
		for i := 'A'; i <= 'Z'; i++ {
			<-letter
			print(string(i))
			if i == 'Z' {
				print("letter print is here1")
			}
			number <- true
			if i == 'Z' {
				print("letter print is here2")
			}
		}
	}()
	number <- true
	group.Wait()
	<-number
}
