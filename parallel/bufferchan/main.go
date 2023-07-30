package main

import "runtime"

func main() {
	c := make(chan struct{})
	ci := make(chan int, 100)
	go func(i chan struct{}, j chan int) {
		for i := 0; i < 10; i++ {
			ci <- i
		}
		close(ci)
		// 写通道
		c <- struct{}{}
	}(c, ci)
	// 打印携程的数量
	println("NumGoroutine=", runtime.NumGoroutine())

	// 读通道等待携程处理完成
	<-c
	println("NumGoroutine=", runtime.NumGoroutine())
	// 缓存通道ci虽然关闭了, 但是缓存内的数据还可以继续读取
	for v := range ci {
		println(v)
	}
}
