package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

// 返回结构体指针
func createPersonPointer(firstName, lastName string) *Person {
	p := Person{FirstName: firstName, LastName: lastName}
	return &p
}

// 返回结构体
func createPerson(firstName, lastName string) Person {
	return Person{FirstName: firstName, LastName: lastName}
}

func main() {
	// 返回结构体指针的示例
	personPointer := createPersonPointer("John", "Doe")
	fmt.Println("First Name (Pointer):", personPointer.FirstName)
	personPointer.LastName = "Smith"
	fmt.Println("Updated Last Name (Pointer):", personPointer.LastName)

	fmt.Println("-------------------------------------")

	person := createPerson("Alice", "Johnson")
	fmt.Println("First Name:", person.FirstName)
	person.LastName = "Brown"
	fmt.Println("Updated Last Name:", person.LastName)

	fmt.Println("-------------------------------------")

	// 返回结构体指针的示例
	personPointer1 := createPersonPointer("John", "Doe")
	personPointer2 := personPointer1  // 多个指针指向同一结构体
	personPointer2.LastName = "Smith" // 修改结构体字段
	fmt.Println("Last Name (Pointer 1):", personPointer1.LastName)
	fmt.Println("Last Name (Pointer 2):", personPointer2.LastName)

	// 返回结构体的示例
	person1 := createPerson("Alice", "Johnson")
	person2 := person1         // 创建新的结构体副本
	person2.LastName = "Brown" // 修改结构体字段
	fmt.Println("Last Name (Struct 1):", person1.LastName)
	fmt.Println("Last Name (Struct 2):", person2.LastName)
}
