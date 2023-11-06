package main

import (
	"fmt"
	"sync"
)

var num int
var wg sync.WaitGroup

// 加入互斥锁
var lock sync.Mutex

func add() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		lock.Lock()
		num = num + 1
		lock.Unlock()
	}
}

func sub() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		lock.Lock()
		num = num - 1
		lock.Unlock()
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
