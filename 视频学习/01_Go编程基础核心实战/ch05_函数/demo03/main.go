package main

import "fmt"

func test(num int) {
	fmt.Println(num)
}

func test2(num1 int, num2 float32, f func(int)) {
	fmt.Println("我在test2调用test函数")
	f(num1)
}

type myFunc func(int)

func test3(num1 int, num2 float32, testFunc myFunc) {
	testFunc(num1)
}

func test4(num1 int, num2 int) (sub int, sum int) {
	sub = num1 - num2
	sum = num1 + num2
	return
}

func main() {
	a := test
	fmt.Printf("a的类型是：%T， test函数的类型是：%T \n", a, test)
	a(10)
	test2(10, 1.2, test)
	test2(111, 1.2, a)
}
