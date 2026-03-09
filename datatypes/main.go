package main

import "fmt"

func main() {

	//var intVar1 int
	//fmt.Println(intVar1)
	//
	//var intVar2 int8
	//fmt.Println(intVar2)
	//
	//var intVar3 int16
	//fmt.Println(intVar3)
	//
	//var intVar4 int32
	//fmt.Println(intVar4)
	//
	//var intVar5 int64
	//fmt.Println(intVar5)
	//
	//var intVar6 uint
	//fmt.Println(intVar6)
	//
	//var intVar7 uint8
	//fmt.Println(intVar7)
	//
	//var intVar8 uint16
	//fmt.Println(intVar8)
	//
	//age3 := 44
	//fmt.Println(age3)
	//
	//var salary float32
	//var bonus float64
	//fmt.Println("salary  %.2f %T", salary, salary)
	//fmt.Println("bonus  %.2f %T", bonus, bonus)

	//num := 100
	//fmt.Println(num)

	var age int = 30
	name := "Dhana"
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)

	var count int
	var price float64
	var username string
	var isAdmin bool

	fmt.Println("count:", count)
	fmt.Println("price:", price)
	fmt.Println("username:", username)
	fmt.Println("isAdmin:", isAdmin)

	const companyName = "River"
	fmt.Println("companyName:", companyName)
	const maxRetries = 3
	fmt.Println("maxRetries:", maxRetries)
	const apiVersion = "v1"
	fmt.Println("apiVersion:", apiVersion)

	const old = 30
	fmt.Println(old)

	productName := "Laptop"
	rate := 500.0
	quantity := 3

	total := rate * float64(quantity)

	fmt.Println("Product:", productName)
	fmt.Println("Rate:", rate)
	fmt.Println("Quantity:", quantity)
	fmt.Println("Total:", total)

}
