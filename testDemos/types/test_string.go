package main

import (
	"fmt"
	"unicode/utf8"
)

func testString() {
	fmt.Println("He said:\"Hello GO!\" ")
	fmt.Println("-------------------")
	fmt.Println(`He said: "Hello GO!"
	Hello World!	
	Hello Gland
	`)
	fmt.Println("-------------------")
	fmt.Println(len("你好"))
	fmt.Println(utf8.RuneCountInString("你好"))
	fmt.Println(utf8.RuneCountInString("你好ab"))
	fmt.Println("Hello, " + "GO!")
}

func main() {
	testString()
}
