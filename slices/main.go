package main

import "fmt"

func main() {

	var intSlice []int
	fmt.Println(intSlice)
	fmt.Println(len(intSlice))
	fmt.Println(cap(intSlice))

	var floatSlice = []float64{33.2, 22.1}
	fmt.Println(floatSlice)
	fmt.Println(len(floatSlice))
	fmt.Println(cap(floatSlice))

	stringSlice := []string{"first", "second", "third"}
	fmt.Println(stringSlice)
	fmt.Println(len(stringSlice))
	fmt.Println(cap(stringSlice))

	employeeIDs := make([]int, 10)
	fmt.Println("employeeIDs:", employeeIDs)
	fmt.Println("total employee ids:", len(employeeIDs))
	fmt.Println(cap(employeeIDs))

	employeeNames := make([]string, 10, 20)
	fmt.Println("employeeIDs:", employeeNames)
	fmt.Println("total employee ids:", len(employeeNames))
	fmt.Println(cap(employeeNames))

	intSlice = append(intSlice, 33)
	fmt.Println(intSlice)
	fmt.Println(len(intSlice))
	fmt.Println(cap(intSlice))

	floatSlice = append(floatSlice, 13.2)
	fmt.Println(floatSlice)
	fmt.Println(len(floatSlice))
	fmt.Println(cap(floatSlice))

	floatSlice = append(floatSlice, 44.2)
	fmt.Println(floatSlice)
	fmt.Println(len(floatSlice))
	fmt.Println(cap(floatSlice))

	floatSlice = append(floatSlice, 64.2, 77.2)
	fmt.Println(floatSlice)
	fmt.Println(len(floatSlice))
	fmt.Println(cap(floatSlice))

	newFloatSlice := floatSlice[2:4]
	fmt.Println(newFloatSlice)
	fmt.Println(len(newFloatSlice))
	fmt.Println(cap(newFloatSlice))

	//Create a slice with values:
	numbers := []int{10, 20, 30}
	numbers = append(numbers, 40, 50)
	fmt.Println(numbers)

	//Create an array-example-3
	number1 := [5]int{10, 20, 30, 40, 50}
	slice := number1[1:4]
	fmt.Println(slice)
	//Create a slice using make() with = example-4
	numberSlice := make([]int, 2, 5)
	fmt.Println("Length:", len(numberSlice))
	fmt.Println("Capacity:", cap(numberSlice))
}
