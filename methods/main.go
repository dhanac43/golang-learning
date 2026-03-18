package main

import "fmt"

type Employee struct {
	EmployeeID int
	FirstName  string
}

func (e Employee) PrintEmployeeInfo() {
	fmt.Printf("Employee ID: %d\n", e.EmployeeID)
	fmt.Printf("First Name: %s\n", e.FirstName)
}

func (e Employee) GetEmployeeID() int {
	return e.EmployeeID
}

func main() {
	e1 := Employee{EmployeeID: 1, FirstName: "Alex"}
	e1.PrintEmployeeInfo()
	fmt.Println(e1.GetEmployeeID())
}
