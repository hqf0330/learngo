package main

import "fmt"

// 函数功能：求和
// 函数名字：getSum，参数列表为空
// 函数返回值：一个参数为一个int，返回值也是一个int的函数
func getSum() func(int) int {
	var sum int = 0
	return func(num int) int {
		sum += num
		return sum
	}
}

// 闭包：返回的匿名函数+匿名函数以外的变量（逃逸现象）

func main() {
	f := getSum()
	fmt.Println(f(1)) // 1
	fmt.Println(f(2)) // 3
	fmt.Println(f(3)) // 6
}
