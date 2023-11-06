package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	str := fmt.Sprintf("Name=%v, Age=%v", s.Name, s.Age)
	return str
}

func main() {
	stu := Student{
		Name: "lili",
		Age:  20,
	}
	fmt.Println(stu)
}
