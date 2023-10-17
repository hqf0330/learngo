package main

import (
	"./tempconv" // 导入当前目录中我们魔改过的 tempconv 包
	"fmt"
)

func main() {
	fmt.Println("摄氏度温度中：")
	fmt.Println("绝对零度：", tempconv.AbsoluteZeroC)
	fmt.Println("冰点：", tempconv.Freezing)
	fmt.Println("沸点：", tempconv.BoilingC)
}
