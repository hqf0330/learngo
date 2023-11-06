package model

import "fmt"

type person struct {
	Name string
	age  int
}

func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

// SetAge 定义set方法，可以包外访问
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("你输入的年龄有问题")
	}
}

func (p *person) GetAge() int {
	return p.age
}
