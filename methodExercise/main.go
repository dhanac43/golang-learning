package main

import "fmt"

// Create a custom type based on float64
type temperature float64

type car struct {
	brand string
	year  int
}

type bankAccount struct {
	owner   string
	balance float64
}

func (t temperature) display() {
	fmt.Println("Temperature is:", t)
}

func (c car) printDetails() {
	fmt.Println("Brand:", c.brand)
	fmt.Println("Year:", c.year)
}

// Method to deposit money
func (b *bankAccount) deposit(amount float64) {
	b.balance += amount
}

func main() {
	var temp temperature = 36.6
	temp.display()

	myCar := car{
		brand: "Toyota",
		year:  2023,
	}

	myCar.printDetails()

	account := bankAccount{
		owner:   "John",
		balance: 100.0,
	}

	account.deposit(50.0)

	fmt.Println("Owner:", account.owner)
	fmt.Println("Balance:", account.balance)
}
