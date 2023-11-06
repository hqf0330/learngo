package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func badTest() {
	s, sep := "", " "
	for _, arg := range os.Args[1:] {
		s += arg + sep
	}
	fmt.Println(s)
}

func goodTest() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {

	startTime := time.Now()
	badTest()
	endTime := time.Now()
	time01 := endTime.Sub(startTime) / time.Microsecond
	fmt.Printf("第一个循环的执行时间：%v\n", time01)

	startTime = time.Now()
	goodTest()
	endTime = time.Now()
	time02 := endTime.Sub(startTime) / time.Microsecond
	fmt.Printf("第二个循环的执行时间：%v\n", time02)

}
