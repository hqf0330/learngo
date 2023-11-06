package main

import (
	"fmt"
)

func modifySlice(slice []int) {
	slice = append(slice, 4) // 在切片的副本上添加元素4
	slice[0] = 99            // 在切片的副本上修改第一个元素
}

func change(tmp *[]int) {
	*tmp = append(*tmp, 4)
	fmt.Printf("%p\n", tmp)
	fmt.Printf("%v %d %d %p\n", *tmp, len(*tmp), cap(*tmp), *tmp)
}

func main() {
	data := []int{1, 2, 3}
	fmt.Println(data)
	modifySlice(data) // 调用函数修改切片
	fmt.Println(data)

	tmp := make([]int, 0, 5)
	tmp = append(tmp, 1, 2, 3)
	fmt.Printf("%p\n", &tmp)
	fmt.Printf("%v %d %d %p\n", tmp, len(tmp), cap(tmp), tmp)
	change(&tmp)
	fmt.Printf("%v %d %d %p\n", tmp, len(tmp), cap(tmp), tmp)
}
