package main

import "fmt"

func main() {
	s := make([]int, 4, 20)
	s[0] = 1
	s[1] = 2
	s[2] = 3
	s[3] = 4

	for i := 0; i < len(s); i++ {
		fmt.Printf("slice[%v] = %v \t", i, s[i])
	}

	fmt.Println("------------------")

	for index, value := range s {
		fmt.Printf("index: %v, value: %v \n", index, value)
	}
}
