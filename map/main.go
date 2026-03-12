package main

import (
	"fmt"
	"strings"
)

func main() {

	//employees := make(map[string]string)
	//employees["1"] = "Alice"
	//employees["2"] = "Bob"
	//employees["3"] = "Charlie"
	//
	//for k, v := range employees {
	//	fmt.Println("key is ", k)
	//	fmt.Println("val is ", v)
	//}
	////id := 3
	////fmt.Println("name at the ids is", employees[id])
	//
	//var salaries map[int]float64
	////salaries[1] = 333.333
	//salaries = map[int]float64{1: 333.00, 2: 444.00}
	//fmt.Println(salaries)
	//
	//addresses := make(map[string]string)
	//addresses["1"] = "Alice"
	//addresses["2"] = "Bob"
	//fmt.Println(addresses)

	//Exercise-1:-phone book
	phoneBook := make(map[string]string)

	phoneBook["Alice"] = "123-456"
	phoneBook["Bob"] = "234-567"
	phoneBook["Charlie"] = "345-678"
	phoneBook["Diana"] = "456-789"

	fmt.Println("Phone Book:")
	for k, v := range phoneBook {
		fmt.Println("key is", k)
		fmt.Println("val is", v)
	}

	phone, ok := phoneBook["Alice"]
	if ok {
		fmt.Println("Alice's number:", phone)
	} else {
		fmt.Println("Alice not found")
	}

	delete(phoneBook, "Bob")

	fmt.Println("Entries after deletion:", len(phoneBook))

	//Exercise 2 — Word Counter

	sentence := "the cat sat on the mat the cat"

	words := strings.Split(sentence, " ")

	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	for word, count := range wordCount {
		fmt.Printf(" %s : %d\n", word, count)
	}
}
