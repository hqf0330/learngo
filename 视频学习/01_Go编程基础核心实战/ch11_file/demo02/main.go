package main

import (
	"fmt"
	"os"
)

func main() {
	// 读取文件
	// 返回的是一个byte的切片，内部自动打开然后关闭文件
	content, err := os.ReadFile("c:/tmp/test.txt")

	if err != nil {
		fmt.Println("读取出错，错误为：", err)
	}
	fmt.Printf("%v", string(content))

}
