package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan <- i
	}

	// 如果在遍历之前没有关闭chan，那么就会出现deadlock的问题
	close(intChan)

	for v := range intChan {
		fmt.Println("value: ", v)
	}
}
