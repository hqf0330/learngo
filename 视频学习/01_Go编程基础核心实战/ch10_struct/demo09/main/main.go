package main

import (
	"../model"
	"fmt"
)

func main() {
	p := model.NewPerson("小黄")
	fmt.Println(*p)

	p.SetAge(18)
	fmt.Println(*p)

	fmt.Println(p.GetAge())
}
