package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var lock sync.RWMutex

func read() {
	defer wg.Done()
	lock.RLock()
	fmt.Println("开始读取数据")
	time.Sleep(time.Second)
	fmt.Println("读数据成功")
	lock.RUnlock()
}

func write() {
	defer wg.Done()
	lock.Lock()
	fmt.Println("开始写数据")
	time.Sleep(time.Second * 5)
	fmt.Println("数据修改成功")
	lock.Unlock()
}

func main() {
	wg.Add(6)

	for i := 0; i < 5; i++ {
		go read()
	}

	go write()

	// go write()
	wg.Wait()
}
