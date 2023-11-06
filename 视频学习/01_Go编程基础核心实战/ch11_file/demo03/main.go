package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("c:/tmp/test.txt")

	// 使用defer来释放资源
	defer file.Close()

	if err != nil {
		fmt.Println("文件打开失败，", err)
	}

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { // 表示已经读到文件结尾
			break
		}
		fmt.Println(str)
	}
	fmt.Println("文件读取完毕")

}
