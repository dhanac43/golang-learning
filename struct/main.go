package main

import (
	"fmt"
)

type productID int64

type user struct {
	userName string
	age      int
	email    string
	salary   float64
}

type admin struct {
	employee   user
	department string
}

func main() {

	var user1 user
	fmt.Printf("%+v\n", user1)

	john := user{"john", 33, "john@abc.com", 444444.33}
	fmt.Printf("%+v\n", john)

	bill := user{
		userName: "bill",
		age:      42,
		email:    "bill@abu.com",
		salary:   99.99,
	}
	fmt.Printf("%+v\n", bill)

	fmt.Println("John age is", john.age)
	fmt.Println("Bill age is", bill.age)

	var peter user
	peter.userName = "peter"
	peter.age = 42
	peter.salary = 99.99
	peter.email = "peter@abc.com"
	fmt.Printf("%+v\n", peter)

	admin1 := admin{
		employee: user{
			userName: "john",
			age:      42,
			email:    "john@abc.com",
			salary:   1000.00,
		},
		department: "peopleOperation",
	}

	admin2 := admin{
		employee:   user{"peter", 33, "peter@abc.com", 99.00},
		department: "IT",
	}

	fmt.Printf("%+v\n", admin1)
	fmt.Printf("%+v\n", admin2)

	fmt.Println("admin 1's name is", admin1.employee.userName)

	var p productID
	p = 3333333
	productInfo(p)

	var v1 int64
	p = productID(v1)

	type myInt = int64
	var v3 myInt = 123
	var v4 int64
	v3 = v4
	fmt.Printf("%d\n", v3)
}

func productInfo(pId productID) {
	fmt.Println("Product ID is", pId)
}
