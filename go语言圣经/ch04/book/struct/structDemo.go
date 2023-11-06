package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

type Point struct {
	X, Y int
}

type address struct {
	hostname string
	port     int
}

var dilbert Employee

func main() {
	dilbert.ID = 1
	dilbert.Name = "Amy"
	dilbert.Address = "shanghai"
	dilbert.DoB = time.Now()
	dilbert.Position = "manager"
	dilbert.Salary = 10000
	dilbert.ManagerID = 10

	fmt.Println("ID:", dilbert.ID)
	fmt.Println("Name:", dilbert.Name)
	fmt.Println("Address:", dilbert.Address)
	fmt.Println("DoB:", dilbert.DoB)
	fmt.Println("Position:", dilbert.Position)
	fmt.Println("Salary:", dilbert.Salary)
	fmt.Println("ManagerID:", dilbert.ManagerID)

	fmt.Println("------------------------------")

	dilbert.Salary += 5000
	position := &dilbert.Position
	*position = "Senior" + *position

	fmt.Println("ID:", dilbert.ID)
	fmt.Println("Name:", dilbert.Name)
	fmt.Println("Address:", dilbert.Address)
	fmt.Println("DoB:", dilbert.DoB)
	fmt.Println("Position:", dilbert.Position)
	fmt.Println("Salary:", dilbert.Salary)
	fmt.Println("ManagerID:", dilbert.ManagerID)

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	(*employeeOfTheMonth).Position += " (proactive team player)"

	fmt.Println("------------------------------")
	fmt.Println("Position:", dilbert.Position)

	pp := &Point{1, 2}

	p2 := new(Point)
	*p2 = Point{1, 2}
	fmt.Println(pp)

	fmt.Println("------------------------------")

	p := Point{1, 2}
	q := Point{2, 1}
	fmt.Println(p.X == q.X && p.Y == q.Y)
	fmt.Println(p == q)

	fmt.Println("------------------------------")

	hits := make(map[address]int)
	hits[address{"google.org", 443}]++

}
