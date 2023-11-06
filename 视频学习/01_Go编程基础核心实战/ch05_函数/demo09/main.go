package main

import "fmt"

func f1() {
	defer fmt.Println("hello world")
	fmt.Println("hello defer!")
}

func f2() {
	defer fmt.Println("hello 1!")
	defer fmt.Println("hello 2!")
	defer fmt.Println("hello 3!")

	fmt.Println("hello defer!")
}

func main() {
	f2()
}
