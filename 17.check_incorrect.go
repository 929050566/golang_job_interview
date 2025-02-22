package main

import (
	"fmt"
	"runtime"
)

func main() {
	var i byte
	go func() {
		for i = 0; i <= 255; i++ {
		}
	}()
	fmt.Println("Dropping mic")
	// Yield execution to force executing other goroutines
	runtime.Gosched()
	// time.Sleep(time.Second)
	runtime.GC()
	fmt.Println("Done")
}

// 会一直执行，因为i是byte类型，最大值是255，所以i永远不会大于255，并且抢占了CPU，所以会一直执行
