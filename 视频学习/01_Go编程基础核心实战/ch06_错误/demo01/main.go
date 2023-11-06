package main

import "fmt"

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("捕捉到错误")
			fmt.Println("err: ", err)
		}
	}()

	defer fmt.Println("准备执行defer")

	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println(result)
}

func main() {
	test()
	fmt.Println("除法正常使用！")
}
