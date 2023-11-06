package main

import "fmt"

func main() {
	var arr [2][3]int16
	fmt.Println(arr)
	fmt.Printf("arr的地址是：%p \n", &arr)
	fmt.Printf("arr[0]的地址是：%p \n", &arr[0])
	fmt.Printf("arr[0][0]的地址是：%p \n", &arr[0][0])

	var arr1 [2][3]int = [2][3]int{{1, 4, 7}, {2, 5, 8}}
	fmt.Println(arr1)

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[i]); j++ {
			fmt.Println(arr1[i][j])
		}
	}

	fmt.Println("----------------")

	for key, value := range arr1 {
		for k, v := range value {
			fmt.Printf("arr[%v][%v]: %v \t", key, k, v)
		}
		fmt.Println()
	}
}
