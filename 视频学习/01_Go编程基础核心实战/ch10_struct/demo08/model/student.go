package model

type student struct {
	Name string
	Age  int
}

func NewStudent(n string, a int) *student {
	return &student{
		Name: n,
		Age:  a,
	}
}
