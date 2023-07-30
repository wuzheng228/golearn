package main

import "fmt"

func main() {
	Case1()
	Case2()
	Case3()
}

func Case1() {
	var s []int
	for i := 0; i < 3; i++ {
		s = append(s, i)
	}
	modifySlice1(s)
	fmt.Printf("case1: %v\n", s)
}

func Case2() {
	var s []int
	for i := 0; i < 3; i++ {
		s = append(s, i)
	}
	modifySlice2(s)
	fmt.Printf("case2: %v\n", s)
}

func Case3() {
	var s []int
	for i := 0; i < 3; i++ {
		s = append(s, i)
	}
	modifySlice3(s)
	fmt.Printf("case3: %v\n", s)
}

func modifySlice3(s []int) {
	s = append(s, 2048)
	s = append(s, 4096)
	fmt.Printf("modifySlice3 len: %v, cap: %v\n", len(s), cap(s))
	s[0] = 1024
}

func modifySlice2(s []int) {
	s = append(s, 2048)
	fmt.Printf("modifySlice2 len: %v, cap: %v\n", len(s), cap(s))
	s[0] = 1024
}

func modifySlice1(s []int) {
	s[0] = 1024
}
