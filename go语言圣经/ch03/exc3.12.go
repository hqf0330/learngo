// 判读两个字符串是否同文异构
package main

import (
	"fmt"
	"sort"
)

func areAnagrams(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	str1Slice := []rune(s1)
	str2Slice := []rune(s2)
	sort.Slice(str1Slice, func(i, j int) bool {
		return str1Slice[i] < str1Slice[j]
	})

	sort.Slice(str2Slice, func(i, j int) bool {
		return str2Slice[i] < str2Slice[j]
	})

	sortedS1 := string(str1Slice)
	sortedS2 := string(str2Slice)

	return sortedS1 == sortedS2

}

func main() {
	fmt.Println(areAnagrams("1230", "123"))
}
