package main

import (
	"fmt"
	"os"
)

func main() {
	// 定义源文件
	srcPath := "c:/tmp/test.txt"
	// 定义目标文件
	destPath := "c:/tmp/test3.txt"

	content, err := os.ReadFile(srcPath)
	if err != nil {
		fmt.Println("读取有问题")
		return
	}

	err = os.WriteFile(destPath, content, 0666)
	if err != nil {
		fmt.Println("文件写错错误")
	}
}
