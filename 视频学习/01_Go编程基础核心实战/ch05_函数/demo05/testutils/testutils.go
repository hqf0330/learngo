package testutils

import "fmt"

var Age int
var Sex int
var Name string

func init() {
	fmt.Println("testutils中的init函数被执行了")
	Age = 19
	Sex = 1
	Name = "小黄"
}
