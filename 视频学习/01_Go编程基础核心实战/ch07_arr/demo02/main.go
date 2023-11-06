package main

import "fmt"

func main() {
	var arr [3]int16
	fmt.Println(len(arr))
	fmt.Println(arr)

	fmt.Printf("arr的地址：%p \n", &arr)
	fmt.Printf("arr第一个元素的地址：%p \n", &arr[0])
	fmt.Printf("arr第二个元素的地址：%p \n", &arr[1])
	fmt.Printf("arr第三个元素的地址：%p \n", &arr[2])

}
