package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "golang你好"
	fmt.Println(len(str))

	// 方式一
	for i, value := range str {
		fmt.Printf("index: %d, value: %c \n", i, value)
	}

	// 方式二
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c \n", r[i])
	}

	// 字符串转数字
	num1, err := strconv.Atoi("666")
	if err != nil {
		return
	}
	fmt.Println(num1)

	// 整数转字符串
	str1 := strconv.Itoa(88)
	fmt.Println(str1)

	// 统计字符串有多少个字串
	count := strings.Count("golangjava", "ja")
	fmt.Println(count)

	flag := strings.EqualFold("go", "Go")
	fmt.Println(flag)
	fmt.Println("go" == "Go")

	index := strings.Index("golangjavaja", "ja")
	fmt.Println(index)

}
