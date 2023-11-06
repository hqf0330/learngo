package main

import (
	"../dbutils"
	"fmt"
)

func main() {
	fmt.Println("执行main方法")
	dbutils.GetConn()
}
