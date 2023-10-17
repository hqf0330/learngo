package main

import "fmt"

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || xv != yv {
			return false
		}
	}
	return true
}

func main() {
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	fmt.Println("--------------------------")

	ages["carol"] = 21

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	if age, ok := ages["bob"]; !ok {
		println(age)
	}

}
