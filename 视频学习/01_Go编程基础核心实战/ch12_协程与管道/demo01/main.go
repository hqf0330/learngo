package main

import (
	"fmt"
	"time"
)

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("线程！！！", i)
		// 阻塞1秒
		time.Sleep(time.Second * 1)
	}
}

func main() {

	// 开启一个协程
	go test()

	// 主线程
	for i := 1; i <= 10; i++ {
		fmt.Println("主线程", i)
		// 阻塞1秒
		time.Sleep(time.Second * 1)
	}
}
