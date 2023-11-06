package main

import "fmt"

type SayHello interface {
	sayHello()
}

type Chinese struct {
}

func (c Chinese) sayHello() {
	fmt.Println("你好")
}

func (c Chinese) dance() {
	fmt.Println("扭秧歌")
}

type American struct {
}

func (a American) sayHello() {
	fmt.Println("hello")
}
func greet(s SayHello) {
	s.sayHello()
	if ch, ok := s.(Chinese); ok {
		ch.dance()
	}
}

// 自定义数据类型实现接口
type integer int

func (i integer) sayHello() {
	fmt.Println("say hi + ", i)
}

func main() {
	var i integer = 10
	var s SayHello = i
	// i.sayHello()
	s.sayHello()

	c := Chinese{}
	greet(c)

	a := American{}
	greet(a)

	//var s SayHello = c
	//s.sayHello()
}
