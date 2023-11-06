package main

import "fmt"

var Func01 = func(num1 int, num2 int) int {
	return num1 * num2
}

func main() {
	// 定义并且使用一个匿名函数：
	result := func(num1 int, num2 int) int {
		return num1 + num2
	}(10, 20)

	fmt.Println(result)

	// 将匿名函数赋值给一个变量，这个变量其实就是函数类型的变量
	// sub等价匿名函数
	sub := func(num1 int, num2 int) int {
		return num1 - num2
	}

	result2 := sub(30, 70)
	fmt.Println(result2)

	result3 := sub(70, 40)
	fmt.Println(result3)

	result4 := Func01(3, 4)
	fmt.Println(result4)
}
