package main

import (
	"fmt"
	"math/rand"
)

func GenerateIntA(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Label:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Label
			}
		}
		close(ch)
	}()
	return ch
}

func GenerateIntB(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Label:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Label
			}
		}
		close(ch)
	}()
	return ch
}

func GenerateInt(done chan struct{}) chan int {
	ch := make(chan int)
	send := make(chan struct{})
	go func() {
	Label:
		for {
			select {
			case ch <- <-GenerateIntA(send):
			case ch <- <-GenerateIntB(send):
			case <-done:
				send <- struct{}{}
				send <- struct{}{}
				break Label
			}
		}

		close(ch)
	}()
	return ch
}

func main() {
	// 创建一个接受退出信号的chan
	done := make(chan struct{})
	ch := GenerateInt(done)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	done <- struct{}{}

	fmt.Println("stop generate")
}
