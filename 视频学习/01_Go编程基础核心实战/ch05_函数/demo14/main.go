package main

import "fmt"

func main() {
	num := new(int)
	fmt.Printf("type: %T, value: %v, address: %v, pointValue: %v", num, num, &num, *num)
}
