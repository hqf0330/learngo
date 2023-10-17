package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	symbol := [...]string{USD: "$", EUR: "yuan", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB])
	r := [...]int{99: -1}
	fmt.Println(r)
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c)
}
