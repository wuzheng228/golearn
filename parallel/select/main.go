package main

func main() {
	ch := make(chan int, 1)
	go func(chan int) {
		for {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}(ch)

	for i := 0; i < 10000; i++ {
		println(<-ch)
	}
}
