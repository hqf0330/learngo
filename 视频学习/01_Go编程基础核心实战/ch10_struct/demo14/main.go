package main

import "fmt"

type C interface {
	c()
}

type B interface {
	b()
}

type A interface {
	B
	C
	a()
}

type Stu struct {
}

func (s Stu) b() {
	panic("b")
}

func (s Stu) c() {
	panic("c")
}

func (s Stu) a() {
	fmt.Println("a")
}

func main() {
	s := Stu{}
	var a A = s
	a.a()
	a.b()
	a.c()

}
