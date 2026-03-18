package main

import "fmt"

func main() {

	var pointerInt *int

	age := 44

	fmt.Println("pointerInt:", pointerInt)

	pointerInt = &age
	fmt.Println("pointerInt:", pointerInt)

	fmt.Println("age value is:", *pointerInt)

	age = 55
	fmt.Println("new age value is:", *pointerInt)

	*pointerInt = 66
	fmt.Println("updated age value is:", age)

	updateAge(age)
	fmt.Println("after updateAge:", age)

	updateAgeWithPointer(&age)
	fmt.Println("after updateAgeWithPointer:", age)

	john := User{"John", 33}
	john.updateAge()
	fmt.Println("updated age value is:", john.age)

	johnAddress := &john
	johnAddress.updateAgeWithPointer()
	fmt.Println("updated age for john is:", john.age)

	employeeBuilder, err := NewEmployeeBuilder().Name("John").Age(33).Email("abc@abc.com").Build()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(employeeBuilder)
}

func updateAge(age int) {
	age = age + 1
}

func updateAgeWithPointer(a *int) {
	*a = *a + 1
}

type User struct {
	name string
	age  int
}

func (u User) updateAge() {
	u.age = u.age + 1
}

func (u *User) updateAgeWithPointer() {
	u.age = u.age + 1
}
