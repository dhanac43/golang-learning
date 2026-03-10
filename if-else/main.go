package main

import "fmt"

func main() {
	//marksA, marksB := 90, 66

	//if marksA >= 75  && marksB >=75 {
	//	log.Println("First Class")
	//} else if marksA >= 50  || marksB >= 50  {
	//	log.Println("Second Class")
	//} else if marksA >= 35 {
	//	log.Println("Third Class")
	//} else {
	//	log.Println("Failed")

	//if average := marksA + marksB/2; average > 50 {
	//	log.Println("average is ", average)
	//}
	//prints numbers 1 to 10 using a for loop.
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	//Write an if statement that prints:Even number
	num := 9
	if num%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd")
	}
	//Loop through this slice and print each fruit.

	fruits := []string{"apple", "banana", "orange"}

	for _, fruit := range fruits {
		fmt.Println(fruit)
	}
	// skip multiples of 3

	for i := 1; i <= 20; i++ {
		if i%3 == 0 {
			continue
		}
		fmt.Println(i)
	}
	//Count how many even numbers exist in this slice.
	nums := []int{3, 6, 7, 10, 13, 14}
	count := 0

	for _, num := range nums {
		if num%2 == 0 {
			count++
		}
	}

	fmt.Printf("Even numbers: %d\n", count)

}
