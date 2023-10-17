// 4.5 编写一个就地处理函数，用于去除[]string slice中相邻的重复字符串元素
package main

import "fmt"

func removeAdjacentDup(slice []string) []string {
	if len(slice) == 0 {
		return slice
	}

	result := []string{slice[0]}

	for i := 1; i < len(slice); i++ {
		if slice[i] != slice[i-1] {
			result = append(result, slice[i])
		}
	}

	return result
}

func main() {
	inputSlice := []string{"apple", "banana", "banana", "cherry", "cherry", "cherry", "date"}
	resultSlice := removeAdjacentDup(inputSlice)
	fmt.Println(resultSlice)
}
