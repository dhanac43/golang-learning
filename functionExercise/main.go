package main

import "fmt"

type User struct {
	name string
	age  int
}

func increment(n *int) {
	*n = *n + 1
}

func (u *User) updateAge(newAge int) {
	u.age = newAge
}

func (u *User) updateName(newName string) {
	u.name = newName
}

func main() {

	num := 10
	increment(&num)
	fmt.Println(num) // 11

	user := User{
		name: "John",
		age:  25,
	}
	// Update using pointer
	user.updateAge(30)
	fmt.Println("Updated age:", user.age) // 30

	user.updateName("Mike")
	fmt.Println("Name after update:", user.name)
}
