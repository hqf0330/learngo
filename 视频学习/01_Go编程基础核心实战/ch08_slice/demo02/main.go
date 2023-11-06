package main

import "fmt"

func main() {
	slice := make([]int, 4, 20)
	fmt.Println(slice)
	fmt.Println("slice's len", len(slice))
	fmt.Println("slice's cap", cap(slice))

	slice2 := []int{1, 4, 7}
	fmt.Println(slice2)
	fmt.Println("slice2's len", len(slice2))
	fmt.Println("slice2's cap", cap(slice2))
}
