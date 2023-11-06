package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) test() {
	p.Name = "露露"
	fmt.Println(p.Name)
}

func main() {
	var p Person
	p.Name = "丽丽"
	p.test()
	fmt.Println(p.Name)
}
