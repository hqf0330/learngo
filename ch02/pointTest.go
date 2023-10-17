package main

import "fmt"

func testPoint01() {
	x := 1
	p := &x
	fmt.Println(x)
	fmt.Println(*p)
	fmt.Println(p)
	fmt.Println(x == *p)
	fmt.Println("通过指针修改数据")
	*p = 2
	fmt.Println(x)
	fmt.Println(p)
}

// 这个变量的地址被返回了，属于逃逸
// 这个时候就会把这个变量分配到堆上，这个函数被释放以后这个地址还能
// 指向正确的地方
func testPoint02() *int {
	v := 1
	return &v
}

func testPoint03(p *int) int {
	*p++
	return *p
}

func main() {
	testPoint01()
	fmt.Println("----------1---------")

	p2 := testPoint02()
	fmt.Println(p2)
	fmt.Println(*p2)
	fmt.Println("----------2---------")

	v := 1
	testPoint03(&v)
	fmt.Println(v)
	fmt.Println("----------3---------")

}

// 写周报
