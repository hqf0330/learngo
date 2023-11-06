package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.OpenFile("c:/tmp/test2.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败")
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("你好，我是新内容 \n")
	}
	writer.Flush()

}
