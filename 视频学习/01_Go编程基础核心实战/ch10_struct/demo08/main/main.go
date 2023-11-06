package main

import (
	"../model"
	"fmt"
)

func main() {
	var s = model.NewStudent("娜娜", 16)
	fmt.Println(*s)
}
