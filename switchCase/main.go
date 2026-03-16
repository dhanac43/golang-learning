package main

import "fmt"

func main() {

	//day := "tuesday"
	//
	//switch day {
	//case "Monday":
	//	fmt.Println("calling task1")
	//case "Tuesday", "tuesday":
	//	fmt.Println("calling task2")
	//case "Wednesday":
	//	fmt.Println("calling task3")
	//case "Thursday":
	//	fmt.Println("calling task4")
	//case "Friday":
	//	fmt.Println("calling task5")
	//case "Saturday":
	//	fmt.Println("calling task6")
	//default:
	//	fmt.Println("calling default task")
	//
	//}

	marks := 33
	switch {

	case marks <= 30:
		fmt.Println("marks <= 30 - Failed")
	case marks > 30 && marks <= 50:
		fmt.Println("3rd class")
	case marks >= 50 && marks < 60:
		fmt.Println("3rd class")
	case marks >= 60:
		fmt.Println("1st class")

	}

	currency := "INR"
	switch currency {
	case "USD":
		fmt.Println("tasks for USD")
	case "EUR":
		fmt.Println("tasks for EUR")
		fallthrough
	case "GBP":
		fmt.Println("tasks for GBP")
	case "AUD":
		fmt.Println("tasks for AUD")
	default:
		fmt.Println("tasks for default currency")
	}
	//Exercise 1:-Day Type Checker
	day := "Tuesday"

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")

	case "Saturday", "Sunday":
		fmt.Println("Weekend")

	default:
		fmt.Println("Unknown")
	}

	//Exercise:2:-HTTP Status Code Explainer
	statusCode := 404

	switch statusCode {

	case 200:
		fmt.Println("200 OK - Request successful")

	case 301:
		fmt.Println("301 Moved Permanently - Resource has been moved")

	case 404:
		fmt.Println("404 Not Found - Resource could not be found")

	case 500:
		fmt.Println("500 Internal Server Error - Server encountered an error")

	default:
		fmt.Println("Unknown status code")

	}

}
