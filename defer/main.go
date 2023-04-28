package main

import (
	"fmt"
	"os"
)

func main() {
	// defer recoverFromPanic()
	fmt.Println("hellow world")
	readDataFromFile()
}

func readDataFromFile() {

	file, err := os.Open("data.txt")
	if err != nil {
		// panic(err)
		return
	}
	defer fmt.Println("defer while err occured")
	fmt.Println(file.Name())
}

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recover from panic", r)
	}
	main()
}
