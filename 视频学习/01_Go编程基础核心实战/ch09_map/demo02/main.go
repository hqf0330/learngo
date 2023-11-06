package main

import "fmt"

func main() {
	a := make(map[string]map[int]string)

	a["班级1"] = make(map[int]string, 3)
	a["班级1"][200101] = "露露"
	a["班级1"][200102] = "丽丽"
	a["班级1"][200103] = "菲菲"

	a["班级2"] = make(map[int]string, 3)
	a["班级2"][200201] = "lulu"
	a["班级2"][200202] = "lili"
	a["班级2"][200203] = "feifei"

	for k1, v1 := range a {
		fmt.Println(k1)
		for k2, v2 := range v1 {
			fmt.Printf("学号为：%v, 姓名为：%v \t", k2, v2)
		}
		fmt.Println()
	}

	fmt.Println(a)
	fmt.Println(a["班级1"])
	fmt.Println(a["班级2"])
}
