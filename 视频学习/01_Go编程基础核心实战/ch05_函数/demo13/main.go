package main

import (
	"fmt"
	"time"
)

func main() {
	// 这个Now函数返回的是一个结构体
	now := time.Now()
	fmt.Printf("%v, 对于的类型为：%T \n", now, now)
	fmt.Printf("年：%v \n", now.Year())
	fmt.Printf("月：%v \n", now.Month())
	fmt.Printf("日：%v \n", now.Day())
	fmt.Printf("时：%v \n", now.Hour())
	fmt.Printf("分：%v \n", now.Minute())
	fmt.Printf("秒：%v \n", now.Second())

	fmt.Printf("%d-%d-%d %d:%d:%d \n", now.Year(),
		now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	dateStr := fmt.Sprintf("%d-%d-%d %d:%d:%d \n", now.Year(),
		now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(dateStr)

	dateStr2 := now.Format("2006/01/02 14:02:03")
	fmt.Println(dateStr2)
}
