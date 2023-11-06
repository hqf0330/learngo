package main

import "fmt"

func main() {
	arr := [6]int{1, 2, 3, 4, 5, 6}
	slice := arr[1:4]
	//fmt.Println(slice[0])
	//fmt.Println(slice[1])
	//fmt.Println(slice[2])
	slice2 := slice[1:2]
	fmt.Println(slice2)
	slice2[0] = 66
	fmt.Println("--------------")
	fmt.Println(arr)
	fmt.Println(slice)
	fmt.Println(slice2)
}
