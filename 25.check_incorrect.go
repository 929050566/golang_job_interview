package main

import "sync"

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

// 在并发编程中，即使是读操作也需要加锁，因为在 Go 中，map 不是线程安全的。
// 如果一个 goroutine 正在写 map，而另一个 goroutine 正在读 map，就会导致数据竞争
// 进而引发运行时错误（panic: concurrent map read and map write）。

// 为了确保线程安全，所有对 map 的访问（包括读和写）都需要加锁。下面是修改后的代码：
