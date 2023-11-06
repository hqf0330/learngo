package main

import "fmt"

func main() {

	// 声明只写管道
	var intChan1 chan<- int
	intChan1 = make(chan int, 3)
	intChan1 <- 1
	// fmt.Println(<-intChan1)
	// 声明只读管道
	var intChan2 <-chan int
	if intChan2 != nil {
		num := <-intChan2
		fmt.Println("num1 ", num)
	}

}
