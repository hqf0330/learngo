package main

import (
	"fmt"
	"os"
	"strings"
)

func test01() {
	s, sep := "", " "
	for _, arg := range os.Args {
		s += arg + sep
	}
	fmt.Println(s)
}

func test02() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func main() {
	fmt.Println(strings.Join(os.Args, " "))
	test01()
	test02()
}
