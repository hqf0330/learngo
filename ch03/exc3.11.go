package main

import (
	"bytes"
	"fmt"
	"strings"
)

func commaProcess(s string) string {
	if len(s) <= 3 {
		return s
	}
	return commaProcess(s[:len(s)-3]) + "," + s[len(s)-3:]
}

func richComma(s string) string {
	var buf bytes.Buffer

	if (s[0] == '-' || s[0] == '+') && len(s[1:]) <= 3 {
		return s
	}

	// 处理负数
	if len(s) > 0 && (s[0] == '-' || s[0] == '+') {
		buf.WriteByte(s[0])
		s = s[1:]
	}

	intPart, fracPart := "", ""
	if dotIndex := strings.Index(s, "."); dotIndex != -1 {
		intPart = s[:dotIndex]
		fracPart = s[dotIndex+1:]
	} else {
		intPart = s
	}

	if fracPart != "" {
		s = commaProcess(intPart) + "." + fracPart
		buf.WriteString(s)
	} else {
		s = commaProcess(intPart)
		buf.WriteString(s)
	}
	return buf.String()
}

func main() {
	fmt.Println(richComma("12345.12"))
	fmt.Println(richComma("-12345.12"))
	fmt.Println(richComma("+12345.12"))
}
