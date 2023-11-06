package main

import "fmt"

func main() {
	var intArr = [6]int{3, 6, 9, 1, 4, 7}
	// var slice = intArr[1:3]
	slice := intArr[1:3]
	fmt.Println(intArr)
	fmt.Println(slice)
	fmt.Println("slice的元素个数：", len(slice))
	fmt.Println("slice的容量：", cap(slice))

	fmt.Println("====================")
	fmt.Printf("数组下标为1位置的地址：%p \n", &intArr[1])
	fmt.Printf("切片下标为0位置的地址：%p \n", &slice[0])

	slice[1] = 16
	fmt.Println("intArr: ", intArr)
	fmt.Println("slice: ", slice)

}
