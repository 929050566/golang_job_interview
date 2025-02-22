// 找问题
package main

import (
	"fmt"
	"time"
)

func main() {
	abc := make(chan int, 1000)
	for i := 0; i < 10; i++ {
		abc <- i
	}
	close(abc)
	go func() {
		for a := range abc {
			fmt.Println("a: ", a)
		}
	}()
	fmt.Println("close")
	time.Sleep(time.Second * 100)
}

// 没问题
