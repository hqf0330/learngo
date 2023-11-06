package main

import (
	"bytes"
	"fmt"
	"strings"
)

func basename1(s string) string {
	// 将最后一个'/'和之前的部分全部都舍弃
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

// 往10进制非负整数的字符串中插入逗号
func comma(s string) string {
	n := len(s)
	if n < 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func intToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		_, err := fmt.Fprintf(&buf, "%d", v)
		if err != nil {
			return ""
		}
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(intToString([]int{1, 2, 3}))
}
