// 找问题
package main

import "sync/atomic"

var value int32

func SetValue(delta int32) {
	for {
		v := value
		if atomic.CompareAndSwapInt32(&value, v, (v + delta)) {
			break
		}
	}
}

// atomic does not need to use a loop to set value
