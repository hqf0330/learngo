package main

import "fmt"

func main() {
	f := 3.141
	i := int(f)
	fmt.Println(f, i)
	f = 1.99
	fmt.Println(int(f))
	c := 1e100
	i = int(c)
	fmt.Println(i)
}
