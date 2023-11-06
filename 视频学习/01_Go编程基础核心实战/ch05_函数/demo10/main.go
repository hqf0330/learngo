package main

import "fmt"

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(t int) {
		t = t + 5
	}(r)
	return 0
}

func f4() (r int) {
	defer func(t *int) {
		*t = *t + 5
	}(&r)
	return 0
}

func main() {
	fmt.Println(f4())
}
