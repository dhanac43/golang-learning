package main

import "fmt"

func main() {

	employees := make(map[string]string)
	employees["1"] = "Alice"
	employees["2"] = "Bob"
	employees["3"] = "Charlie"

	for k, v := range employees {
		fmt.Println("key is ", k)
		fmt.Println("val is ", v)
	}
	//id := 3
	//fmt.Println("name at the ids is", employees[id])

	var salaries map[int]float64
	//salaries[1] = 333.333
	salaries = map[int]float64{1: 333.00, 2: 444.00}
	fmt.Println(salaries)

	addresses := make(map[string]string)
	addresses["1"] = "Alice"
	addresses["2"] = "Bob"
	fmt.Println(addresses)
}
