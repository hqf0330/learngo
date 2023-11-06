package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := strings.Replace("goandjavagogo", "go", "golang", -1)
	fmt.Println(str1)

	arr := strings.Split("go-python-java", "-")
	fmt.Println(arr)

	fmt.Println(strings.TrimSpace("   go and  java  "))
	fmt.Println(strings.Trim("~golang~", "~"))
	fmt.Println(strings.TrimLeft("~golang~", "~"))
	fmt.Println(strings.TrimRight("~golang~", "~"))
	fmt.Println(strings.HasPrefix("https://www.baidu.com", "https"))
	fmt.Println(strings.HasSuffix("https://www.baidu.com", ".com"))
}
