package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Printf("%v \n", intChan)

	intChan <- 10
	num := 20
	intChan <- num
	// 不能存放大于容量的数据
	fmt.Printf("管道的实际长度：%v, 管道的容量：%v \n", len(intChan), cap(intChan))

	num1 := <-intChan
	fmt.Println(num1)

	num2 := <-intChan
	fmt.Println(num2)

	num3 := <-intChan
	fmt.Println(num3)
}
