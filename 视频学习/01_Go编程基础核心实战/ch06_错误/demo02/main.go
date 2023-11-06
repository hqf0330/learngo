package main

import (
	"errors"
	"fmt"
)

func test() (err error) {

	num1 := 10
	num2 := 0
	if num2 == 0 {
		return errors.New("除数不为0")
	} else {
		result := num1 / num2
		fmt.Println(result)
		return nil
	}

}

func main() {
	err := test()
	if err != nil {
		fmt.Println("自定义错误", err)
		panic(err)
	}
	fmt.Println("除法正常使用！")
}
