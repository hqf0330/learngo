package main

import "fmt"

func main() {
	// 定义map
	b := make(map[int]string)
	b[1001] = "张三"
	b[1002] = "李四"
	fmt.Println(b)
	b[1002] = "王五"
	fmt.Println(b)
	delete(b, 1001)
	fmt.Println(b)

	value, flag := b[1001]
	fmt.Println(value)
	fmt.Println(flag)
}
