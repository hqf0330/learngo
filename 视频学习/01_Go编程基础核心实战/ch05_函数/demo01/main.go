package main

import "fmt"

func exchangeNumValue(num1 int, num2 int) {
	var t = num1
	num1 = num2
	num2 = t
}

func exchangeNumPoint(num1 *int, num2 *int) {
	var t = *num1
	*num1 = *num2
	*num2 = t
}

func main() {
	num1 := 10
	num2 := 20
	fmt.Printf("num1: %v, num2: %v", num1, num2)
	exchangeNumPoint(&num1, &num2)
	fmt.Println()
	fmt.Printf("num1: %v, num2: %v", num1, num2)
}
