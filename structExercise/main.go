package main

import "fmt"

// Define a struct called book
type book struct {
	title  string
	author string
	price  float64
}

// Define a custom type temperature
type temperature float64

func main() {
	// Create a book variable
	myBook := book{
		title:  "The Go Programming Language",
		author: "Alan",
		price:  25.99,
	}

	// Print the book details
	fmt.Println("Book Details:")
	fmt.Println("Title:", myBook.title)
	fmt.Println("Author:", myBook.author)
	fmt.Println("Price:", myBook.price)

	// Create a temperature variable
	var temp temperature = 36.6
	fmt.Println("Temperature:", temp)
}
