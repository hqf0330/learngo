package main

import "fmt"

type Teacher struct {
	Name   string
	Age    int
	School string
}

func main() {

	t := new(Teacher)
	(*t).Name = "MSB"
	(*t).Age = 45
	t.School = "TingHua"
	fmt.Println(*t)

	//t := Teacher{
	//	Name:   "zss",
	//	Age:    31,
	//	School: "黑龙江大学",
	//}
	//fmt.Println(t)
}
