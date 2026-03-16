package main

import "fmt"

func main() {

	//Exercise 1:Write a function multiply that returns the product of two integers.

	fmt.Println(multiply(3, 4))

	//Exercise 2:Write a function that returns both sum and difference of two numbers.
	sum, diff := sumAndDifference(10, 4)
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)

	//Exercise 3:Write a variadic function that finds the maximum number.
	fmt.Println(max(1, 5, 3, 10, 2))

	//Exercise 4:-Create a program where an init() function prints:
	fmt.Println("Program running")
}

// Exercise 1:Write a function multiply that returns the product of two integers.
func multiply(x, y int) int {
	return x * y
}

// Exercise 2:Write a function that returns both sum and difference of two numbers.
func sumAndDifference(a, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

// Exercise 3:Write a variadic function that finds the maximum number.
func max(nums ...int) int {
	maxValue := nums[0]

	for _, n := range nums {
		if n > maxValue {
			maxValue = n
		}
	}

	return maxValue
}

// Exercise 4:-Create a program where an init() function prints:
func init() {
	fmt.Println("Application starting...")
}
