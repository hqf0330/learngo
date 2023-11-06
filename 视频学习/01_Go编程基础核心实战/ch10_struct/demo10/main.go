package main

import "fmt"

type Animal struct {
	Age    int
	Weight float32
}

func (an *Animal) Shout() {
	fmt.Println("我可以大声喊叫")
}

func (an *Animal) ShowInfo() {
	fmt.Printf("动物的年龄是：%v, 动物的体重是：%v \n", an.Age, an.Weight)
}

type Cat struct {
	Animal
	Age int
}

func (c *Cat) ShowInfo() {
	fmt.Printf("小猫的年龄是：%v, 动物的体重是：%v \n", c.Age, c.Weight)
}

func (c *Cat) scratch() {
	fmt.Println("我是小猫，可以挠人")
}

func main() {
	//cat := &Cat{}
	//cat.Age = 3
	//cat.Animal.Weight = 10.6
	//cat.Shout()
	//cat.Animal.ShowInfo()
	//cat.scratch()

	cat := &Cat{}
	cat.Weight = 10.6
	cat.Age = 10
	cat.ShowInfo()
	cat.Animal.ShowInfo()

}
