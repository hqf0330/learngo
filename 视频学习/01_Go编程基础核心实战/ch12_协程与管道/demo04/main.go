package main

import (
	"fmt"
	"sync"
)

var num int
var wg sync.WaitGroup

func add() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		num = num + 1
	}
}

func sub() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		num = num - 1
	}
}

func main() {
	// 多协程操作同一个数据
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(num)
}
