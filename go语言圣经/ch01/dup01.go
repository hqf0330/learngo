// dup2 打印输入中多次出现的行的个数和文本、
// 它从 stdin 或指定的文件列表读取

package main

import (
	"bufio"
	"fmt"
	"os"
)

func findRepeatLine() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func main() {
	findRepeatLine()
}
