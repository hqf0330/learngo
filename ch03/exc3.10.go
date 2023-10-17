package main

import (
	"bytes"
	"fmt"
)

func commaNoRecursion(s string) string {

	if len(s) <= 3 {
		return s
	}

	var buf bytes.Buffer
	for i := len(s) - 1; i >= 0; i-- {
		buf.WriteByte(s[i])
		if (len(s)-i)%3 == 0 && i > 0 {
			buf.WriteByte(',')
		}
	}

	result := buf.String()
	reversed := []rune(result)
	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}

	return string(reversed)

}

func main() {
	fmt.Println(commaNoRecursion("12345"))
}
