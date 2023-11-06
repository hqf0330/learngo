package main

import "fmt"

type Teacher struct {
	Name   string
	Age    int
	School string
}

func main() {
	var t1 Teacher
	// 未赋值的时候是存在初始值的，按照字段定义的类型来初始化
	fmt.Println(t1)
	t1.Name = "MSB"
	t1.Age = 44
	t1.School = "QingHua"
	fmt.Println(t1)
	fmt.Println(t1.Age + 10)
}
