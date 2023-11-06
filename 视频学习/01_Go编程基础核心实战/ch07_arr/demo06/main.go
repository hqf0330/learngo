package main

import "fmt"

func test(arr [3]int) {
	arr[0] = 7
}

func test1(arr *[3]int) {
	(*arr)[0] = 7
}

func main() {
	var arr3 = [3]int{3, 6, 9}
	test1(&arr3)
	fmt.Println(arr3)
}
