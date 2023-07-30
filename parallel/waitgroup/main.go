package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

type Person struct {
	Name string
	Age  int
}

var p Person

var urls = []string{
	"https://www.wzfry.com/",
	"https://www.google.com/",
	"https://www.baidu.com/",
}

func main() {
	fmt.Printf("%v\n", wg)
	fmt.Printf("%v\n", p)
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			r, err := http.Get(url)
			if err == nil {
				println(r.Status)
			}
		}(url)
	}
	wg.Wait()
}
