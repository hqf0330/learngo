// 练习4.3：重写reverse函数，使用数组指针作为参数而不是slice

package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseArr(p *[5]int) {
	for i, j := 0, len(*p)-1; i < j; i, j = i+1, j-1 {
		(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
	}
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	reverse(arr[:])
	fmt.Println(arr)
	reverseArr(&arr)
	fmt.Println(arr)
}
