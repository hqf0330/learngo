package main

import "fmt"

func newTest01() {
	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)
}

var global *int

func newTest02() {
	var x int
	x = 1
	global = &x
}

func main() {
	newTest01()
	newTest02()

}
