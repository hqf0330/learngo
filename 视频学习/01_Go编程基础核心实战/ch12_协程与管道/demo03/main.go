package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for i := 1; i <= 5; i++ {
		// 线程开启就+1
		wg.Add(1)
		go func(n int) {
			// 线程结束就减1
			defer wg.Done()
			fmt.Println(n)
		}(i)
	}

	// 主线程阻塞，什么时候wg计数器减为0，那么就停止
	wg.Wait()
}
