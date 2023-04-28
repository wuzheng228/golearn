package main

import "fmt"

type Inter interface {
	Ping()
	Pang()
}

type St struct{}

func (St) Ping() {
	println("ping")
}

func (*St) Pang() {
	println("pang")
}

func main() {
	var st *St = nil
	var it Inter = st

	fmt.Printf("%p\n", st)
	fmt.Printf("%p\n", it)

	if it != nil {
		it.Pang()

		instance := *it.(*St)

		fmt.Printf("%f", instance)
	}
}
