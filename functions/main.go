package main

import (
	"fmt"

	"github.com/icrowley/fake"
)

func main() {
	result := addInts(2, 3)
	fmt.Println("result of addition of 2 + 3 is ", result)

	fmt.Println("\n result of addition of 4 + 4 is %d", addInts(4, 4))

	//var salary float32
	//
	//result = addInts(100, salary)

	//var name string
	//name = addInts(33, 33)
	//calculateDiscount(444)me,
	name, city, day, m := randomEmployeeDetails()

	fmt.Printf("Random employee details %s, %s, %d, %s\n", name, city, day, m)
}

func addInts(x, y int) int {
	return x + y
}

func randomEmployeeDetails() (string, string, int, string) {
	name := fake.FullName()
	city := fake.City()
	day := fake.Day()
	month := fake.Month()
	return name, city, day, month
}
