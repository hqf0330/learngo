package main

import "fmt"

// 多继承

type A struct {
	a int
	b string
}

type B struct {
	c int
	d string
	a int
}

type C struct {
	A
	B
	int
}

type D struct {
	*A
	*B
	int
}

type E struct {
	a int
	b int
	c B
}

func main() {

	e := E{
		a: 1,
		b: 2,
		c: B{1, "hh", 10},
	}

	fmt.Println(e)
	fmt.Println(e.c.a)

	c := C{
		A:   A{10, "a"},
		B:   B{20, "b", 50},
		int: 888,
	}
	fmt.Println(c.b)
	fmt.Println(c.d)
	fmt.Println(c.A.a)
	fmt.Println(c.B.a)
	fmt.Println(c.int)

	d := D{
		A: &A{
			a: 10,
			b: "a",
		},
		B: &B{
			c: 20,
			d: "b",
			a: 50,
		},
		int: 888,
	}
	fmt.Println(d)
	fmt.Println(*d.A)
	fmt.Println(*d.B)
}
