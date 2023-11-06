package main

import (
	"fmt"
	"time"
)

func main() {
	// 使用匿名函数启动一个协程，直接调用匿名函数
	//go func() {
	//	fmt.Println(1)
	//}()

	for i := 1; i <= 5; i++ {
		go func(n int) {
			fmt.Println(n)
		}(i)
	}

	time.Sleep(time.Second)
}
