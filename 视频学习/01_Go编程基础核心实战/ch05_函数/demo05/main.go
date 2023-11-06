package main

import (
	utils "./testutils"
	"fmt"
)

var num int = test()

func test() int {
	fmt.Println("main.go文件的test函数被执行了")
	return 10
}

func init() {
	fmt.Println("main.go文件的init函数被执行了")
}

func main() {
	fmt.Println("main.go文件的main函数被执行了")
	fmt.Println("Age = ", utils.Age)
	fmt.Println("Sex = ", utils.Sex)
	fmt.Println("Name = ", utils.Name)
}
