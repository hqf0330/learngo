package main

import "fmt"

type AInterface interface {
	a()
}

type BInterface interface {
	b()
}

type Stu struct {
}

func (s Stu) a() {
	fmt.Println("使用AInterface")
}

func (s Stu) b() {
	fmt.Println("使用BInterface")
}

func main() {
	var s Stu
	var a AInterface = s
	var b BInterface = s
	a.a()
	b.b()
	fmt.Println("----------------------")
	s.a()
	s.b()
}
