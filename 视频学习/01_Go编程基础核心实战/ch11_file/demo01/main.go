package main

import (
	"fmt"
	"os"
)

func main() {
	// 打开文件
	file, err := os.Open("c:/tmp/test.txt")

	// 关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("文件关闭出错，错误为：", err)
		}
	}(file)

	if err != nil {
		fmt.Println("文件打开出错，对应错误为：", err)
	}

	fmt.Println(file)
}
