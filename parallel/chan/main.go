package main

import "runtime"

func main() {
	c := make(chan struct{})
	go func() {
		sum := 0
		for i := 0; i < 10000; i++ {
			sum += i
		}
		println(sum)
		c <- struct{}{}
	}()
	println("NumberGoroutine=", runtime.NumGoroutine())

	// 读通道c进行等待,忽略结果
	<-c
}
