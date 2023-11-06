package main

import "fmt"

func main() {
	var scores [5]int
	scores[0] = 95
	scores[1] = 91
	scores[2] = 39
	scores[3] = 60
	scores[4] = 21

	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}

	avg := sum / len(scores)
	fmt.Printf("成绩总和：%v，平均成绩为：%v", sum, avg)
}
