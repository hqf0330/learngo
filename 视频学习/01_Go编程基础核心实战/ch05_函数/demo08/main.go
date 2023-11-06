package main

import "fmt"

func main() {
	add(30, 60)
}

func add(num1 int, num2 int) int {
	// 遇到defer关键字，不会立即执行defer的语句，
	//而是将defer语句压入到一个defer专用栈中
	// 然后继续执行下面的语句
	defer fmt.Println("num1=", num1)
	defer fmt.Println("num2=", num2)
	num1 += 90
	num2 += 50
	// 在函数执行完毕以后，从defer栈中取出语句开始执行
	sum := num1 + num2
	fmt.Println("sum=", sum)
	return sum
}
