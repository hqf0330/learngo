// 4.4 编写一个rota函数，一次遍历就可完成元素旋转

package main

import "fmt"

func rotate(arr []int, k int) {
	n := len(arr)
	k = k % n

	if k == 0 {
		return
	}

	reverseRotate(arr[:n-k])
	reverseRotate(arr[n-k:])
	reverseRotate(arr)
}

func reverseRotate(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	arr := [7]int{1, 2, 3, 4, 5, 6, 7}
	rotate(arr[:], 3)
	fmt.Println(arr)
}
